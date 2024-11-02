package controller

import (
	"context"
	"fmt"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

// Login processes user login requests and returns a login response.
func (c *Controller) Login(ctx context.Context, req *authService.LoginRequest) (*authService.LoginResponse, error) {
	fmt.Println(ctx, req)
	return nil, nil
}
