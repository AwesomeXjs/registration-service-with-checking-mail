package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// Login retrieves the user's login information by email, checking Redis first for cached data.
func (r *Repository) Login(ctx context.Context, email string) (*model.LoginResponse, error) {
	logger.Debug("getting email in repository", zap.String("EMAIL", email))

	var loginResponse model.LoginResponse
	err := r.redisClient.GetObject(ctx, email, &loginResponse)
	if nil == err {
		logger.Info("found user in redis", zap.Any("user", loginResponse))
		return &loginResponse, nil
	}

	queryBuilder := squirrel.Select(IDColumn, HashPasswordColumn, RoleColumn).
		From(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{EmailColumn: email}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "Login",
		QueryRaw: query,
	}

	err = r.db.DB().ScanOneContext(ctx, &loginResponse, q, args...)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			logger.Debug("failed to get user from db", zap.Error(err))
			return nil, fmt.Errorf("user not found")
		}
		logger.Error("failed to get user from db", zap.Error(err))
		return nil, fmt.Errorf("failed to get user from db: %v", err)
	}

	err = r.redisClient.SetObject(ctx, email, loginResponse, 15*time.Minute)
	if err != nil {
		logger.Warn("failed to set user role in redis", zap.Error(err))
	}

	return &loginResponse, nil
}
