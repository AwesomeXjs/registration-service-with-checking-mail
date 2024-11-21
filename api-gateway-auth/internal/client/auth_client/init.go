package auth_client

import (
	"context"

	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GRPCClient wraps the AuthClient interface, providing methods to interact with
// the authentication gRPC service. It delegates requests to the underlying AuthClient.
type GRPCClient struct {
	authClient AuthClient
}

// New creates a new GRPCClient instance, initializing it with the provided AuthClient.
func New(authClient AuthClient) AuthClient {
	return &GRPCClient{
		authClient: authClient,
	}
}

// Registration delegates the registration request to the underlying authClient.
// It sends the registration request to the authentication service and returns the response.
func (g *GRPCClient) Registration(ctx context.Context, in *authService.RegistrationRequest, opts ...grpc.CallOption) (*authService.RegistrationResponse, error) {
	return g.authClient.Registration(ctx, in, opts...) // Delegates to the Registration method of the AuthClient.
}

// Login delegates the login request to the underlying authClient.
// It sends the login credentials to the authentication service and returns the response.
func (g *GRPCClient) Login(ctx context.Context, in *authService.LoginRequest, opts ...grpc.CallOption) (*authService.LoginResponse, error) {
	return g.authClient.Login(ctx, in, opts...) // Delegates to the Login method of the AuthClient.
}

// ValidateToken delegates the token validation request to the underlying authClient.
// It sends the token to the authentication service for validation and returns the response.
func (g *GRPCClient) ValidateToken(ctx context.Context, in *authService.ValidateTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.authClient.ValidateToken(ctx, in, opts...) // Delegates to the ValidateToken method of the AuthClient.
}

// GetAccessToken delegates the request to get a new access token using a refresh token to the underlying authClient.
// It sends the refresh token to the authentication service and returns the new access token response.
func (g *GRPCClient) GetAccessToken(ctx context.Context, in *authService.GetAccessTokenRequest, opts ...grpc.CallOption) (*authService.GetAccessTokenResponse, error) {
	return g.authClient.GetAccessToken(ctx, in, opts...) // Delegates to the GetAccessToken method of the AuthClient.
}

// UpdatePassword delegates the password update request to the underlying authClient.
// It sends the new password to the authentication service and returns an empty response.
func (g *GRPCClient) UpdatePassword(ctx context.Context, in *authService.UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.authClient.UpdatePassword(ctx, in, opts...) // Delegates to the UpdatePassword method of the AuthClient.
}
