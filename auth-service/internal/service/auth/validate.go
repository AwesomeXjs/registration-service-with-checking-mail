package auth

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"go.uber.org/zap"
)

// ValidateToken verifies the access token and returns true if valid.
func (s *ServiceAuth) ValidateToken(_ context.Context, accessToken string) (bool, error) {

	const mark = "Service.Auth.ValidateToken"

	logger.Debug("get access token in service", mark, zap.String("ACCESS_TOKEN", accessToken))

	_, err := s.authHelper.VerifyToken(accessToken)
	if err != nil {
		logger.Error("failed to verify token", mark, zap.Error(err))
		return false, fmt.Errorf("failed to verify token: %v", err)
	}
	return true, nil
}
