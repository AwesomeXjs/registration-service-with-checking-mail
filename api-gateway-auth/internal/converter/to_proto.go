package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
)

// FromModelToProtoRegister converts a RegistrationRequest model to a gRPC RegistrationRequest.
func FromModelToProtoRegister(info *model.RegistrationRequest) *authService.RegistrationRequest {
	return &authService.RegistrationRequest{
		Email:    info.Email,
		Password: info.Password,
		Name:     info.Name,
		Surname:  info.Surname,
		Role:     info.Role,
	}
}

// FromModelToProtoLogin converts a LoginRequest model to a gRPC LoginRequest.
func FromModelToProtoLogin(info *model.LoginRequest) *authService.LoginRequest {
	return &authService.LoginRequest{
		Email:    info.Email,
		Password: info.Password,
	}
}

// FromModelToProtoUpdatePass converts an UpdatePasswordRequest model to a gRPC UpdatePasswordRequest.
func FromModelToProtoUpdatePass(info *model.UpdatePasswordRequest) *authService.UpdatePasswordRequest {
	return &authService.UpdatePasswordRequest{
		Email:       info.Email,
		NewPassword: info.NewPassword,
	}
}

// FromModelToProtoGetAccessToken creates a gRPC GetAccessTokenRequest from a refresh token string.
func FromModelToProtoGetAccessToken(token string) *authService.GetAccessTokenRequest {
	return &authService.GetAccessTokenRequest{
		RefreshToken: token,
	}
}

// ToProtoValidateToken creates a gRPC ValidateTokenRequest from an access token string.
func ToProtoValidateToken(token string) *authService.ValidateTokenRequest {
	return &authService.ValidateTokenRequest{
		AccessToken: token,
	}
}
