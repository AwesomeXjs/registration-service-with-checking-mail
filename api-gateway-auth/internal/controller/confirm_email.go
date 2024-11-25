package controller

import (
	"net/http"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ConfirmEmail - verify email
// @Summary verify email
// @Tags Verify
// @Security ApiKeyAuth
// @Description get new access token from refresh token
// @ID verify-email
// @Accept  json
// @Produce  json
// @Param input body model.ConfirmEmailRequest true "email info"
// @Success 200
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/confirm-email [post]
func (c *Controller) ConfirmEmail(ctx echo.Context) error {
	accessToken, err := c.hh.GetAccessTokenFromHeader(ctx)
	if err != nil {
		logger.Warn("failed to get access token from header", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	var Request model.ConfirmEmailRequest
	err = ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return response.RespHelper(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	}

	_, err = govalidator.ValidateStruct(Request)
	if err != nil {
		logger.Warn("failed to validate struct", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnprocessableEntity, "Bad Request", err.Error())
	}

	_, err = c.mailClient.CheckUniqueCode(ctx.Request().Context(), &mail_v1.CheckUniqueCodeRequest{
		Code:        Request.Code,
		Email:       Request.Email,
		AccessToken: accessToken,
	})
	if err != nil {
		logger.Error("failed to check unique code", zap.Error(err))
		return response.RespHelper(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
	}

	return nil
}
