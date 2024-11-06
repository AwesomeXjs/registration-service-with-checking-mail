package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/client/auth_client"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

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

func (c *Controller) Login(ctx echo.Context) error {
	var Request model.LoginRequest
	err := ctx.Bind(&Request)
	if err != nil {
		return err
	}
	// тут валидация

	result, err := c.authClient.Login(ctx.Request().Context(), &authService.LoginRequest{
		Email:    Request.Email,
		Password: Request.Password,
	})
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    result.GetRefreshToken(),
		Path:     "/",
		HttpOnly: true,
	}
	ctx.SetCookie(cookie)
	// тут обработка результата

	return ctx.JSON(200, &result)
}

func (c *Controller) GetAccessToken(ctx echo.Context) error {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "No refresh token found",
			})
		}
		return err
	}

	accessToken, err := c.authClient.GetAccessToken(ctx.Request().Context(), &authService.GetAccessTokenRequest{
		RefreshToken: cookie.Value,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(200, &model.GetAccessTokenResponse{
		AccessToken: accessToken.GetAccessToken()})

}
