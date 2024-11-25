package grpc_auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IAuthClient interface {
	ConfirmEmail(ctx context.Context, in *authService.ConfirmEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
