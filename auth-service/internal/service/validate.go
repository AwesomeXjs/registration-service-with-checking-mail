package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

// ValidateToken verifies the access token and returns true if valid.
func (s *Service) ValidateToken(_ context.Context, accessToken string) (bool, error) {
	logger.Debug("get access token in service", zap.String("ACCESS_TOKEN", accessToken))

	_, err := s.authHelper.VerifyToken(accessToken)
	if err != nil {
		logger.Error("failed to verify token", zap.Error(err))
		return false, fmt.Errorf("failed to verify token: %v", err)
	}
	return true, nil
}
