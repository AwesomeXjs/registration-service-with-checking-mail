package controller

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

func (c *Controller) GetAccessToken(ctx context.Context, req *authService.GetAccessTokenRequest) (*authService.GetAccessTokenResponse, error) {

	return nil, nil
}
