package controller

import (
	"github.com/labstack/echo/v4"
)

func (c *Controller) InitRoutes(server *echo.Echo) {
	// Swagger init
	//server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group("/api")
	{
		// song routes
		v1 := api.Group("/v1")
		{
			v1.GET("/login", c.Login)
			v1.GET("/get-access-token", c.GetAccessToken)

		}
	}
}
