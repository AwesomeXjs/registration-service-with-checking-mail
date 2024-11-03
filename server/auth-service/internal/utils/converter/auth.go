package converter

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

func ToInfoFromProto(info *authService.RegistrationRequest) *model.UserInfo {
	return &model.UserInfo{
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
		Name:     info.GetName(),
		Surname:  info.GetSurname(),
		Role:     info.GetRole(),
	}
}

func ToProtoFromRegResponse(resp *model.RegistrationResponse) *authService.RegistrationResponse {
	return &authService.RegistrationResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		UserId:       resp.UserId,
	}
}
