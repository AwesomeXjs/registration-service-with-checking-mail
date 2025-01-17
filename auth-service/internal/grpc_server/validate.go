package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ValidateToken checks the validity of the provided access token and returns an empty response.
func (c *GrpcServer) ValidateToken(ctx context.Context,
	req *authService.ValidateTokenRequest) (*emptypb.Empty, error) {

	const mark = "GrpcServer.ValidateToken"

	_, err := c.svc.Auth.ValidateToken(ctx, req.GetAccessToken())
	if err != nil {
		logger.Error("failed to validate token", mark, zap.Any("req", req))
		return nil, err
	}
	return nil, nil
}
