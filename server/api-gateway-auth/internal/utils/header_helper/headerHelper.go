package header_helper

import (
	"github.com/labstack/echo/v4"
)

type IHeaderHelper interface {
	GetRefreshTokenFromCookie(c echo.Context, key string) (string, error)
	SetRefreshTokenInCookie(c echo.Context, key string, value string)
	GetAccessTokenFromHeader(c echo.Context) (string, error)
}
