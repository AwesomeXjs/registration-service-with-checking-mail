package controller

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// GetAccessToken handles the request for generating an access token.
func (c *Controller) GetAccessToken(ctx context.Context, req *authService.GetAccessTokenRequest) (*authService.GetAccessTokenResponse, error) {
	res, err := c.svc.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		logger.Error(err.Error(), zap.Any("req", req))
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}
	return converter.ToProtoFromNewPairTokens(res), nil
}
