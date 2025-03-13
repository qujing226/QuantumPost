package sarama

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/qujing226/quantum_post/pkg/log"
	"time"
)

func NewBatchConsumerHandler[T any](l log.Logger, fn func(msgs []*sarama.ConsumerMessage, ts []T) error) sarama.ConsumerGroupHandler {
	return &BatchConsumerHandler[T]{
		l:             l,
		fn:            fn,
		batchSize:     10,
		batchDuration: 1 * time.Second,
	}
}

type BatchConsumerHandler[T any] struct {
	l             log.Logger
	fn            func(msgs []*sarama.ConsumerMessage, ts []T) error
	batchSize     int
	batchDuration time.Duration
}

func (b *BatchConsumerHandler[T]) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BatchConsumerHandler[T]) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil

}

func (b *BatchConsumerHandler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgCh := claim.Messages()
	batchSize := b.batchSize
	duration := b.batchDuration
	for {
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		var lastMsg *sarama.ConsumerMessage
		done := false

		msgs := make([]*sarama.ConsumerMessage, 0, batchSize)
		ts := make([]T, 0, batchSize)

		for i := 0; i < batchSize && !done; i++ {
			select {
			case <-ctx.Done():
				done = true
			case msg, ok := <-msgCh:
				if !ok {
					cancel()
				}
				lastMsg = msg
				var t T
				err := json.Unmarshal(msg.Value, &t)
				if err != nil {
					b.l.Error("MQ undo serialize failed",
						log.Error(err),
						log.String("topic: ", msg.Topic),
						log.Int32("partition: ", msg.Partition),
						log.Int64("offset: ", msg.Offset))
					continue
				}
				msgs = append(msgs, msg)
				ts = append(ts, t)
			}
		}
		cancel()
		if len(msgs) == 0 {
			continue
		}
		err := b.fn(msgs, ts)
		if err != nil {
			b.l.Error("batch MQ handler logic error:", log.Error(err))
		}

		if lastMsg != nil {
			session.MarkMessage(lastMsg, "pass")
		}
	}
}
