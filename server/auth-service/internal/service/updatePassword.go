package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

// UpdatePassword hashes the new password and updates the user's password in the database.
func (s *Service) UpdatePassword(ctx context.Context, updatePassInfo *model.UpdatePassInfo) error {
	hashPassword, err := s.authHelper.HashPassword(updatePassInfo.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	err = s.repo.UpdatePassword(ctx, &model.UpdatePassDb{
		Email:        updatePassInfo.Email,
		HashPassword: hashPassword,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}
