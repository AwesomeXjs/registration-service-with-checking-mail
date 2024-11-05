package controller

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/validator"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// Registration handles user registration requests and returns a registration response.
func (c *Controller) Registration(ctx context.Context, req *authService.RegistrationRequest) (*authService.RegistrationResponse, error) {
	logger.Debug("registration", zap.Any("req", req))

	err := validator.Validate(ctx,
		validator.ValidateName(req.GetName()),
		validator.ValidateRole(req.GetRole()),
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetPassword()),
		validator.ValidateSurname(req.GetSurname()))
	if err != nil {
		logger.Info("failed to validate", zap.Error(err))
		return nil, err
	}

	res, err := c.svc.Registration(ctx, converter.ToInfoFromProto(req))
	if err != nil {
		logger.Error(err.Error(), zap.Any("req", req))
		return nil, fmt.Errorf("failed to registration: %v", err)
	}

	logger.Debug("new pair tokens: ", zap.Any("tokens", res))

	return converter.ToProtoFromRegResponse(res), nil
}
