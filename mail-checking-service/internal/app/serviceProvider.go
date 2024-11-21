package app

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/configs"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/controller"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/logger"
	"go.uber.org/zap"
)

const (
	topicName = "registration"
)

type serviceProvider struct {
	grpcConfig  configs.GRPCConfigs
	kafkaConfig configs.KafkaConfig

	kafkaConsumer *kafka.Consumer

	controller   *controller.Controller
	kafkaHandler *kafka.KafkaHandler
}

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GRPCConfig initializes and returns the gRPC configuration if not already set.
func (s *serviceProvider) GRPCConfig() configs.GRPCConfigs {
	if s.grpcConfig == nil {
		cfg, err := configs.NewGRPCConfig()
		if err != nil {
			logger.Fatal("failed to get grpc config", zap.Error(err))
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) KafkaConfig() configs.KafkaConfig {
	if s.kafkaConfig == nil {
		cfg, err := configs.NewKafkaConfig()
		if err != nil {
			logger.Fatal("failed to get kafka config", zap.Error(err))
		}
		s.kafkaConfig = cfg
	}
	return s.kafkaConfig
}

// Controller initializes and returns the controller layer to handle business logic requests.
func (s *serviceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New()
	}
	return s.controller
}

func (s *serviceProvider) KafkaHandler() *kafka.KafkaHandler {
	if s.kafkaHandler == nil {
		handler := kafka.NewKafkaHandler()
		s.kafkaHandler = handler
	}
	return s.kafkaHandler
}

func (s *serviceProvider) KafkaConsumer(number int) *kafka.Consumer {
	if s.kafkaConsumer == nil {
		consumer, err := kafka.NewConsumer(s.KafkaHandler(),
			s.KafkaConfig().Address(), topicName, "mail", number)
		if err != nil {
			logger.Fatal("failed to get kafka consumer", zap.Error(err))
		}
		s.kafkaConsumer = consumer

	}
	return s.kafkaConsumer
}
