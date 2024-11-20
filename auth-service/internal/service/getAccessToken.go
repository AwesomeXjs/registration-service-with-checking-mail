package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

// GetAccessToken generates a new access token and refresh token based on the provided refresh token.
func (s *Service) GetAccessToken(ctx context.Context, refreshToken string) (*model.NewPairTokens, error) {
	logger.Debug("getting refresh token on service", zap.String("REFRESH_TOKEN", refreshToken))

	claims, err := s.authHelper.VerifyToken(refreshToken)
	if err != nil {
		logger.Error("failed to verify token", zap.Error(err))
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	info, err := s.repo.GetAccessToken(ctx, claims.ID)
	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}
	access, err := s.authHelper.GenerateAccessToken(info)
	if err != nil {
		logger.Error("failed to generate access token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refresh, err := s.authHelper.GenerateRefreshToken(info.ID)
	if err != nil {
		logger.Error("failed to generate refresh token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &model.NewPairTokens{AccessToken: access, RefreshToken: refresh}, nil
}
