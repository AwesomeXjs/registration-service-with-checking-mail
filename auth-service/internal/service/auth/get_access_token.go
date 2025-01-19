package auth

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"go.uber.org/zap"
)

// GetAccessToken generates a new access token and refresh token based on the provided refresh token.
func (s *ServiceAuth) GetAccessToken(ctx context.Context, refreshToken string) (*model.NewPairTokens, error) {

	const mark = "Service.Auth.GetAccessToken"

	logger.Debug("getting refresh token on service", mark, zap.String("REFRESH_TOKEN", refreshToken))

	claims, err := s.AuthHelper.VerifyToken(refreshToken)
	if err != nil {
		logger.Error("failed to verify token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	info, err := s.Repo.Auth.GetAccessToken(ctx, claims.ID)
	if err != nil {
		logger.Error("failed to get access token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}
	access, err := s.AuthHelper.GenerateAccessToken(info)
	if err != nil {
		logger.Error("failed to generate access token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refresh, err := s.AuthHelper.GenerateRefreshToken(info.ID)
	if err != nil {
		logger.Error("failed to generate refresh token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &model.NewPairTokens{AccessToken: access, RefreshToken: refresh}, nil
}
