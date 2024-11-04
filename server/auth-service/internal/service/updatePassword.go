package service

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

func (s *Service) UpdatePassword(ctx context.Context, updatePassInfo *model.UpdatePassInfo) error {
	// захешировать пароль
	hashPassword, err := s.authHelper.HashPassword(updatePassInfo.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// сходить в базу данных обновить пароль
	err = s.repo.UpdatePassword(ctx, &model.UpdatePassDb{
		Email:        updatePassInfo.Email,
		HashPassword: hashPassword,
	})
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	// тут отослать в кафку сообщение о смене почты

	return nil
}
