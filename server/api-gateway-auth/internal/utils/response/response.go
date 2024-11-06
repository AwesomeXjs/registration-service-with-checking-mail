package response

import (
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Response struct defines the structure of the response body.
// It includes a title, detail about the error or message, the request URI, and the current time.
type Response struct {
	Title   string `json:"title"`   // The main message or title of the response
	Detail  string `json:"detail"`  // Detailed explanation of the response (e.g., error details)
	Request string `json:"request"` // The URI of the request that generated the response
	Time    string `json:"time"`    // The timestamp when the response was generated
}

// RespHelper is a utility function to send JSON responses.
// It accepts the HTTP context, a status code, a message, and detailed information to send in the response.
func RespHelper(ctx echo.Context, statusCode int, message, detail string) error {
	// Construct the response object using the provided parameters.
	err := ctx.JSON(statusCode, Response{
		Title:   message,                                  // Title of the response (message)
		Detail:  detail,                                   // Detailed message or error information
		Request: ctx.Request().RequestURI,                 // The URI of the request that generated the response
		Time:    time.Now().Format("2006-01-02 15:04:05"), // The current timestamp in a specific format
	})

	// If there's an error in writing the response, log it and return the error.
	if err != nil {
		logger.Info("failed to write response", zap.String("error", err.Error()))
		return err
	}

	// Return nil indicating that the response was successfully written.
	return nil
}
