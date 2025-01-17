package grpc_auth_client

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// GRPCClient wraps the AuthClient interface, providing methods to interact with
// the authentication gRPC service. It delegates requests to the underlying AuthClient.
type GRPCClient struct {
	authClient authService.AuthV1Client
}

// New creates a new GRPCClient instance, initializing it with the provided AuthClient.
func New(authClient authService.AuthV1Client) AuthClient {
	return &GRPCClient{
		authClient: authClient,
	}
}

// Registration delegates the registration request to the underlying authClient.
// It sends the registration request to the authentication service and returns the response.
func (g *GRPCClient) Registration(ctx context.Context,
	request *model.RegistrationRequest) (*model.RegistrationResponse, string, error) {

	const mark = "Client.grpc_auth_client.Registration"

	result, err := g.authClient.Registration(ctx, converter.FromModelToProtoRegister(request)) // Delegates to the Registration method of the AuthClient.

	if err != nil {
		logger.Error("failed to register user", mark, zap.Error(err))
		return nil, "", fmt.Errorf("failed to register user: %v", err)
	}

	return converter.ToModelFromProtoRegister(result), result.GetRefreshToken(), nil
}

// Login delegates the login request to the underlying authClient.
// It sends the login credentials to the authentication service and returns the response.
func (g *GRPCClient) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, string, error) {

	const mark = "Client.grpc_auth_client.Login"

	result, err := g.authClient.Login(ctx, converter.FromModelToProtoLogin(request)) // Delegates to the Login method of the AuthClient.
	if err != nil {
		logger.Error("failed to login", mark, zap.Error(err))
		return nil, "", fmt.Errorf("failed to login user: %v", err)
	}

	return converter.ToModelFromProtoLogin(result), result.GetRefreshToken(), nil // Delegates to the Login method of the AuthClient.
}

// ValidateToken delegates the token validation request to the underlying authClient.
// It sends the token to the authentication service for validation and returns the response.
func (g *GRPCClient) ValidateToken(ctx context.Context, accessToken string) error {

	const mark = "Client.grpc_auth_client.ValidateToken"

	_, err := g.authClient.ValidateToken(ctx, converter.ToProtoValidateToken(accessToken))
	if err != nil {
		logger.Error("failed to validate token", mark, zap.Error(err))
		return fmt.Errorf("failed to validate token: %v", err)
	}
	return nil // Delegates to the ValidateToken method of the AuthClient.
}

// GetAccessToken delegates the request to get a new access token using a refresh token to the underlying authClient.
// It sends the refresh token to the authentication service and returns the new access token response.
func (g *GRPCClient) GetAccessToken(ctx context.Context, refreshToken string) (string, string, error) {

	const mark = "Client.grpc_auth_client.GetAccessToken"

	result, err := g.authClient.GetAccessToken(ctx, converter.FromModelToProtoGetAccessToken(refreshToken)) // Delegates to the GetAccessToken method of the AuthClient.
	if err != nil {
		logger.Error("failed to get access token", mark, zap.Error(err))
		return "", "", fmt.Errorf("failed to get access token: %v", err)
	}
	return result.GetRefreshToken(), result.GetAccessToken(), nil
}

// UpdatePassword delegates the password update request to the underlying authClient.
// It sends the new password to the authentication service and returns an empty response.
func (g *GRPCClient) UpdatePassword(ctx context.Context, request *model.UpdatePasswordRequest) error {

	const mark = "Client.grpc_auth_client.UpdatePassword"

	_, err := g.authClient.UpdatePassword(ctx, converter.FromModelToProtoUpdatePass(request))
	if err != nil {
		logger.Error("failed to update password", mark, zap.Error(err))
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil // Delegates to the UpdatePassword method of the AuthClient.
}
