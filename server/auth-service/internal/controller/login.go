package controller

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/validator"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// Login processes user login requests and returns a login response.
func (c *Controller) Login(ctx context.Context, req *authService.LoginRequest) (*authService.LoginResponse, error) {
	err := validator.Validate(
		ctx,
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetPassword()))
	if err != nil {
		logger.Error(err.Error(), zap.Any("validate", err.Error()))
		return nil, err
	}

	result, err := c.svc.Login(ctx, converter.ToLoginInfoFromProto(req))
	if err != nil {
		return nil, err
	}

	return converter.ToProtoFromLoginResponse(result), nil
}
