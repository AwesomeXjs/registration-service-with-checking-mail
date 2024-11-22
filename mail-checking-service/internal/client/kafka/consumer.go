package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

const (
	sessionTimeout = 7000
	noTimeout      = -1
)

type Consumer struct {
	consumer       *kafka.Consumer
	stop           bool
	handler        Handler
	consumerNumber int
}

func NewConsumer(handler Handler, addresses []string, topic string, consumerGroup string, consumerNumber int) (*Consumer, error) {
	// конфигурация
	config := &kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(addresses, ","),
		"group.id":                 consumerGroup,
		"session.timeout.ms":       sessionTimeout,
		"enable.auto.offset.store": false, // автоматическое сохранение оффсета в локальную память консьюмера
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  5000,
		"auto.offset.reset":        "earliest",
	}

	// создание консьюмера
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	// подписка на топик
	if err = consumer.Subscribe(topic, nil); err != nil {
		return nil, fmt.Errorf("failed to subscribe: %w", err)
	}

	return &Consumer{
		consumer:       consumer,
		handler:        handler,
		consumerNumber: consumerNumber,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		if c.stop {
			break
		}
		kafkaMsg, err := c.consumer.ReadMessage(noTimeout)
		if err != nil {
			fmt.Printf("Failed to read message: %s\n", err)
			continue
		}
		if kafkaMsg == nil {
			continue
		}

		if err = c.handler.HandleMessage(ctx, kafkaMsg.Value, kafkaMsg.TopicPartition.Offset, c.consumerNumber); err != nil {
			logger.Error("failed to handle message", zap.Error(err))
		}

		if _, err = c.consumer.StoreMessage(kafkaMsg); err != nil {
			fmt.Printf("Failed to store message: %s\n", err)
		}
	}
}

func (c *Consumer) Stop() error {
	c.stop = true
	if _, err := c.consumer.Commit(); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	fmt.Println("commited offset")
	return c.consumer.Close()
}
