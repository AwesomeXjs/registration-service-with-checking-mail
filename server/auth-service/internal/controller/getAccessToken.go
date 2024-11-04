package controller

import (
	"context"
	"fmt"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

// GetAccessToken handles the request for generating an access token.
func (c *Controller) GetAccessToken(ctx context.Context, req *authService.GetAccessTokenRequest) (*authService.GetAccessTokenResponse, error) {
	fmt.Println(ctx, req)
	return nil, nil
}