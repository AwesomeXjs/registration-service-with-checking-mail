package events

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
)

// IEventRepository defines the interface for event-related operations.
type IEventRepository interface {
	// SendEvent inserts a new event into the database.
	SendEvent(ctx context.Context, event *model.SendEvent) error

	// GetEvents retrieves pending events from the database.
	GetEvents(ctx context.Context) ([]model.EventData, error)

	// CompleteEvents marks events with specified IDs as completed in the database.
	CompleteEvents(ctx context.Context, completeIDs []int) error
}

// EventRepository is the concrete implementation of IEventRepository.
type EventRepository struct {
	db db.Client // Database client for executing queries.
}

// New creates a new instance of EventRepository.
func New(db db.Client) IEventRepository {
	return &EventRepository{
		db: db,
	}
}
