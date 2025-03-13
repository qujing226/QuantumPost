package sarama

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/qujing226/quantum_post/pkg/log"
)

type Handler[T any] struct {
	l  log.Logger
	fn func(msg *sarama.ConsumerMessage, t T) error
}

func NewHandler[T any](logger log.Logger, fn func(msg *sarama.ConsumerMessage, t T) error) sarama.ConsumerGroupHandler {
	return &Handler[T]{
		l:  logger,
		fn: fn,
	}
}

func (h *Handler[T]) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler[T]) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgCh := claim.Messages()
	for msg := range msgCh {
		var t T
		err := json.Unmarshal(msg.Value, &t)
		if err != nil {
			h.l.Error("MQ undo serialize failed",
				log.Error(err),
				log.String("topic: ", msg.Topic),
				log.Int32("partition: ", msg.Partition),
				log.Int64("offset: ", msg.Offset))
			continue
		}
		// 重试策略
		for i := 0; i < 3; i++ {
			err = h.fn(msg, t)
			if err == nil {
				break
			}
			h.l.Error("MQ handler logic error",
				log.Error(err),
				log.String("topic: ", msg.Topic),
				log.Int32("partition: ", msg.Partition),
				log.Int64("offset: ", msg.Offset))
		}
		if err != nil {
			h.l.Error("MQ handler logic error, all retry failed: ",
				log.Error(err),
				log.String("topic: ", msg.Topic),
				log.Int32("partition: ", msg.Partition),
				log.Int64("offset: ", msg.Offset))
		} else {
			session.MarkMessage(msg, "pass")
		}
	}
	return nil
}
