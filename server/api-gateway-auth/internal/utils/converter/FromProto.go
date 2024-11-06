package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

func ToModelFromProtoRegister(info *authService.RegistrationResponse) *model.RegistrationResponse {
	return &model.RegistrationResponse{
		AccessToken: info.GetAccessToken(),
		UserID:      info.GetUserId(),
	}
}

func ToModelFromProtoLogin(info *authService.LoginResponse) *model.LoginResponse {
	return &model.LoginResponse{
		AccessToken: info.GetAccessToken(),
		UserID:      info.GetUserId(),
	}
}
