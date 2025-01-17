package events

// Constants for database column and table names related to the outbox table.
const (
	IDColumn        = "id"         // Column name for the unique identifier.
	TableNameOutbox = "outbox"     // Table name for storing events.
	EventColumn     = "event"      // Column name for the event type or name.
	PayloadColumn   = "payload"    // Column name for the event payload data.
	CreatedAtColumn = "created_at" // Column name for the event creation timestamp.
	StatusColumn    = "status"     // Column name for the event status.
)
