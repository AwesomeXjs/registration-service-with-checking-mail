package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

func FromModelToProtoRegister(info *model.RegistrationRequest) *authService.RegistrationRequest {
	return &authService.RegistrationRequest{
		Email:    info.Email,
		Password: info.Password,
		Name:     info.Name,
		Surname:  info.Surname,
		Role:     info.Role,
	}
}

func FromModelToProtoLogin(info *model.LoginRequest) *authService.LoginRequest {
	return &authService.LoginRequest{
		Email:    info.Email,
		Password: info.Password,
	}
}

func FromModelToProtoUpdatePass(info *model.UpdatePasswordRequest) *authService.UpdatePasswordRequest {
	return &authService.UpdatePasswordRequest{
		Email:       info.Email,
		NewPassword: info.NewPassword,
	}
}

func FromModelToProtoGetAccessToken(token string) *authService.GetAccessTokenRequest {
	return &authService.GetAccessTokenRequest{
		RefreshToken: token,
	}
}
func ToProtoValidateToken(token string) *authService.ValidateTokenRequest {
	return &authService.ValidateTokenRequest{
		AccessToken: token,
	}
}
