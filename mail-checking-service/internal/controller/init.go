package controller

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
)

// Controller handles gRPC requests for the authentication service.
// It embeds UnimplementedMailV1Server to provide default implementations
// for the gRPC server methods, allowing customization where needed.
type Controller struct {
	mail_v1.UnimplementedMailV1Server
}

// New creates a new instance of Controller with the provided service.
// The service is used to manage authentication-related operations.
func New() *Controller {
	return &Controller{}
}
