package repository

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) ConfirmEmail(ctx context.Context, email string) error {
	queryBuilder := squirrel.Update(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(Verification, true).
		Where(squirrel.Eq{EmailColumn: email})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", zap.Error(err))
		return fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "Confirm Email",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	fmt.Println(res)
	if err != nil {
		logger.Error("failed to confirm email", zap.Error(err))
		return fmt.Errorf("failed to confirm email: %v", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Error("user not found", zap.Error(err))
		return fmt.Errorf("user not found")
	}

	return nil
}
