package events

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
)

// SendEvent inserts a new event into the outbox table.
func (e *EventRepository) SendEvent(ctx context.Context, event *model.SendEvent) error {

	const mark = "Repository.Events.SendEvent"

	builderInsert := squirrel.Insert(TableNameOutbox).PlaceholderFormat(squirrel.Dollar).
		Columns(EventColumn, PayloadColumn).
		Values(event.Event, event.Payload)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return fmt.Errorf("failed to build insert query: %v", err)
	}

	q := db.Query{
		Name:     "SendEvent",
		QueryRaw: query,
	}

	_, err = e.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to send event", mark, zap.Error(err))
		return fmt.Errorf("failed to send event: %v", err)
	}

	return nil
}
