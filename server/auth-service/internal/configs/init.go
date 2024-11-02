package configs

import (
	"fmt"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		logger.Error("Error loading .env file", zap.String("path", path))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}
