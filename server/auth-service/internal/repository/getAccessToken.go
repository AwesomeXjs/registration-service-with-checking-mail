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

func (r *Repository) GetAccessToken(ctx context.Context, userID string) (*model.AccessTokenInfo, error) {
	queryBuilder := squirrel.Select(consts.IdColumn, consts.RoleColumn).
		From(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{consts.IdColumn: userID}).
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

	var accessTokenInfo model.AccessTokenInfo
	err = r.db.DB().ScanOneContext(ctx, &accessTokenInfo, q, args...)
	if err != nil {
		logger.Error("failed to get access token from db", zap.Error(err))
		return nil, fmt.Errorf("failed to get access token from db: %v", err)
	}

	return &accessTokenInfo, nil
}
