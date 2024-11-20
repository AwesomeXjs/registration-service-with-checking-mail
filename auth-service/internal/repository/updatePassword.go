package repository

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// UpdatePassword updates the password of a user in the database based on their email address.
// The function builds an SQL update query to set the new hashed password and updated timestamp.
// If the query fails to build, an error is logged and returned. After executing the query,
// if no rows are affected, it logs a warning indicating the user was not found.
// If the update is successful, it attempts to delete the user’s cached data from Redis.
// Any errors encountered while interacting with Redis are logged but not returned.
func (r *Repository) UpdatePassword(ctx context.Context, updatePassDb *model.UpdatePassDb) error {
	queryBuilder := squirrel.Update(consts.TableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(consts.HashPasswordColumn, updatePassDb.HashPassword).
		Set(consts.UpdatedAtColumn, updatePassDb.UpdatedAt).
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

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	fmt.Println(res)
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return fmt.Errorf("failed to update password: %v", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("user not found", zap.String("email", updatePassDb.Email))
		return fmt.Errorf("user not found")
	}

	err = r.redisClient.Delete(ctx, updatePassDb.Email)
	if err != nil {
		logger.Error("failed to delete user from redis", zap.Error(err))
	}

	return nil
}
