package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
)

// ToProtoFromLoginResponse converts the AuthResponse model into a gRPC-compatible LoginResponse.
// This is useful for sending authentication details back to the client after a successful login.
func ToProtoFromLoginResponse(resp *model.AuthResponse) *authService.LoginResponse {
	return &authService.LoginResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		UserId:       resp.UserID,
	}
}

// ToProtoFromNewPairTokens transforms the NewPairTokens model into a gRPC GetAccessTokenResponse.
// This is used to provide the client with a new pair of access and refresh tokens.
func ToProtoFromNewPairTokens(resp *model.NewPairTokens) *authService.GetAccessTokenResponse {
	return &authService.GetAccessTokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
}

// ToProtoFromRegResponse converts the AuthResponse model into a RegistrationResponse for gRPC.
// This allows the client to receive authentication tokens after successful registration.
func ToProtoFromRegResponse(resp *model.AuthResponse) *authService.RegistrationResponse {
	return &authService.RegistrationResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		UserId:       resp.UserID,
	}
}
