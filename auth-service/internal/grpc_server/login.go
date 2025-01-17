package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/validator"
	"go.uber.org/zap"
)

// Login processes user login requests and returns a login response.
func (c *GrpcServer) Login(ctx context.Context,
	req *authService.LoginRequest) (*authService.LoginResponse, error) {

	const mark = "GrpcServer.Login"

	logger.Debug("get user data in controller", mark, zap.Any("req", req))

	err := Validate(ctx, req)
	if err != nil {
		logger.Error("failed to validate", mark, zap.Error(err))
		return nil, err
	}

	result, err := c.svc.Auth.Login(ctx, converter.ToLoginInfoFromProto(req))
	if err != nil {
		logger.Error("failed to login", mark, zap.Error(err))
		return nil, err
	}

	logger.Debug("new pair tokens: ", mark, zap.Any("tokens", result))

	return converter.ToProtoFromLoginResponse(result), nil
}

// Validate checks the validity of the provided login request.
func Validate(ctx context.Context, req *authService.LoginRequest) error {
	err := validator.Validate(ctx,
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetPassword()))
	if err != nil {
		return err
	}
	return nil
}
