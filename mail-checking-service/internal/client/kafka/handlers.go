package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler interface {
	HandleMessage(kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error
}

type KafkaHandler struct {
}

func NewKafkaHandler() *KafkaHandler {
	return &KafkaHandler{}
}

func (h *KafkaHandler) HandleMessage(kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error {
	fmt.Println("message:", string(kafkaMsg), "offset:", offset, "consumer number:", consumerNumber)
	return nil
}
