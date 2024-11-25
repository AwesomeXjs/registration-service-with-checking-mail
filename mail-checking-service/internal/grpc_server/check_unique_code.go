package grpc_server

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CheckUniqueCode handles the gRPC request to verify the uniqueness of a code.
// Returns an empty protobuf message on success or an error if the operation fails.
func (c *GrpcServer) CheckUniqueCode(ctx context.Context, req *mail_v1.CheckUniqueCodeRequest) (*emptypb.Empty, error) {
	result, err := c.redisClient.Get(ctx, req.GetEmail())
	if err != nil {
		logger.Warn("failed to get code from redis", zap.Error(err))
		return nil, fmt.Errorf("failed to get code from redis")
	}
	if result != req.GetCode() {
		logger.Error("code not found in redis", zap.Any("code", result))
		return nil, fmt.Errorf("code is invalid")
	}

	err = c.authClient.ValidateToken(ctx, req.GetAccessToken())
	if err != nil {
		logger.Error("failed to validate token", zap.Error(err))
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	err = c.authClient.ConfirmEmail(ctx, req.GetEmail())
	if err != nil {
		logger.Error("failed to confirm email", zap.Error(err))
		return nil, fmt.Errorf("failed to confirm email: %v", err)
	}

	return nil, nil
}
