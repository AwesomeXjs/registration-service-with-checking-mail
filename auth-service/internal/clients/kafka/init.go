package kafka

import (
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// IProducer defines an interface for a Kafka producer with methods for sending messages and closing.
type IProducer interface {
	// Produce sends a message to a specified Kafka topic.
	Produce(message, topic, key string, timestamp time.Time) error

	// Close gracefully shuts down the producer and flushes pending messages.
	Close() error
}

// Producer wraps a Kafka producer to handle message production to Kafka topics.
type Producer struct {
	// producer is the underlying Kafka producer instance.
	producer *kafka.Producer
}

// NewProducer creates and initializes a new Producer with the given Kafka broker addresses.
func NewProducer(addresses []string) (*Producer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(addresses, ","),
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return &Producer{producer: producer}, nil
}
