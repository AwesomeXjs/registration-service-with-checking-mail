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

func (r *Repository) Registration(ctx context.Context, infoToDb *model.InfoToDb) (string, error) {
	builderInsert := squirrel.Insert(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(consts.IdColumn, consts.EmailColumn, consts.HashPasswordColumn, consts.RoleColumn).
		Values(infoToDb.ID, infoToDb.Email, infoToDb.HashPassword, infoToDb.Role).
		Suffix(consts.ReturningID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", zap.Error(err))
		return "", fmt.Errorf("failed to build insert query: %v", err)
	}

	q := db.Query{
		Name:     "Registration",
		QueryRaw: query,
	}

	var ID string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&ID)
	if err != nil {
		logger.Error("failed to get user from db", zap.Error(err))
		return "", fmt.Errorf("failed to get user from db: %v", err)
	}

	return ID, nil
}
