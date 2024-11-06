package controller

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/client/auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/header_helper"
)

type Controller struct {
	authClient auth_client.AuthClient
	hh         header_helper.IHeaderHelper
}

func New(authClient auth_client.AuthClient, hh header_helper.IHeaderHelper) *Controller {
	return &Controller{
		authClient: authClient,
		hh:         hh,
	}
}
