package auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	authClient *AuthClient
}

func New(authClient *AuthClient) *GRPCClient {
	return &GRPCClient{
		authClient: authClient,
	}
}

func (G GRPCClient) Login(ctx context.Context, request *authService.LoginRequest) (*authService.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCClient) Registration(ctx context.Context, request *authService.RegistrationRequest) (*authService.RegistrationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCClient) GetAccessToken(ctx context.Context, request *authService.GetAccessTokenRequest) (*authService.GetAccessTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCClient) UpdatePassword(ctx context.Context, request *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCClient) Validate(ctx context.Context, request *authService.ValidateTokenRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
