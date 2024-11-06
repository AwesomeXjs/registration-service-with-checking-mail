package controller

import (
	"net/http"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ValidateToken - validate access token
// @Summary validate access token
// @Security ApiKeyAuth
// @Tags Token
// @Description validate access token
// @ID validate-access-token
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/validate-token [get]
func (c *Controller) ValidateToken(ctx echo.Context) error {
	accessToken, err := c.hh.GetAccessTokenFromHeader(ctx)
	if err != nil {
		logger.Warn("failed to get access token from header", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	_, err = c.authClient.ValidateToken(ctx.Request().Context(), converter.ToProtoValidateToken(accessToken))
	if err != nil {
		if strings.Contains(err.Error(), "failed to verify") {
			logger.Warn("failed to validate token", zap.Error(err))
			return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", "failed to verify token")
		}
		logger.Error("failed to validate token", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	return nil
}
