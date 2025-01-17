package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// Registration inserts a new user into the database and returns the user ID.
func (a *RepositoryAuth) Registration(ctx context.Context, infoToDb *model.InfoToDb) (int, error) {

	const mark = "Repository.Auth.Registration"

	builderInsert := squirrel.Insert(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(EmailColumn, HashPasswordColumn, RoleColumn).
		Values(infoToDb.Email, infoToDb.HashPassword, infoToDb.Role).
		Suffix(ReturningID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return 0, fmt.Errorf("failed to build insert query: %v", err)
	}

	q := db.Query{
		Name:     "Registration",
		QueryRaw: query,
	}

	var ID int
	err = a.db.DB().QueryRowContext(ctx, q, args...).Scan(&ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			logger.Warn("email already registered", mark, zap.Error(err))
			return 0, fmt.Errorf("email already registered")
		}

		logger.Error("failed to register user", mark, zap.Error(err))
		return 0, fmt.Errorf("failed to register user: %v", err)
	}

	return ID, nil
}
