package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UpdatePassword - update password
// @Summary UpdatePassword
// @Security ApiKeyAuth
// @Tags Auth
// @Description update password
// @ID update-password
// @Accept  json
// @Produce  json
// @Param input body model.UpdatePasswordRequest true "new info"
// @Success 200
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/update-password [patch]
func (c *Controller) UpdatePassword(ctx echo.Context) error {
	accessToken, err := c.hh.GetAccessTokenFromHeader(ctx)
	if err != nil {
		logger.Warn("failed to get access token from header", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}
	logger.Debug("get access token from header", zap.String("ACCESS_TOKEN", accessToken))

	err = c.authClient.ValidateToken(ctx.Request().Context(), accessToken)
	if err != nil {
		logger.Warn("failed to validate token", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	var Request model.UpdatePasswordRequest
	err = ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	logger.Debug("update password request: ", zap.Any("request", Request))

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Error("failed to validate struct", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	err = c.authClient.UpdatePassword(ctx.Request().Context(), &Request)
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	return nil
}
