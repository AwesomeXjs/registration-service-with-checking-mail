package grpc_auth_client

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/metrics"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
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

	const mark = "Client.grpc_auth_client.ConfirmEmail"

	_, err := a.authClient.ConfirmEmail(ctx, &authService.ConfirmEmailRequest{
		Email: email,
	})
	if err != nil {
		logger.Error("failed to confirm email", mark, zap.Error(err))
		return fmt.Errorf("failed to confirm email: %v", err)
	}

	metrics.IncSuccessVerificationCounter()

	return nil
}

// ValidateToken sends a request to validate the provided access token.
func (a *AuthClient) ValidateToken(ctx context.Context, accessToken string) error {

	const mark = "Client.grpc_auth_client.ValidateToken"

	_, err := a.authClient.ValidateToken(ctx, &authService.ValidateTokenRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		logger.Error("failed to validate token", mark, zap.Error(err))
		return fmt.Errorf("failed to validate token: %v", err)
	}

	return nil
}
