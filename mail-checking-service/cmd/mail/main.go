package main

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	myApp, err := app.New(ctx)
	if err != nil {
		logger.Fatal("failed to init app", zap.Error(err))
	}

	err = myApp.Run(ctx)
	if err != nil {
		logger.Fatal("failed to run app", zap.Error(err))
	}
}
