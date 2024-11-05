package repository

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// UpdatePassword updates the user's password in the database based on the provided email.
func (r *Repository) UpdatePassword(ctx context.Context, updatePassDb *model.UpdatePassDb) error {
	fmt.Println(updatePassDb)
	queryBuilder := squirrel.Update(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(consts.HashPasswordColumn, updatePassDb.HashPassword).
		Set(consts.UpdatedAtColumn, updatePassDb.UpdatedAt).
		Where(squirrel.Eq{consts.EmailColumn: updatePassDb.Email})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", zap.Error(err))
		return fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "UpdatePassword",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	fmt.Println(res)
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return fmt.Errorf("failed to update password: %v", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("user not found", zap.String("email", updatePassDb.Email))
		return fmt.Errorf("user not found")
	}

	return nil
}
