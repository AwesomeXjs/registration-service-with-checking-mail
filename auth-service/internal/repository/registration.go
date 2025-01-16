package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// Registration inserts a new user into the database and returns the user ID.
func (r *Repository) Registration(ctx context.Context, infoToDb *model.InfoToDb) (string, error) {
	builderInsert := squirrel.Insert(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(EmailColumn, HashPasswordColumn, RoleColumn).
		Values(infoToDb.Email, infoToDb.HashPassword, infoToDb.Role).
		Suffix(ReturningID)

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
