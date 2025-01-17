package auth

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"go.uber.org/zap"
)

// Login authenticates a user with the provided login information and generates access and refresh tokens.
func (s *ServiceAuth) Login(ctx context.Context, loginInfo *model.LoginInfo) (*model.AuthResponse, error) {

	const mark = "Service.Auth.Login"

	logger.Debug("get data in service", mark, zap.Any("loginInfo", loginInfo))

	result, err := s.repo.Auth.Login(ctx, loginInfo.Email)
	if err != nil {
		logger.Error("failed to login", mark, zap.Error(err))
		return nil, err
	}

	ok := s.authHelper.ValidatePassword(result.HashPassword, loginInfo.Password)
	if !ok {
		logger.Error("invalid password", mark, zap.Error(err))
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, err := s.authHelper.GenerateAccessToken(&model.AccessTokenInfo{
		ID:   result.UserID,
		Role: result.Role,
	})
	if err != nil {
		logger.Error("failed to generate access token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.authHelper.GenerateRefreshToken(result.UserID)
	if err != nil {
		logger.Error("failed to generate refresh token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	logger.Debug("new pair tokens in service: ", mark, zap.Any("tokens", result))

	return converter.ToModelAuthResponse(accessToken, refreshToken, result.UserID), nil
}
