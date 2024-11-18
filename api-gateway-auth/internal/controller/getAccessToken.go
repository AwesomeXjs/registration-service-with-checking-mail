package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/consts"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
	refreshToken, err := c.hh.GetRefreshTokenFromCookie(ctx, consts.RefreshTokenKey)
	if err != nil {
		logger.Error("failed to get refresh token from cookie", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}
	logger.Debug("get refresh token", zap.String("REFRESH_TOKEN", refreshToken))

	accessToken, err := c.authClient.GetAccessToken(ctx.Request().Context(), converter.FromModelToProtoGetAccessToken(refreshToken))
	if err != nil {
		logger.Error("failed to get access token", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	c.hh.SetRefreshTokenInCookie(ctx, consts.RefreshTokenKey, accessToken.RefreshToken)

	return response.RespHelper(ctx, http.StatusOK, "OK", accessToken.AccessToken)

}
