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

func (r *Repository) UpdatePassword(ctx context.Context, updatePassDb *model.UpdatePassDb) error {
	queryBuilder := squirrel.Update(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(consts.HashPasswordColumn, updatePassDb.HashPassword).
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

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}
