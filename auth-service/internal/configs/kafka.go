package configs

import (
	"fmt"
	"os"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

// Kafka environment variable keys.
// These constants represent the environment variable names for Kafka broker addresses.
const (
	Kafka1 = "KAFKA1"
	Kafka2 = "KAFKA2"
	Kafka3 = "KAFKA3"
)

// KafkaConfig defines an interface for Kafka configuration.
// It includes a method to retrieve Kafka broker addresses.
type KafkaConfig interface {
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
func NewKafkaConfig() (KafkaConfig, error) {
	kafka1 := os.Getenv(Kafka1)
	if len(kafka1) == 0 {
		logger.Error("failed to get kafka host", zap.String("kafka host", Kafka1))
		return nil, fmt.Errorf("KAFKA_1 is not set")
	}
	kafka2 := os.Getenv(Kafka2)
	if len(kafka2) == 0 {
		logger.Error("failed to get kafka host", zap.String("kafka host", Kafka2))
		return nil, fmt.Errorf("KAFKA_2 is not set")
	}
	kafka3 := os.Getenv(Kafka3)
	if len(kafka3) == 0 {
		logger.Error("failed to get kafka host", zap.String("kafka host", Kafka3))
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
