package main

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/api-gateway-auth/internal/utils/logger"
	"go.uber.org/zap"
)

// @title Authentication API
// @version 1.0
// @description API Server for Authentication
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()

	myApp, err := app.New(ctx)
	if err != nil {
		logger.Fatal("failed to init app", zap.Error(err))
	}

	err = myApp.Run()
	if err != nil {
		logger.Fatal("failed to run app", zap.Error(err))
	}
}
