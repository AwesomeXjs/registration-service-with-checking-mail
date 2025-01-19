package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	// RefreshTokenKey is the key used to store and retrieve the refresh token from a storage mechanism (e.g., cookie or cache).
	RefreshTokenKey = "refresh_token"
)

// GetAccessToken - get access token
// @Summary Get Access Token
// @Security ApiKeyAuth
// @Tags Token
// @Description get new access token from refresh token
// @ID get-access-token
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/get-access-token [get]
func (c *Controller) GetAccessToken(ctx echo.Context) error {

	const mark = "Controller.GetAccessToken"

	refreshToken, err := c.hh.GetRefreshTokenFromCookie(ctx, RefreshTokenKey)
	if err != nil {
		logger.Error("failed to get refresh token from cookie", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}
	logger.Debug("get refresh token", mark, zap.String("REFRESH_TOKEN", refreshToken))

	refreshToken, accessToken, err := c.authClient.GetAccessToken(ctx.Request().Context(), refreshToken)
	if err != nil {
		logger.Error("failed to get access token", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	c.hh.SetRefreshTokenInCookie(ctx, RefreshTokenKey, refreshToken)

	return response.RespHelper(ctx, http.StatusOK, "OK", accessToken)

}
