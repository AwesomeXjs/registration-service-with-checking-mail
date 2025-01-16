package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"go.uber.org/zap"
)

// UpdatePassword hashes the new password and updates the user's password in the database.
func (s *Service) UpdatePassword(ctx context.Context, updatePassInfo *model.UpdatePassInfo) error {
	logger.Debug("get data in service", zap.Any("updatePassInfo", updatePassInfo))

	hashPassword, err := s.authHelper.HashPassword(updatePassInfo.NewPassword)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return fmt.Errorf("failed to hash password: %v", err)
	}
	err = s.repo.UpdatePassword(ctx, converter.FromUpdatePassInfoToDbPassInfo(updatePassInfo, hashPassword))
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return fmt.Errorf("failed to update password: %v", err)
	}

	logger.Debug("password updated", zap.Any("updatePassInfo", updatePassInfo))
	return nil
}
