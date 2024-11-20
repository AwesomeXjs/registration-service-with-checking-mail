package controller

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// GetAccessToken handles the request for generating an access token.
func (c *Controller) GetAccessToken(ctx context.Context, req *authService.GetAccessTokenRequest) (*authService.GetAccessTokenResponse, error) {
	logger.Debug("getting refresh token", zap.String("REFRESH_TOKEN", req.GetRefreshToken()))

	res, err := c.svc.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		return nil, err
	}

	logger.Debug("new pair tokens: ", zap.Any("tokens", res))
	return converter.ToProtoFromNewPairTokens(res), nil
}
