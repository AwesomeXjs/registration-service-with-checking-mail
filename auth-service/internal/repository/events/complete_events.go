package events

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// CompleteEvents updates the status of events with the provided IDs to "complete".
func (e *EventRepository) CompleteEvents(ctx context.Context, completeIDs []int) error {

	const mark = "Repository.Events.CompleteEvents"

	updateBuilder := squirrel.Update("outbox").
		PlaceholderFormat(squirrel.Dollar).
		Set("status", "complete").
		Where(squirrel.Eq{IDColumn: completeIDs})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build update query", mark, zap.Error(err))
		return fmt.Errorf("failed to build update query: %v", err)
	}

	q := db.Query{
		Name:     "CompleteEvents",
		QueryRaw: query,
	}

	_, err = e.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to complete events", mark, zap.Error(err))
		return fmt.Errorf("failed to complete events: %v", err)
	}

	return nil
}
