package response

import (
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Response struct {
	Title   string `json:"title"`
	Detail  string `json:"detail"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

// ResponseHelper helper function
func ResponseHelper(ctx echo.Context, statusCode int, message, detail string) error {
	err := ctx.JSON(statusCode, Response{Title: message, Detail: detail, Request: ctx.Request().RequestURI, Time: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		logger.Info("failed to write response", zap.String("error", err.Error()))
		return err
	}
	return nil
}
