package kafka

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var unknownType = errors.New("unknown event type")

const (
	flushTimeout = 5000
)

type IProducer interface {
	Produce(message, topic, key string, timestamp time.Time) error
	Close() error
}

type Producer struct {
	producer *kafka.Producer
}

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

func (p *Producer) Produce(message, topic, key string, timestamp time.Time) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Key:            []byte(key),
		Timestamp:      timestamp,
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
		return unknownType
	}
}

func (p *Producer) Close() error {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
	return nil
}
