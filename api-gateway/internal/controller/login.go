package controller

import (
	"net/http"
	"strings"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Login - login
// @Summary Login
// @Tags Auth
// @Description login into system
// @ID login
// @Accept  json
// @Produce  json
// @Param input body model.LoginRequest true "login info"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/login [post]
func (c *Controller) Login(ctx echo.Context) error {

	const mark = "Controller.Login"

	var Request model.LoginRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	logger.Debug("login request: ", mark, zap.Any("request", Request))

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Warn("failed to validate struct", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	span, contextWithTrace := opentracing.StartSpanFromContext(ctx.Request().Context(), "Login")
	defer span.Finish()

	span.SetTag("email", Request.Email)

	result, refreshToken, err := c.authClient.Login(contextWithTrace, &Request)
	if err != nil {
		if strings.Contains(err.Error(), "invalid password") {
			logger.Warn("failed to login", mark, zap.Error(err))
			return response.RespHelper(ctx, http.StatusBadRequest, "Bad request", "invalid password")
		}
		logger.Error("failed to login", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	c.hh.SetRefreshTokenInCookie(ctx, RefreshTokenKey, refreshToken)

	return ctx.JSON(http.StatusOK, result)
}
