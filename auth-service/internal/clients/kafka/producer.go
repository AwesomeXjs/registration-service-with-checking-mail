package kafka

import (
	"errors"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// unknownType represents an error for unrecognized event types.
var errUnknownType = errors.New("unknown event type")

const (
	// flushTimeout defines the timeout in milliseconds for flushing the Kafka producer.
	flushTimeout = 5000
)

// Produce sends a message with a key and timestamp to the specified Kafka topic.
func (p *Producer) Produce(message, topic, key string) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Key:            []byte(key),
		Timestamp:      time.Now(),
	}
	kafkaChan := make(chan kafka.Event)

	if err := p.producer.Produce(kafkaMessage, kafkaChan); err != nil {
		return err
	}

	event := <-kafkaChan
	switch event.(type) {
	case *kafka.Message:
		m := event.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return m.TopicPartition.Error
		}
		return nil
	default:
		return errUnknownType
	}
}

// Close gracefully shuts down the Kafka producer after flushing pending messages.
func (p *Producer) Close() error {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
	return nil
}
