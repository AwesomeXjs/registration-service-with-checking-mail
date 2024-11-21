package controller

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/client/auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/utils/header_helper"
)

// Controller handles the authentication and header-related operations.
type Controller struct {
	authClient auth_client.AuthClient      // Authentication client for interacting with the auth service.
	hh         header_helper.IHeaderHelper // Header helper for managing tokens in headers and cookies.
}

// New creates a new instance of the Controller.
// It takes an authentication client and a header helper as dependencies.
func New(authClient auth_client.AuthClient, hh header_helper.IHeaderHelper) *Controller {
	return &Controller{
		authClient: authClient,
		hh:         hh,
	}
}
