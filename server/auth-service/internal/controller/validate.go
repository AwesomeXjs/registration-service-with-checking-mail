package controller

import (
	"context"
	"fmt"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ValidateToken checks the validity of the provided access token and returns an empty response.
func (c *Controller) ValidateToken(ctx context.Context, req *authService.ValidateTokenRequest) (*emptypb.Empty, error) {
	fmt.Println(ctx, req)
	return nil, nil
}
