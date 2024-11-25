package grpc_auth_client

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
)

// AuthClient provides an interface for communicating with the authentication gRPC service.
// It includes methods for user login, registration, password updates, token generation, and token validation.
type AuthClient interface {
	// Login performs user authentication using email and password.
	// It returns a LoginResponse containing an access token, refresh token and userID.
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, string, error)

	// Registration registers a new user with the given details.
	// It returns a RegistrationResponse containing an access token, refresh token and the user ID.
	Registration(ctx context.Context, request *model.RegistrationRequest) (*model.RegistrationResponse, string, error)

	// UpdatePassword updates the user's password with a new one.
	// It returns an empty response on success.
	UpdatePassword(ctx context.Context, request *model.UpdatePasswordRequest) error

	// GetAccessToken generates a new access token using the provided refresh token.
	// It returns a GetAccessTokenResponse containing the new pair of access and refresh tokens.
	GetAccessToken(ctx context.Context, refreshToken string) (string, string, error)

	// ValidateToken verifies the validity of an access token.
	// It returns an empty response if the token is valid.
	ValidateToken(ctx context.Context, accessToken string) error
}
