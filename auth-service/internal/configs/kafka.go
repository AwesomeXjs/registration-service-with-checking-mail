package configs

import (
	"fmt"
	"os"
)

const (
	Kafka1 = "KAFKA3"
	Kafka2 = "KAFKA2"
	Kafka3 = "KAFKA3"
)

type KafkaConfig interface {
	Address() []string
}

type kafkaConfig struct {
	kafka1 string
	kafka2 string
	kafka3 string
}

func NewKafkaConfig() (KafkaConfig, error) {
	kafka1 := os.Getenv(Kafka1)
	if len(kafka1) == 0 {
		return nil, fmt.Errorf("KAFKA_1 is not set")
	}
	kafka2 := os.Getenv(Kafka2)
	if len(kafka2) == 0 {
		return nil, fmt.Errorf("KAFKA_2 is not set")
	}
	kafka3 := os.Getenv(Kafka3)
	if len(kafka3) == 0 {
		return nil, fmt.Errorf("KAFKA_3 is not set")
	}

	return &kafkaConfig{
		kafka1: kafka1,
		kafka2: kafka2,
		kafka3: kafka3,
	}, nil
}

func (m *kafkaConfig) Address() []string {
	return []string{m.kafka1, m.kafka2, m.kafka3}
}
