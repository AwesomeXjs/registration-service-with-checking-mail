package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"

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

	userID, err := s.repo.Registration(ctx, converter.FromUserInfoToDbModel(userInfo, HashedPassword))
	if err != nil {
		logger.Error("failed to registration", zap.Error(err))
		return nil, err
	}

	err = s.kafkaProducer.Produce(userInfo.Email, topicRegistration, strconv.Itoa(userID))
	if err != nil {
		logger.Error("failed to produce message", zap.Error(err))
		return nil, fmt.Errorf("failed to produce message: %v", err)
	}

	accessToken, err := s.authHelper.GenerateAccessToken(&model.AccessTokenInfo{
		ID:   userID,
		Role: userInfo.Role,
	})
	if err != nil {
		logger.Error("failed to generate access token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.authHelper.GenerateRefreshToken(userID)
	if err != nil {
		logger.Error("failed to generate refresh token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       int64(userID),
	}, nil
}
