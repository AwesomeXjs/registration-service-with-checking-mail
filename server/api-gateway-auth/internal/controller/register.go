package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/response"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (c *Controller) Registration(ctx echo.Context) error {
	var Request model.RegistrationRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Warn("failed to bind request", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Warn("failed to validate struct", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	result, err := c.authClient.Registration(ctx.Request().Context(), converter.FromModelToProtoRegister(&Request))
	if err != nil {
		logger.Error("failed to register", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	c.hh.SetRefreshTokenInCookie(ctx, "refresh_token", result.GetRefreshToken())

	return ctx.JSON(200, converter.ToModelFromProtoRegister(result))
}
