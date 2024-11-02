package controller

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/service"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

type Controller struct {
	authService.UnimplementedAuthV1Server
	svc service.IService
}

func New(svc service.IService) *Controller {
	return &Controller{
		svc: svc,
	}
}
