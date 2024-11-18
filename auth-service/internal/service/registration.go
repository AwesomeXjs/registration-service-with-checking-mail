package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Registration handles the registration of a new user, hashes their password,
// and generates access and refresh tokens upon successful registration.
func (s *Service) Registration(ctx context.Context, userInfo *model.UserInfo) (*model.AuthResponse, error) {
	HashedPassword, err := s.authHelper.HashPassword(userInfo.Password)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &model.InfoToDb{
		ID:           uuid.NewString(),
		Email:        userInfo.Email,
		HashPassword: HashedPassword,
		Role:         userInfo.Role,
	}

	userID, err := s.repo.Registration(ctx, user)
	if err != nil {
		logger.Error("failed to registration", zap.Error(err))
		return nil, err
	}

	// ТУТ ОТПРАВЛЯЕМ СНАЧАЛО СОЗДАНИЕ В СЕРВИСЕ ЮЗЕРОВ
	// ПОТОМ ОТПРАВЛЯЕМ ПИСЬМО В КАФКУ

	accessToken, err := s.authHelper.GenerateAccessToken(&model.AccessTokenInfo{
		ID:   userID,
		Role: user.Role,
	})
	if err != nil {
		logger.Error("failed to generate access token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.authHelper.GenerateRefreshToken(user.ID)
	if err != nil {
		logger.Error("failed to generate refresh token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
	}, nil
}
