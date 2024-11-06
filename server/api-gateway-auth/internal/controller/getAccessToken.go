package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/response"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (c *Controller) GetAccessToken(ctx echo.Context) error {
	cookie, err := c.hh.GetRefreshTokenFromCookie(ctx, "refresh_token")
	if err != nil {
		logger.Error("failed to get refresh token from cookie", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	accessToken, err := c.authClient.GetAccessToken(ctx.Request().Context(), &authService.GetAccessTokenRequest{
		RefreshToken: cookie,
	})
	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	return response.ResponseHelper(ctx, http.StatusOK, "OK", accessToken.AccessToken)

}
