package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"go.uber.org/zap"
)

var (
	topicRegistration = "registration"
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
		Email:        userInfo.Email,
		HashPassword: HashedPassword,
		Role:         userInfo.Role,
	}

	userID, err := s.repo.Registration(ctx, user)
	if err != nil {
		logger.Error("failed to registration", zap.Error(err))
		return nil, err
	}

	err = s.kafkaProducer.Produce(user.Email, topicRegistration, user.ID, time.Now())
	if err != nil {
		logger.Error("failed to produce message", zap.Error(err))
		return nil, fmt.Errorf("failed to produce message: %v", err)
	}

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
