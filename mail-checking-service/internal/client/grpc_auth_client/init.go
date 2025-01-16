package grpc_auth_client

import (
	"context"
	"fmt"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

// AuthClient implements IAuthClient for interacting with AuthService via gRPC.
type AuthClient struct {
	authClient authService.AuthV1Client // gRPC client for AuthService
}

// NewAuthClient creates and returns a new AuthClient instance.
func NewAuthClient(authClient authService.AuthV1Client) IAuthClient {
	return &AuthClient{
		authClient: authClient,
	}
}

// ConfirmEmail sends a request to confirm the provided email address.
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

// ValidateToken sends a request to validate the provided access token.
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
