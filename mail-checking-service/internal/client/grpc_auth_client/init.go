package grpc_auth_client

import (
	"context"
	"fmt"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

type AuthClient struct {
	authClient authService.AuthV1Client
}

func NewAuthClient(authClient authService.AuthV1Client) IAuthClient {
	return &AuthClient{
		authClient: authClient,
	}
}

func (a *AuthClient) ConfirmEmail(ctx context.Context, email string) error {
	_, err := a.authClient.ConfirmEmail(ctx, &authService.ConfirmEmailRequest{
		Email: email,
	})
	if err != nil {
		logger.Error("failed to confirm email", zap.Error(err))
		return fmt.Errorf("failed to confirm email: %v", err)
	}
	return nil
}

func (a *AuthClient) ValidateToken(ctx context.Context, accessToken string) error {
	_, err := a.authClient.ValidateToken(ctx, &authService.ValidateTokenRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		logger.Error("failed to validate token", zap.Error(err))
		return fmt.Errorf("failed to validate token: %v", err)
	}

	return nil

}
