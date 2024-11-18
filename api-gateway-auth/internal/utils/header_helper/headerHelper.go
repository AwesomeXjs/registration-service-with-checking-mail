package header_helper

import (
	"github.com/labstack/echo/v4"
)

// IHeaderHelper defines methods for handling tokens in HTTP headers and cookies.
type IHeaderHelper interface {
	// GetRefreshTokenFromCookie retrieves the refresh token from the cookie.
	GetRefreshTokenFromCookie(c echo.Context, key string) (string, error)

	// SetRefreshTokenInCookie sets the refresh token in the cookie.
	SetRefreshTokenInCookie(c echo.Context, key string, value string)

	// GetAccessTokenFromHeader retrieves the access token from the request header (Bearer ...).
	GetAccessTokenFromHeader(c echo.Context) (string, error)
}
