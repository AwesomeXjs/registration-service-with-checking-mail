package controller

import (
	_ "github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/docs" // swagger docs
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// InitRoutes initializes all the routes for the application.
func (c *Controller) InitRoutes(server *echo.Echo) {
	// Swagger init
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group("/api")
	{
		// Auth routes
		v1 := api.Group("/v1")
		{
			v1.POST("/login", c.Login)
			v1.GET("/get-access-token", c.GetAccessToken)
			v1.POST("/register", c.Registration)
			v1.GET("/validate-token", c.ValidateToken)
			v1.PATCH("/update-password", c.UpdatePassword)
		}
	}
}
