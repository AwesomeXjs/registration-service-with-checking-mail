package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// GetAccessToken retrieves the access token information for a given user ID.
// It first attempts to get the user's role from Redis.
// If the role is not found in Redis, it queries the database to retrieve it.
// The role is then cached in Redis for future access.
func (r *Repository) GetAccessToken(ctx context.Context, userID string) (*model.AccessTokenInfo, error) {
	logger.Debug("getting userID in repository", zap.String("USER_ID", userID))

	val, err := r.redisClient.Get(ctx, userID+" for role")
	if nil == err {
		return &model.AccessTokenInfo{
			ID:   userID,
			Role: val,
		}, nil
	}

	queryBuilder := squirrel.Select(consts.RoleColumn).
		From(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{consts.IDColumn: userID}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "GetAccessToken",
		QueryRaw: query,
	}

	var role string
	err = r.db.DB().ScanOneContext(ctx, &role, q, args...)
	if err != nil {
		logger.Error("failed to get role from db", zap.Error(err))
		return nil, fmt.Errorf("failed to get role from db: %v", err)
	}

	err = r.redisClient.Set(ctx, userID+" for role", role, 24*time.Hour)
	if err != nil {
		logger.Error("failed to set role in redis", zap.Error(err))
		return nil, fmt.Errorf("failed to set role in redis: %v", err)
	}

	logger.Debug("getting user role from database", zap.String("User role", role))

	return &model.AccessTokenInfo{
		ID:   userID,
		Role: role,
	}, nil
}
