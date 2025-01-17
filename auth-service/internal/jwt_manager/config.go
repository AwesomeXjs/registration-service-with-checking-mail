package jwt_manager

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"go.uber.org/zap"
)

// AuthConfig holds the configuration for authentication tokens.
type AuthConfig struct {
	secretKey            []byte
	refreshTokenDuration time.Duration
	accessTokenDuration  time.Duration
}

// NewAuthConfig creates a new AuthConfig instance by reading environment variables.
func NewAuthConfig() (*AuthConfig, error) {
	const mark = "jwt_manager.NewAuthConfig"

	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) == 0 {
		logger.Error("failed to get secret key", mark, zap.String("secret key", "SECRET_KEY"))
		return nil, fmt.Errorf("env %s is empty", "SECRET_KEY")
	}

	refreshTokenDuration, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_DURATION"))
	if err != nil {
		logger.Error("failed to get refresh token duration", mark, zap.Error(err))
		return nil, fmt.Errorf("env %s is empty", "REFRESH_TOKEN_DURATION")
	}
	accessTokenDuration, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		logger.Error("failed to get access token duration", mark, zap.Error(err))
		return nil, fmt.Errorf("env %s is empty", "ACCESS_TOKEN_DURATION")
	}

	return &AuthConfig{
		secretKey:            []byte(secretKey),
		refreshTokenDuration: time.Duration(refreshTokenDuration) * time.Hour,
		accessTokenDuration:  time.Duration(accessTokenDuration) * time.Hour,
	}, nil
}

// GetSecretKey returns the secret key for token generation.
func (a *AuthConfig) GetSecretKey() []byte {
	return a.secretKey
}

// GetRefreshTokenDuration returns the duration for refresh tokens.
func (a *AuthConfig) GetRefreshTokenDuration() time.Duration {
	return a.refreshTokenDuration
}

// GetAccessTokenDuration returns the duration for access tokens.
func (a *AuthConfig) GetAccessTokenDuration() time.Duration {
	return a.accessTokenDuration
}
