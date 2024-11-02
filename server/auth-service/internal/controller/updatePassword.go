package controller

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Controller) UpdatePassword(ctx context.Context, req *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {

	return nil, nil
}
