package controller

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

// Registration handles user registration requests and returns a registration response.
func (c *Controller) Registration(ctx context.Context, req *authService.RegistrationRequest) (*authService.RegistrationResponse, error) {

	return nil, nil
}
