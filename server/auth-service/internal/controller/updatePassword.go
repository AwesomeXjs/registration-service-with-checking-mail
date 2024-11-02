package controller

import (
	"context"
	"fmt"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePassword processes requests to update a user's password and returns an empty response.
func (c *Controller) UpdatePassword(ctx context.Context, req *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {
	fmt.Println(ctx, req)
	return nil, nil
}
