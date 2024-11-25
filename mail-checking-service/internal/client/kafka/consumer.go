package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

// Constants for configuring Kafka consumer behavior.
const (
	sessionTimeout = 7000 // Session timeout in milliseconds. If no heartbeat is received within this time, the consumer will be considered inactive.
	noTimeout      = -1   // Indicates no timeout for reading messages from Kafka.
)

// Consumer represents a Kafka consumer that reads messages from a topic and processes them using a handler.
type Consumer struct {
	consumer       *kafka.Consumer // The underlying Kafka consumer instance.
	stop           bool            // Indicates whether the consumer should stop processing messages.
	handler        IKafkaHandler   // The handler for processing Kafka messages.
	consumerNumber int             // Identifier for this consumer instance (useful in multi-consumer scenarios).
}

// NewConsumer creates a new Kafka consumer instance with the specified configuration.
// Returns a pointer to the Consumer instance or an error if initialization fails.
func NewConsumer(handler IKafkaHandler, addresses []string, topic string, consumerGroup string, consumerNumber int) (*Consumer, error) {
	// Kafka consumer configuration.
	config := &kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(addresses, ","), // List of Kafka broker addresses.
		"group.id":                 consumerGroup,                // Consumer group ID for managing offsets and load balancing.
		"session.timeout.ms":       sessionTimeout,               // Timeout for detecting inactive consumers.
		"enable.auto.offset.store": false,                        // Prevent automatic offset storage; manual storage is used instead.
		"enable.auto.commit":       true,                         // Automatically commit offsets at intervals.
		"auto.commit.interval.ms":  5000,                         // Interval for automatic offset commits.
		"auto.offset.reset":        "earliest",                   // Reset behavior for new consumers (start from the earliest offset).
	}

	// Create a new Kafka consumer.
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	// Subscribe the consumer to the specified topic.
	if err = consumer.Subscribe(topic, nil); err != nil {
		return nil, fmt.Errorf("failed to subscribe: %w", err)
	}

	return &Consumer{
		consumer:       consumer,
		handler:        handler,
		consumerNumber: consumerNumber,
	}, nil
}

// Start begins consuming messages from Kafka and processing them using the provided handler.
// - `ctx`: The context for controlling the lifecycle of the consumer.
func (c *Consumer) Start(ctx context.Context) {
	for {
		if c.stop {
			break // Stop consuming messages if the consumer is stopped.
		}
		// Read a message from Kafka with no timeout.
		kafkaMsg, err := c.consumer.ReadMessage(noTimeout)
		if err != nil {
			fmt.Printf("Failed to read message: %s\n", err)
			continue
		}
		if kafkaMsg == nil {
			continue // Skip processing if the message is nil.
		}

		// Handle the Kafka message using the provided handler.
		if err = c.handler.HandleMessage(ctx, kafkaMsg.Value, kafkaMsg.TopicPartition.Offset, c.consumerNumber); err != nil {
			logger.Warn("failed to handle message", zap.Error(err))
		}

		// Store the message's offset to ensure it can be committed later.
		if _, err = c.consumer.StoreMessage(kafkaMsg); err != nil {
			fmt.Printf("Failed to store message: %s\n", err)
		}
	}
}

// Stop gracefully stops the Kafka consumer and commits the latest offsets.
// Returns an error if the commit or consumer close operation fails.
func (c *Consumer) Stop() error {
	c.stop = true // Signal the consumer to stop processing.
	if _, err := c.consumer.Commit(); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	fmt.Println("committed offset")
	return c.consumer.Close() // Close the consumer connection.
}
