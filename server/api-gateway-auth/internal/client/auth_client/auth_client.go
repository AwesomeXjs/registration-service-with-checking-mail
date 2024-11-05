package auth_client

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthClient interface {
	Login(ctx context.Context, request *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error)
	Registration(ctx context.Context, request *auth_v1.RegistrationRequest) (*auth_v1.RegistrationResponse, error)
	GetAccessToken(ctx context.Context, request *auth_v1.GetAccessTokenRequest) (*auth_v1.GetAccessTokenResponse, error)
	UpdatePassword(ctx context.Context, request *auth_v1.UpdatePasswordRequest) (*emptypb.Empty, error)
	Validate(ctx context.Context, request *auth_v1.ValidateTokenRequest) (*emptypb.Empty, error)
}
