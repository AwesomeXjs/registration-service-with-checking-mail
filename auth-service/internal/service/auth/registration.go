package auth

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/goccy/go-json"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"go.uber.org/zap"
)

var (
	topicRegistration = "registration"
)

// Registration handles the registration of a new user, hashes their password,
// and generates access and refresh tokens upon successful registration.
func (s *ServiceAuth) Registration(ctx context.Context, userInfo *model.UserInfo) (*model.AuthResponse, error) {

	const mark = "Service.Auth.Registration"

	HashedPassword, err := s.AuthHelper.HashPassword(userInfo.Password)
	if err != nil {
		logger.Error("failed to hash password", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	var UserID int
	err = s.Tx.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		UserID, errTx = s.Repo.Auth.Registration(ctx, converter.FromUserInfoToDbModel(userInfo, HashedPassword))
		if errTx != nil {
			logger.Error("failed to registration", mark, zap.Error(err))
			return errTx
		}

		payload, errTx := json.Marshal(converter.ToModelPayload(topicRegistration, UserID, userInfo))
		if errTx != nil {
			logger.Error("failed to marshal", mark, zap.Error(err))
			return errTx
		}

		errTx = s.Repo.Event.SendEvent(ctx, converter.ToModelSendEvent(topicRegistration, payload))
		if errTx != nil {
			logger.Error("failed to send event", mark, zap.Error(err))
			return errTx
		}

		return nil
	})

	if err != nil {
		logger.Error("failed to registration", mark, zap.Error(err))
		return nil, err
	}

	accessToken, err := s.AuthHelper.GenerateAccessToken(converter.ToModelAccessTokenInfo(UserID, userInfo))
	if err != nil {
		logger.Error("failed to generate access token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := s.AuthHelper.GenerateRefreshToken(UserID)
	if err != nil {
		logger.Error("failed to generate refresh token", mark, zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return converter.ToModelAuthResponse(accessToken, refreshToken, UserID), nil
}
