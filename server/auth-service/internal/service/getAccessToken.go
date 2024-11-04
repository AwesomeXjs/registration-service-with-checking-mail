package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

func (s *Service) GetAccessToken(ctx context.Context, refreshToken string) (*model.NewPairTokens, error) {
	claims, err := s.authHelper.VerifyToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	fmt.Println(claims.ID)
	info, err := s.repo.GetAccessToken(ctx, claims.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}
	access, err := s.authHelper.GenerateAccessToken(info)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refresh, err := s.authHelper.GenerateRefreshToken(info.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return &model.NewPairTokens{AccessToken: access, RefreshToken: refresh}, nil
}
