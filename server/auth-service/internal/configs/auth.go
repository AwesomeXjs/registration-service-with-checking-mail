package configs

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
)

type AuthConfig struct {
	secretKey            []byte
	refreshTokenDuration time.Duration
	accessTokenDuration  time.Duration
}

func NewAuthConfig() (*AuthConfig, error) {
	secretKey := os.Getenv(consts.SecretKey)
	if len(secretKey) == 0 {
		logger.Error("failed to get secret key", zap.String("secret key", consts.SecretKey))
		return nil, fmt.Errorf("env %s is empty", consts.SecretKey)
	}

	refreshTokenDuration, err := strconv.Atoi(os.Getenv(consts.RefreshTokenDuration))
	if err != nil {
		logger.Error("failed to get refresh token duration", zap.Error(err))
		return nil, fmt.Errorf("env %s is empty", consts.RefreshTokenDuration)
	}
	accessTokenDuration, err := strconv.Atoi(os.Getenv(consts.AccessTokenDuration))
	if err != nil {
		logger.Error("failed to get access token duration", zap.Error(err))
		return nil, fmt.Errorf("env %s is empty", consts.AccessTokenDuration)
	}

	return &AuthConfig{
		secretKey:            []byte(secretKey),
		refreshTokenDuration: time.Duration(refreshTokenDuration) * time.Hour,
		accessTokenDuration:  time.Duration(accessTokenDuration) * time.Hour,
	}, nil
}

func (a *AuthConfig) GetSecretKey() []byte {
	return a.secretKey
}
func (a *AuthConfig) GetRefreshTokenDuration() time.Duration {
	return a.refreshTokenDuration
}
func (a *AuthConfig) GetAccessTokenDuration() time.Duration {
	return a.accessTokenDuration
}
