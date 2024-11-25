package grpc_auth_client

import (
	"context"
)

type IAuthClient interface {
	ConfirmEmail(ctx context.Context, email string) error
	ValidateToken(ctx context.Context, accessToken string) error
}
