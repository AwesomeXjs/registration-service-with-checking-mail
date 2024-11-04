package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

// Login authenticates a user with the provided login information and generates access and refresh tokens.
func (s *Service) Login(ctx context.Context, loginInfo *model.LoginInfo) (*model.AuthResponse, error) {
	result, err := s.repo.Login(ctx, loginInfo.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	ok := s.authHelper.ValidatePassword(result.HashPassword, loginInfo.Password)
	if !ok {
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, err := s.authHelper.GenerateAccessToken(&model.AccessTokenInfo{
		ID:   result.UserID,
		Role: result.Role,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.authHelper.GenerateRefreshToken(result.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       result.UserID,
	}, nil
}
