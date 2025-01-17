package events

import (
	"context"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
)

// GetEvents retrieves up to 10 pending events from the outbox table.
func (e *EventRepository) GetEvents(ctx context.Context) ([]model.EventData, error) {

	const mark = "Repository.Events.GetEvents"

	builderSelect := squirrel.
		Select(IDColumn, EventColumn, PayloadColumn, CreatedAtColumn).
		PlaceholderFormat(squirrel.Dollar).
		From(TableNameOutbox).
		Where(squirrel.Eq{StatusColumn: "pending"}).
		Limit(10)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		logger.Error("failed to build select query", mark, zap.Error(err))
		return nil, err
	}

	q := db.Query{
		Name:     "GetEvents",
		QueryRaw: query,
	}

	rows, err := e.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to get events", mark, zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var events []model.EventData
	for rows.Next() {
		var event model.EventData
		if err := rows.Scan(&event.ID, &event.Event, &event.Payload, &event.CreatedAt); err != nil {
			logger.Error("failed to scan row", mark, zap.Error(err))
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
