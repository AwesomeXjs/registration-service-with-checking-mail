package middlewares

import (
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		logger.Info("Request details",
			zap.String("method", c.Request().Method),
			zap.String("path", c.Request().URL.Path),
			zap.Duration("duration", time.Since(start)),
		)

		return err
	}
}
