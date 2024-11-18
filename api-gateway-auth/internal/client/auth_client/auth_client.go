package auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AuthClient provides an interface for communicating with the authentication gRPC service.
// It includes methods for user login, registration, password updates, token generation, and token validation.
type AuthClient interface {
	// Login performs user authentication using email and password.
	// It returns a LoginResponse containing an access token, refresh token and userID.
	Login(ctx context.Context, in *authService.LoginRequest, opts ...grpc.CallOption) (*authService.LoginResponse, error)

	// Registration registers a new user with the given details.
	// It returns a RegistrationResponse containing an access token, refresh token and the user ID.
	Registration(ctx context.Context, in *authService.RegistrationRequest, opts ...grpc.CallOption) (*authService.RegistrationResponse, error)

	// UpdatePassword updates the user's password with a new one.
	// It returns an empty response on success.
	UpdatePassword(ctx context.Context, in *authService.UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)

	// GetAccessToken generates a new access token using the provided refresh token.
	// It returns a GetAccessTokenResponse containing the new pair of access and refresh tokens.
	GetAccessToken(ctx context.Context, in *authService.GetAccessTokenRequest, opts ...grpc.CallOption) (*authService.GetAccessTokenResponse, error)

	// ValidateToken verifies the validity of an access token.
	// It returns an empty response if the token is valid.
	ValidateToken(ctx context.Context, in *authService.ValidateTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
