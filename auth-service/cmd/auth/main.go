package main

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"go.uber.org/zap"
)

func main() {
	const mark = "Cmd.Main"

	ctx := context.Background()

	myApp, err := app.New(ctx)
	if err != nil {
		logger.Fatal("failed to init app", mark, zap.Error(err))
	}

	err = myApp.Run(ctx)
	if err != nil {
		logger.Fatal("failed to run app", mark, zap.Error(err))
	}
}
