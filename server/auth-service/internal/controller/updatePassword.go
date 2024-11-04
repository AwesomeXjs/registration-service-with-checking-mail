package controller

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/validator"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePassword processes requests to update a user's password and returns an empty response.
func (c *Controller) UpdatePassword(ctx context.Context, req *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {
	err := validator.Validate(
		ctx,
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetNewPassword()))
	if err != nil {
		return nil, err
	}

	err = c.svc.UpdatePassword(ctx, converter.ToUpdatePassFromProto(req))
	if err != nil {
		return nil, fmt.Errorf("failed to update password: %v", err)
	}
	return nil, nil
}
