package repository

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

// IRepository defines the interface for the Repository,
// representing data access methods.
type IRepository interface {
	Registration(ctx context.Context, registrationRequest *model.InfoToDb) (string, error)
	Login(ctx context.Context, email string) (*model.LoginResponse, error)
}
