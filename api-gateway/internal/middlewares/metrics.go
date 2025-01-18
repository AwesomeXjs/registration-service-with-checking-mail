package middlewares

import (
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/metrics"
	"github.com/labstack/echo/v4"
)

// MetricsInterceptor is a middleware that measures request counts and response times,
// and records them using the metrics package.
func MetricsInterceptor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		metrics.IncRequestCounter()

		timeStart := time.Now()

		err := next(c)

		diffTime := time.Since(timeStart)

		if err != nil {
			metrics.IncResponseCounter("error", c.Request().Method)
			metrics.ObserveResponseTime("error", diffTime.Seconds())
		} else {
			metrics.IncResponseCounter("success", c.Request().Method)
			metrics.ObserveResponseTime("success", diffTime.Seconds())
		}

		return err
	}
}
