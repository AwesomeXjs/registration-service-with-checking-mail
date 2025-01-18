package eventsender

import (
	"context"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/metrics"

	"github.com/goccy/go-json"
	"go.uber.org/zap"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/events"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
)

// ISender defines the interface for sending events at regular intervals.
type ISender interface {
	// Start begins the process of fetching and sending events at a specified interval.
	Start(ctx context.Context, period time.Duration)
}

// Sender is responsible for sending events to a Kafka producer and marking them as complete.
type Sender struct {
	producer kafka.IProducer         // Kafka producer to send events.
	db       db.Client               // Database client for accessing event repository.
	repo     events.IEventRepository // Event repository for fetching and completing events.
}

// New creates a new Sender instance with the provided Kafka producer, database client, and event repository.
func New(producer kafka.IProducer, db db.Client, repo events.IEventRepository) *Sender {
	return &Sender{
		producer: producer,
		db:       db,
		repo:     repo,
	}
}

// Start starts a loop to fetch events periodically and send them to Kafka.
func (s *Sender) Start(ctx context.Context, period time.Duration) {
	const mark = "Service.EventSender.Start"

	go func() {
		for {
			// Fetch events from the repository
			data, err := s.repo.GetEvents(ctx)
			if err != nil {
				logger.Error("failed to get events", mark, zap.Error(err))
				continue
			}

			// If no events, continue to next cycle
			if len(data) == 0 {
				continue
			}

			// Send events to Kafka and get their IDs for completion
			ids := SendEvents(data, s)

			// Mark events as completed in the repository
			err = s.repo.CompleteEvents(ctx, ids)
			if err != nil {
				logger.Error("failed to complete events", mark, zap.Error(err))
				continue
			}

			// Wait for the specified period or termination signal
			select {
			case <-ctx.Done():
				return
			case <-time.After(period):
			}
		}
	}()
}

// SendEvents processes the fetched events, sends them to Kafka, and returns the IDs of successfully processed events.
func SendEvents(data []model.EventData, s *Sender) []int {
	const mark = "Service.EventSender.SendEvents"

	successIDs := make([]int, 0)
	for _, event := range data {
		var payload model.Payload

		// Unmarshal the event payload
		if err := json.Unmarshal([]byte(event.Payload), &payload); err != nil {
			logger.Error("failed to unmarshal payload", mark, zap.Error(err))
			continue
		}

		// Produce the event to Kafka
		if err := s.producer.Produce(payload.Message, payload.Topic, payload.Key); err != nil {
			logger.Error("failed to produce event", mark, zap.Error(err))
			continue
		}
		metrics.IncVerificationCounter()
		successIDs = append(successIDs, event.ID)
	}
	return successIDs
}
