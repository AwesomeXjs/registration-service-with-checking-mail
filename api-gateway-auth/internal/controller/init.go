package controller

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/grpc_auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/mail_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/headers_manager"
)

// Controller handles the authentication and header-related operations.
type Controller struct {
	authClient grpc_auth_client.AuthClient // Authentication client for interacting with the auth service.
	mailClient mail_client.MailClient
	hh         headers_manager.IHeaderHelper // Header helper for managing tokens in headers and cookies.
}

// New creates a new instance of the Controller.
// It takes an authentication client and a header helper as dependencies.
func New(authClient grpc_auth_client.AuthClient, mailClient mail_client.MailClient, hh headers_manager.IHeaderHelper) *Controller {
	return &Controller{
		authClient: authClient,
		mailClient: mailClient,
		hh:         hh,
	}
}
