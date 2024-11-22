package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
)

// ToInfoFromProto converts a RegistrationRequest from the authService to a UserInfo model.
// This is used to transform incoming registration data into a format suitable for further processing.
func ToInfoFromProto(info *authService.RegistrationRequest) *model.UserInfo {
	return &model.UserInfo{
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
		Name:     info.GetName(),
		Surname:  info.GetSurname(),
		Role:     info.GetRole(),
	}
}

// ToLoginInfoFromProto converts a LoginRequest from the authService to a LoginInfo model.
// This function maps login details from the incoming request to the model structure used in the application.
func ToLoginInfoFromProto(info *authService.LoginRequest) *model.LoginInfo {
	return &model.LoginInfo{
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
	}
}

// ToUpdatePassFromProto converts an UpdatePasswordRequest from the authService to an UpdatePassInfo model.
// It is used to extract and prepare email and new password information for updating user credentials.
func ToUpdatePassFromProto(info *authService.UpdatePasswordRequest) *model.UpdatePassInfo {
	return &model.UpdatePassInfo{
		Email:       info.GetEmail(),
		NewPassword: info.GetNewPassword(),
	}
}
