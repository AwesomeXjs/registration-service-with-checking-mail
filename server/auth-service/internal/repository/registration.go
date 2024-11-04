package repository

import (
	"context"
	"fmt"
	"strings"

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
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			logger.Warn("email already registered", zap.Error(err))
			return "", fmt.Errorf("email already registered")
		}

		logger.Error("failed to register user", zap.Error(err))
		return "", fmt.Errorf("failed to register user: %v", err)
	}

	return ID, nil
}
