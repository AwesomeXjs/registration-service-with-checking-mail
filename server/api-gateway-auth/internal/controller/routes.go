package controller

import (
	"github.com/labstack/echo/v4"
)

func (e *Controller) InitRoutes(server *echo.Echo) {
	// Swagger init
	//server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "Hello, World!",
		})
	})
	{
		// song routes
		_ = api.Group("/v1")
		{

		}
	}
}
