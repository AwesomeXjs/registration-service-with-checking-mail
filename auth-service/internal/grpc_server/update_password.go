package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/validator"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePassword processes requests to update a user's password and returns an empty response.
func (c *GrpcServer) UpdatePassword(ctx context.Context, req *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {
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
