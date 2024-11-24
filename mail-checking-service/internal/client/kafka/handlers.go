package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/mail"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/client/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// IKafkaHandler defines an interface for handling Kafka messages and sending emails.
type IKafkaHandler interface {
	// HandleMessage processes a Kafka message.
	// - ctx: Context for managing lifecycle and cancellation.
	// - kafkaMsg: The message received from Kafka as a byte slice.
	// - offset: The Kafka offset of the message.
	// - consumerNumber: The identifier for the consumer processing the message.
	// Returns an error if processing fails.
	HandleMessage(ctx context.Context, kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error
}

// Handler implements the IKafkaHandler interface and provides methods for message handling and email sending.
type Handler struct {
	redisClient redis.IRedis     // Redis client for caching verification codes.
	mailClient  mail.IMailClient // Email configuration for sending emails.
}

// NewKafkaHandler creates a new instance of Handler.
// - redisClient: A Redis client for storing verification codes.
// - configEmail: The email configuration for sending messages.
// Returns an implementation of IKafkaHandler.
func NewKafkaHandler(redisClient redis.IRedis, mailClient mail.IMailClient) IKafkaHandler {
	return &Handler{
		redisClient: redisClient,
		mailClient:  mailClient,
	}
}

// HandleMessage processes a Kafka message by generating a verification code,
// sending it via email, and storing the code in Redis.
// - ctx: Context for managing lifecycle and cancellation.
// - kafkaMsg: The message received from Kafka as a byte slice.
// - offset: The Kafka offset of the message.
// - consumerNumber: The identifier for the consumer processing the message.
// Returns an error if processing fails.
func (h *Handler) HandleMessage(ctx context.Context, kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error {
	// Generate a unique verification code.
	code := uuid.NewString()

	// Send the code via email to the recipient specified in the Kafka message.
	if err := h.mailClient.SendEmail(ctx, string(kafkaMsg), "Verification Code", fmt.Sprintf("Your code is: %s", code)); err != nil {
		logger.Warn("failed to send verification code", zap.String("code", code), zap.String("email", string(kafkaMsg)), zap.Error(err))
		return fmt.Errorf("%w", err)
	}

	// Store the verification code in Redis with a 1-hour expiration.
	if err := h.redisClient.Set(ctx, string(kafkaMsg), code, time.Hour); err != nil {
		return fmt.Errorf("failed to set code: %w", err)
	}

	// Log the successful message processing.
	logger.Info("message received",
		zap.Int("consumer", consumerNumber),
		zap.String("message", string(kafkaMsg)),
		zap.Int64("offset", int64(offset)),
	)
	return nil
}
