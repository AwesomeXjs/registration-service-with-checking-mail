package grpc_auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthClient struct {
	authClient IAuthClient
}

func NewAuthClient(authClient IAuthClient) IAuthClient {
	return &AuthClient{
		authClient: authClient,
	}
}

func (a *AuthClient) ConfirmEmail(ctx context.Context, in *authService.ConfirmEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return a.authClient.ConfirmEmail(ctx, in, opts...)
}
