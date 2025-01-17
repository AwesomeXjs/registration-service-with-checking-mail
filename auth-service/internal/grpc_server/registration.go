package grpc_server

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/validator"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// Registration handles user registration requests and returns a registration response.
func (c *GrpcServer) Registration(ctx context.Context,
	req *authService.RegistrationRequest) (*authService.RegistrationResponse, error) {

	const mark = "GrpcServer.Registration"

	logger.Debug("registration", mark, zap.Any("req", req))

	err := ValidateUserData(ctx, req)
	if err != nil {
		logger.Info("failed to validate", mark, zap.Error(err))
		return nil, err
	}

	res, err := c.svc.Auth.Registration(ctx, converter.ToInfoFromProto(req))
	if err != nil {
		logger.Error(err.Error(), mark, zap.Any("req", req))
		return nil, fmt.Errorf("failed to registration: %v", err)
	}

	logger.Debug("new pair tokens: ", mark, zap.Any("tokens", res))

	return converter.ToProtoFromRegResponse(res), nil
}

// ValidateUserData validates the user registration request data.
func ValidateUserData(ctx context.Context, req *authService.RegistrationRequest) error {
	return validator.Validate(ctx,
		validator.ValidateName(req.GetName()),
		validator.ValidateRole(req.GetRole()),
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetPassword()),
		validator.ValidateSurname(req.GetSurname()),
	)
}
