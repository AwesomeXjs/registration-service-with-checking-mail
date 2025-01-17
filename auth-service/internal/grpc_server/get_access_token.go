package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// GetAccessToken handles the request for generating an access token.
func (c *GrpcServer) GetAccessToken(ctx context.Context,
	req *authService.GetAccessTokenRequest) (*authService.GetAccessTokenResponse, error) {
	const mark = "GrpcServer.GetAccessToken"

	logger.Debug("getting refresh token", mark, zap.String("REFRESH_TOKEN", req.GetRefreshToken()))

	res, err := c.svc.Auth.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		logger.Error("failed to get access token", mark, zap.Error(err))
		return nil, err
	}

	logger.Debug("new pair tokens: ", mark, zap.Any("tokens", res))
	return converter.ToProtoFromNewPairTokens(res), nil
}
