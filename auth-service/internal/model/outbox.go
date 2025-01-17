package model

import "time"

// EventData represents an event stored in the database with its details.
type EventData struct {
	ID        int       `db:"id"`         // Unique identifier of the event.
	Event     string    `db:"event"`      // Name or type of the event.
	Payload   string    `db:"payload"`    // Data associated with the event.
	CreatedAt time.Time `db:"created_at"` // Timestamp when the event was created.
}

// SendEvent represents an event ready to be sent, excluding metadata.
type SendEvent struct {
	Event   string `db:"event"`   // Name or type of the event.
	Payload string `db:"payload"` // Data associated with the event.
}

// Payload represents the structure of a message payload used for communication.
type Payload struct {
	Message string `json:"key"`     // The key of the payload message.
	Topic   string `json:"topic"`   // The topic associated with the payload.
	Key     string `json:"message"` // The actual message content of the payload.
}
