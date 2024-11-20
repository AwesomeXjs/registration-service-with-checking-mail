package kafka

import (
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	sessionTimeout = 7000
	noTimeout      = -1
)

type Handler interface {
	HandleMessage(kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error
}

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
		consumerNumber: consumerNumber}, nil
}

func (c *Consumer) Start() {
	for {
		// если консьюмер остановлен - выходим из цикла
		if c.stop {
			break
		}
		// читаем сообщение
		kafkaMsg, err := c.consumer.ReadMessage(noTimeout)
		if err != nil {
			fmt.Printf("Failed to read message: %s\n", err)
			continue
		}
		if kafkaMsg == nil {
			continue
		}

		// тут обрабатываем сообщение (передаем в другой сервис)
		if err = c.handler.HandleMessage(kafkaMsg.Value, kafkaMsg.TopicPartition.Offset, c.consumerNumber); err != nil {
			fmt.Printf("Failed to handle message: %s\n", err)
		}

		// после обработки мы должны вручную сохранить оффсет локально
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
