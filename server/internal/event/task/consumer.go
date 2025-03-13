package task

import "github.com/IBM/sarama"

type KafkaTaskConsumer struct {
	client sarama.Client
	repo   repository
}
