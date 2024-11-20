package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

// Login authenticates a user with the provided login information and generates access and refresh tokens.
func (s *Service) Login(ctx context.Context, loginInfo *model.LoginInfo) (*model.AuthResponse, error) {
	logger.Debug("get data in service", zap.Any("loginInfo", loginInfo))

	result, err := s.repo.Login(ctx, loginInfo.Email)
	if err != nil {
		logger.Error("failed to login", zap.Error(err))
		return nil, err
	}

	ok := s.authHelper.ValidatePassword(result.HashPassword, loginInfo.Password)
	if !ok {
		logger.Error("invalid password", zap.Error(err))
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, err := s.authHelper.GenerateAccessToken(&model.AccessTokenInfo{
		ID:   result.UserID,
		Role: result.Role,
	})
	if err != nil {
		logger.Error("failed to generate access token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.authHelper.GenerateRefreshToken(result.UserID)
	if err != nil {
		logger.Error("failed to generate refresh token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	logger.Debug("new pair tokens in service: ", zap.Any("tokens", result))

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       result.UserID,
	}, nil
}
