package repository

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

func (r *Repository) Registration(ctx context.Context, infoToDb *model.InfoToDb) (string, error) {

	return "RETURNING ID", nil
}
