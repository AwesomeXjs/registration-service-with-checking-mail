package configs

import (
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// LoadEnv loads environment variables from a specified .env file.
// It returns an error if the file cannot be loaded.
func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		logger.Error("Error loading .env file", zap.String("path", path))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}
