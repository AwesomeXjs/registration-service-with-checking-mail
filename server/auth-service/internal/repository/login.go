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

// Login retrieves the user's login information by email, checking Redis first for cached data.
func (r *Repository) Login(ctx context.Context, email string) (*model.LoginResponse, error) {

	queryBuilder := squirrel.Select(consts.IDColumn, consts.HashPasswordColumn, consts.RoleColumn).
		From(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{consts.EmailColumn: email}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	fmt.Println(query)

	q := db.Query{
		Name:     "Login",
		QueryRaw: query,
	}
	var loginResponse model.LoginResponse
	err = r.db.DB().ScanOneContext(ctx, &loginResponse, q, args...)
	if err != nil {
		logger.Error("failed to get user from db", zap.Error(err))
		return nil, fmt.Errorf("failed to get user from db: %v", err)
	}

	return &loginResponse, nil
}
