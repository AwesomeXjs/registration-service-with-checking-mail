package controller

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/service"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
)

// Controller handles gRPC requests for the authentication service.
// It embeds UnimplementedAuthV1Server to provide default implementations
// for the gRPC server methods, allowing customization where needed.
type Controller struct {
	authService.UnimplementedAuthV1Server
	svc service.IService
}

// New creates a new instance of Controller with the provided service.
// The service is used to manage authentication-related operations.
func New(svc service.IService) *Controller {
	return &Controller{
		svc: svc,
	}
}
