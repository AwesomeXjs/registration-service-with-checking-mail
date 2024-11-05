package controller

import "github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/client/auth_client"

/*
тут валидация и работа с хедерами
и обращение к другим сервисам
*/

type Controller struct {
	authClient auth_client.AuthClient
}

func New(authClient auth_client.AuthClient) *Controller {
	return &Controller{
		authClient: authClient,
	}
}
