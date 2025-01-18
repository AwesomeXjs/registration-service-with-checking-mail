package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
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

	const mark = "Controller.UpdatePassword"

	accessToken, err := c.hh.GetAccessTokenFromHeader(ctx)
	if err != nil {
		logger.Warn("failed to get access token from header", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}
	logger.Debug("get access token from header", mark, zap.String("ACCESS_TOKEN", accessToken))

	span, contextWithTrace := opentracing.StartSpanFromContext(ctx.Request().Context(), "UpdatePassword")
	defer span.Finish()

	span.SetTag("access_token", accessToken)

	err = c.authClient.ValidateToken(contextWithTrace, accessToken)
	if err != nil {
		logger.Warn("failed to validate token", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	var Request model.UpdatePasswordRequest
	err = ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}
	logger.Debug("update password request: ", mark, zap.Any("request", Request))

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Error("failed to validate struct", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	err = c.authClient.UpdatePassword(ctx.Request().Context(), &Request)
	if err != nil {
		logger.Error("failed to update password", mark, zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	return nil
}
