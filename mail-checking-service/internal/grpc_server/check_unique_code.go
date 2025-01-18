package grpc_server

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CheckUniqueCode handles the gRPC request to verify the uniqueness of a code.
// Returns an empty protobuf message on success or an error if the operation fails.
func (c *GrpcServer) CheckUniqueCode(ctx context.Context, req *mail_v1.CheckUniqueCodeRequest) (*emptypb.Empty, error) {

	const mark = "GrpcServer.CheckUniqueCode"

	result, err := c.redisClient.Get(ctx, req.GetEmail())
	if err != nil {
		logger.Warn("failed to get code from redis", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to get code from redis")
	}
	if result != req.GetCode() {
		logger.Error("code not found in redis", mark, zap.Any("code", result))
		return nil, fmt.Errorf("code is invalid")
	}

	span, contextWithTraceValidateToken := opentracing.StartSpanFromContext(ctx, "validate token")
	defer span.Finish()

	span.SetTag("token", req.GetAccessToken())

	err = c.authClient.ValidateToken(contextWithTraceValidateToken, req.GetAccessToken())
	if err != nil {
		logger.Error("failed to validate token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	span, contextWithTrace := opentracing.StartSpanFromContext(ctx, "confirm mail")
	defer span.Finish()

	span.SetTag("email", req.GetEmail())

	err = c.authClient.ConfirmEmail(contextWithTrace, req.GetEmail())
	if err != nil {
		logger.Error("failed to confirm email", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to confirm email: %v", err)
	}

	return nil, nil
}
