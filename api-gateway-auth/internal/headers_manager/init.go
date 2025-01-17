package headers_manager

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/response"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
)

// HeaderHelper provides methods for working with authentication tokens in cookies and headers.
type HeaderHelper struct {
}

// New creates a new instance of HeaderHelper.
func New() *HeaderHelper {
	return &HeaderHelper{}
}

// GetRefreshTokenFromCookie retrieves the refresh token from the cookie.
// Returns an error if the token is missing or invalid.
func (h HeaderHelper) GetRefreshTokenFromCookie(c echo.Context, key string) (string, error) {

	const mark = "HeadersManager.GetRefreshTokenFromCookie"

	cookie, err := c.Cookie(key)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			logger.Error("failed to get refresh token from cookie", mark, zap.Error(err))
			return "", response.RespHelper(c, http.StatusUnauthorized, "Unauthorized", err.Error())
		}
		return "", err
	}
	if cookie.Value == "" {
		logger.Error("no refresh token found", mark, zap.Error(err))
		return "", errors.New("no refresh token found")
	}

	return cookie.Value, nil
}

// SetRefreshTokenInCookie sets the refresh token in the cookie.
// It marks the cookie as HttpOnly to prevent client-side access.
func (h HeaderHelper) SetRefreshTokenInCookie(c echo.Context, key string, value string) {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(cookie)
}

// GetAccessTokenFromHeader retrieves the access token from the Authorization header.
// It expects the token to be prefixed with "Bearer ".
func (h HeaderHelper) GetAccessTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", errors.New("invalid authorization header format")
	}

	token := strings.TrimPrefix(authHeader, bearerPrefix)
	return token, nil
}
