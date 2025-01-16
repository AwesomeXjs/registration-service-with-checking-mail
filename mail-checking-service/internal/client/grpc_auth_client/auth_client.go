package grpc_auth_client

import (
	"context"
)

// IAuthClient defines the interface for interacting with an authentication client.
// It provides methods to confirm email addresses and validate access tokens.
type IAuthClient interface {
	// ConfirmEmail confirms the given email address.
	// It returns an error if the confirmation fails.
	ConfirmEmail(ctx context.Context, email string) error

	// ValidateToken validates the provided access token.
	// It returns an error if the token is invalid or expired.
	ValidateToken(ctx context.Context, accessToken string) error
}
