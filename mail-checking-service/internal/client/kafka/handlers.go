package kafka

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/logger"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type Handler interface {
	HandleMessage(kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error
}

type KafkaHandler struct{}

func NewKafkaHandler() *KafkaHandler {
	return &KafkaHandler{}
}

func (h *KafkaHandler) HandleMessage(kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error {
	logger.Info("message received", zap.Int("consumer", consumerNumber), zap.String("message", string(kafkaMsg)), zap.Int64("offset", int64(offset)))
	return nil
}
