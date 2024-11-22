package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type Handler interface {
	HandleMessage(ctx context.Context, kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error
	sendEmail(ctx context.Context, config IMailConfig, to, subject, body string) error
}

type KafkaHandler struct {
	redisClient redis.IRedis
	configEmail IMailConfig
}

func NewKafkaHandler(redisClient redis.IRedis, configEmail IMailConfig) *KafkaHandler {
	return &KafkaHandler{
		redisClient: redisClient,
		configEmail: configEmail,
	}
}

func (h *KafkaHandler) HandleMessage(ctx context.Context, kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error {
	code := uuid.NewString()

	if err := h.sendEmail(ctx, h.configEmail, string(kafkaMsg), "Verification Code", fmt.Sprintf("Your code is: %s", code)); err != nil {
		logger.Warn("Your code is: ", zap.String("code", code))
		return fmt.Errorf("%w", err)
	}

	if err := h.redisClient.Set(ctx, string(kafkaMsg), code, time.Hour); err != nil {
		return fmt.Errorf("failed to set code: %w", err)
	}

	logger.Info("message received", zap.Int("consumer", consumerNumber), zap.String("message", string(kafkaMsg)), zap.Int64("offset", int64(offset)))
	return nil
}

func (h *KafkaHandler) sendEmail(ctx context.Context, config IMailConfig, to, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "xxx@example.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer(config.GetHost(), config.GetPort(), config.GetUserName(), config.GetPassword())

	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
