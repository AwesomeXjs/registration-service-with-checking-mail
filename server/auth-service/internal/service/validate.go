package service

import (
	"context"
	"fmt"
)

// ValidateToken verifies the access token and returns true if valid.
func (s *Service) ValidateToken(_ context.Context, accessToken string) (bool, error) {
	_, err := s.authHelper.VerifyToken(accessToken)
	if err != nil {
		return false, fmt.Errorf("failed to verify token: %v", err)
	}
	return true, nil
}
