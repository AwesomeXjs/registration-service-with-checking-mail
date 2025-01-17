package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/validator"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePassword processes requests to update a user's password and returns an empty response.
func (c *GrpcServer) UpdatePassword(ctx context.Context,
	req *authService.UpdatePasswordRequest) (*emptypb.Empty, error) {

	const mark = "GrpcServer.UpdatePassword"

	logger.Debug("update password", mark, zap.Any("req", req))

	err := validator.Validate(
		ctx,
		validator.ValidateEmail(req.GetEmail()),
		validator.ValidatePassword(req.GetNewPassword()))
	if err != nil {
		logger.Error("failed to validate", mark, zap.Error(err))
		return nil, err
	}

	err = c.svc.Auth.UpdatePassword(ctx, converter.ToUpdatePassFromProto(req))
	if err != nil {
		logger.Error("failed to update password", mark, zap.Error(err))
		return nil, err
	}
	return nil, nil
}
