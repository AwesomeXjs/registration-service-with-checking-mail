package controller

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/validator"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePassword processes requests to update a user's password and returns an empty response.
func (c *Controller) UpdatePassword(ctx context.Context, req *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {
	logger.Debug("update password", zap.Any("req", req))

	err := validator.Validate(
		ctx,
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetNewPassword()))
	if err != nil {
		logger.Error("failed to validate", zap.Error(err))
		return nil, err
	}

	err = c.svc.UpdatePassword(ctx, converter.ToUpdatePassFromProto(req))
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return nil, err
	}
	return nil, nil
}
