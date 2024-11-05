package auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthClient interface {
	Login(ctx context.Context, in *authService.LoginRequest, opts ...grpc.CallOption) (*authService.LoginResponse, error)
	Registration(ctx context.Context, in *authService.RegistrationRequest, opts ...grpc.CallOption) (*authService.RegistrationResponse, error)
	UpdatePassword(ctx context.Context, in *authService.UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAccessToken(ctx context.Context, in *authService.GetAccessTokenRequest, opts ...grpc.CallOption) (*authService.GetAccessTokenResponse, error)
	ValidateToken(ctx context.Context, in *authService.ValidateTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
