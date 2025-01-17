package auth

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

// ConfirmEmail updates the user's email verification status to confirmed in the database.
// It accepts a context and the user's email as parameters and returns an error if any issues occur.
func (a *RepositoryAuth) ConfirmEmail(ctx context.Context, email string) error {

	const mark = "Repository.Auth.ConfirmEmail"

	queryBuilder := squirrel.Update(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(Verification, true).
		Where(squirrel.Eq{EmailColumn: email})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query", mark, zap.Error(err))
		return fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "Confirm Email",
		QueryRaw: query,
	}

	res, err := a.db.DB().ExecContext(ctx, q, args...)
	fmt.Println(res)
	if err != nil {
		logger.Error("failed to confirm email", mark, zap.Error(err))
		return fmt.Errorf("failed to confirm email: %v", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Error("user not found", mark, zap.Error(err))
		return fmt.Errorf("user not found")
	}

	return nil
}
