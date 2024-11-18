package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

// ToModelFromProtoRegister converts a RegistrationResponse from gRPC to a RegistrationResponse model.
func ToModelFromProtoRegister(info *authService.RegistrationResponse) *model.RegistrationResponse {
	return &model.RegistrationResponse{
		AccessToken: info.GetAccessToken(),
		UserID:      info.GetUserId(),
	}
}

// ToModelFromProtoLogin converts a LoginResponse from gRPC to a LoginResponse model.
func ToModelFromProtoLogin(info *authService.LoginResponse) *model.LoginResponse {
	return &model.LoginResponse{
		AccessToken: info.GetAccessToken(),
		UserID:      info.GetUserId(),
	}
}
