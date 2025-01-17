package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// GetAccessToken retrieves the access token information for a given user ID.
// It first attempts to get the user's role from Redis.
// If the role is not found in Redis, it queries the database to retrieve it.
// The role is then cached in Redis for future access.
func (a *RepositoryAuth) GetAccessToken(ctx context.Context, userID int) (*model.AccessTokenInfo, error) {

	const mark = "Repository.Auth.GetAccessToken"

	val, err := a.redisClient.Get(ctx, strconv.Itoa(userID)+" for role")
	if nil == err {
		return &model.AccessTokenInfo{
			ID:   userID,
			Role: val,
		}, nil
	}

	queryBuilder := squirrel.Select(RoleColumn).
		From(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{IDColumn: userID}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "GetAccessToken",
		QueryRaw: query,
	}

	var role string
	err = a.db.DB().ScanOneContext(ctx, &role, q, args...)
	if err != nil {
		logger.Error("failed to get role from db", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to get role from db: %v", err)
	}

	err = a.redisClient.Set(ctx, strconv.Itoa(userID)+" for role", role, 24*time.Hour)
	if err != nil {
		logger.Error("failed to set role in redis", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to set role in redis: %v", err)
	}

	logger.Debug("getting user role from database", mark, zap.String("User role", role))

	return &model.AccessTokenInfo{
		ID:   userID,
		Role: role,
	}, nil
}
