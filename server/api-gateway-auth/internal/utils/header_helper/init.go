package header_helper

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type HeaderHelper struct {
}

func New() *HeaderHelper {
	return &HeaderHelper{}
}

func (h HeaderHelper) GetRefreshTokenFromCookie(c echo.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "No refresh token found",
			})
		}
		return "", err
	}
	if cookie.Value == "" {
		return "", errors.New("no refresh token found")
	}

	return cookie.Value, nil
}

func (h HeaderHelper) SetRefreshTokenInCookie(c echo.Context, key string, value string) {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(cookie)
}

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
