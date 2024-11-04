package controller

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ValidateToken checks the validity of the provided access token and returns an empty response.
func (c *Controller) ValidateToken(ctx context.Context, req *authService.ValidateTokenRequest) (*emptypb.Empty, error) {
	_, err := c.svc.ValidateToken(ctx, req.GetAccessToken())
	if err != nil {
		logger.Error(err.Error(), zap.Any("req", req))
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}
	return nil, nil
}
