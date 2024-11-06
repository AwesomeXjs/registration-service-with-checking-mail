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

func (g *GRPCClient) Registration(ctx context.Context, in *authService.RegistrationRequest, opts ...grpc.CallOption) (*authService.RegistrationResponse, error) {
	return g.authClient.Registration(ctx, in, opts...)
}

func (g *GRPCClient) Login(ctx context.Context, in *authService.LoginRequest, opts ...grpc.CallOption) (*authService.LoginResponse, error) {
	return g.authClient.Login(ctx, in, opts...)
}

func (g *GRPCClient) ValidateToken(ctx context.Context, in *authService.ValidateTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.authClient.ValidateToken(ctx, in, opts...)
}

func (g *GRPCClient) GetAccessToken(ctx context.Context, in *authService.GetAccessTokenRequest, opts ...grpc.CallOption) (*authService.GetAccessTokenResponse, error) {
	return g.authClient.GetAccessToken(ctx, in, opts...)
}

func (g *GRPCClient) UpdatePassword(ctx context.Context, in *authService.UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.authClient.UpdatePassword(ctx, in, opts...)
}
