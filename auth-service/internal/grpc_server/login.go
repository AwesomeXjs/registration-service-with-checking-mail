package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/validator"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// Login processes user login requests and returns a login response.
func (c *GrpcServer) Login(ctx context.Context, req *authService.LoginRequest) (*authService.LoginResponse, error) {
	logger.Debug("get user data in controller", zap.Any("req", req))

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
		logger.Error("failed to login", zap.Error(err))
		return nil, err
	}

	logger.Debug("new pair tokens: ", zap.Any("tokens", result))

	return converter.ToProtoFromLoginResponse(result), nil
}
