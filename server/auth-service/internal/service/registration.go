package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Service) Registration(ctx context.Context, userInfo *model.UserInfo) (model.RegistrationResponse, error) {
	HashedPassword, err := s.authHelper.HashPassword(userInfo.Password)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return model.RegistrationResponse{}, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &model.InfoToDb{
		ID:           uuid.NewString(),
		Email:        userInfo.Email,
		HashPassword: HashedPassword,
		Role:         userInfo.Role,
	}

	res, err := s.repo.Registration(ctx, user)
	if err != nil {
		logger.Error("failed to registration", zap.Error(err))
		return model.RegistrationResponse{}, fmt.Errorf("failed to registration: %v", err)
	}

	accessToken, err := s.authHelper.GenerateAccessToken(user)
	if err != nil {
		logger.Error("failed to generate access token", zap.Error(err))
		return model.RegistrationResponse{}, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.authHelper.GenerateRefreshToken(user)
	if err != nil {
		logger.Error("failed to generate refresh token", zap.Error(err))
		return model.RegistrationResponse{}, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return model.RegistrationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       res,
	}, nil
}
