package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Registration - Registration
// @Summary Registration
// @Tags Auth
// @Description Registration new user
// @ID registration
// @Accept  json
// @Produce  json
// @Param input body model.RegistrationRequest true "registration info"
// @Success 200 {object} model.RegistrationResponse
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/register [post]
func (c *Controller) Registration(ctx echo.Context) error {

	const mark = "Controller.Registration"

	var Request model.RegistrationRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Warn("failed to bind request", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	logger.Debug("registration request: ", mark, zap.Any("request", Request))

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Warn("failed to validate struct", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	result, refreshToken, err := c.authClient.Registration(ctx.Request().Context(), &Request)
	if err != nil {
		logger.Error("failed to register", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	c.hh.SetRefreshTokenInCookie(ctx, RefreshTokenKey, refreshToken)

	return ctx.JSON(http.StatusOK, result)
}
