package kafka

import (
	"fmt"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

// Kafka environment variable keys.
// These constants represent the environment variable names for Kafka broker addresses.
const (
	FirstKafka  = "KAFKA1"
	SecondKafka = "KAFKA2"
	ThirdKafka  = "KAFKA3"
)

// IKafkaConfig defines an interface for Kafka configuration.
// It includes a method to retrieve Kafka broker addresses.
type IKafkaConfig interface {
	// Address returns a list of Kafka broker addresses.
	Address() []string
}

// kafkaConfig is a concrete implementation of KafkaConfig.
// It stores the Kafka broker addresses as string fields.
type kafkaConfig struct {
	kafka1 string
	kafka2 string
	kafka3 string
}

// NewKafkaConfig initializes a KafkaConfig by reading broker addresses from environment variables.
// Returns:
// - KafkaConfig: An instance of kafkaConfig if all environment variables are set correctly.
// - error: An error if any of the required environment variables are missing.
func NewKafkaConfig() (IKafkaConfig, error) {
	kafka1 := os.Getenv(FirstKafka)
	if len(kafka1) == 0 {
		logger.Error("failed to get kafka host", zap.String("kafka host", FirstKafka))
		return nil, fmt.Errorf("KAFKA_1 is not set")
	}
	kafka2 := os.Getenv(SecondKafka)
	if len(kafka2) == 0 {
		logger.Error("failed to get kafka host", zap.String("kafka host", SecondKafka))
		return nil, fmt.Errorf("KAFKA_2 is not set")
	}
	kafka3 := os.Getenv(ThirdKafka)
	if len(kafka3) == 0 {
		logger.Error("failed to get kafka host", zap.String("kafka host", ThirdKafka))
		return nil, fmt.Errorf("KAFKA_3 is not set")
	}

	return &kafkaConfig{
		kafka1: kafka1,
		kafka2: kafka2,
		kafka3: kafka3,
	}, nil
}

// Address returns a slice of Kafka broker addresses.
// The addresses are retrieved from the internal fields kafka1, kafka2, and kafka3.
func (m *kafkaConfig) Address() []string {
	return []string{m.kafka1, m.kafka2, m.kafka3}
}
