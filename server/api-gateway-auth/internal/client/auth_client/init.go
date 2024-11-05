package auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	authClient AuthClient
}

func New(authClient AuthClient) *GRPCClient {
	return &GRPCClient{
		authClient: authClient,
	}
}

func (G *GRPCClient) Registration(ctx context.Context, in *authService.RegistrationRequest, opts ...grpc.CallOption) (*authService.RegistrationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GRPCClient) Login(ctx context.Context, in *authService.LoginRequest, opts ...grpc.CallOption) (*authService.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GRPCClient) ValidateToken(ctx context.Context, in *authService.ValidateTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GRPCClient) GetAccessToken(ctx context.Context, in *authService.GetAccessTokenRequest, opts ...grpc.CallOption) (*authService.GetAccessTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GRPCClient) UpdatePassword(ctx context.Context, in *authService.UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
