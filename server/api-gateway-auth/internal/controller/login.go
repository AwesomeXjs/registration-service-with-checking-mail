package controller

import (
	"net/http"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/response"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (c *Controller) Login(ctx echo.Context) error {
	var Request model.LoginRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Warn("failed to validate struct", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	result, err := c.authClient.Login(ctx.Request().Context(), converter.FromModelToProtoLogin(&Request))

	if err != nil {
		if strings.Contains(err.Error(), "invalid password") {
			logger.Warn("failed to login", zap.Error(err))
			return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad request", "invalid password")
		}
		logger.Error("failed to login", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	c.hh.SetRefreshTokenInCookie(ctx, "refresh_token", result.GetRefreshToken())

	return ctx.JSON(200, converter.ToModelFromProtoLogin(result))
}
