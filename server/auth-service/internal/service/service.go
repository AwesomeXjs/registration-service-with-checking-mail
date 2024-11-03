package service

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

// IService defines the interface for the Service,
// representing business logic operations.
type IService interface {
	Registration(ctx context.Context, registrationRequest *model.UserInfo) (model.RegistrationResponse, error)
}
