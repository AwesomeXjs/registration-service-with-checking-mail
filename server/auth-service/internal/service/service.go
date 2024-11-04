package service

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

// IService defines the interface for the Service,
// representing business logic operations.
type IService interface {
	Registration(ctx context.Context, registrationRequest *model.UserInfo) (*model.AuthResponse, error)
	Login(ctx context.Context, loginRequest *model.LoginInfo) (*model.AuthResponse, error)
	GetAccessToken(ctx context.Context, refreshToken string) (*model.NewPairTokens, error)
}
