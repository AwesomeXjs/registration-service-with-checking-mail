package auth

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"go.uber.org/zap"
)

// UpdatePassword hashes the new password and updates the user's password in the database.
func (s *ServiceAuth) UpdatePassword(ctx context.Context, updatePassInfo *model.UpdatePassInfo) error {

	const mark = "Service.Auth.UpdatePassword"

	logger.Debug("get data in service", mark, zap.Any("updatePassInfo", updatePassInfo))

	hashPassword, err := s.authHelper.HashPassword(updatePassInfo.NewPassword)
	if err != nil {
		logger.Error("failed to hash password", mark, zap.Error(err))
		return fmt.Errorf("failed to hash password: %v", err)
	}
	err = s.repo.Auth.UpdatePassword(ctx, converter.FromUpdatePassInfoToDbPassInfo(updatePassInfo, hashPassword))
	if err != nil {
		logger.Error("failed to update password", mark, zap.Error(err))
		return fmt.Errorf("failed to update password: %v", err)
	}

	logger.Debug("password updated", mark, zap.Any("updatePassInfo", updatePassInfo))
	return nil
}
