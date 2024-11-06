package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/response"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (c *Controller) UpdatePassword(ctx echo.Context) error {

	var Request model.UpdatePasswordRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Error("failed to validate struct", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	_, err = c.authClient.UpdatePassword(ctx.Request().Context(), &authService.UpdatePasswordRequest{
		Email:       Request.Email,
		NewPassword: Request.NewPassword,
	})
	if err != nil {
		logger.Error("failed to update password", zap.Error(err))
		return response.ResponseHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	return nil
}
