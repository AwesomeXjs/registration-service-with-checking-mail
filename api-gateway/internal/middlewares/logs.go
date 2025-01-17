package middlewares

import (
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Logger is a middleware that logs details of each HTTP request.
// It logs the request method, path, and duration. In case of an error, it also logs the error details.
func Logger(next echo.HandlerFunc) echo.HandlerFunc {

	const mark = "middlewares.Logger"

	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		if err != nil {
			logger.Error("Failed to process request",
				mark,
				zap.Error(err),
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.Duration("duration", time.Since(start)))
		}

		logger.Info("Request details",
			mark,
			zap.String("method", c.Request().Method),
			zap.String("path", c.Request().URL.Path),
			zap.Duration("duration", time.Since(start)),
		)

		return err
	}
}
