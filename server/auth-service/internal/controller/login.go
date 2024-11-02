package controller

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

func (c *Controller) Login(ctx context.Context, req *authService.LoginRequest) (*authService.LoginResponse, error) {

	return nil, nil
}
