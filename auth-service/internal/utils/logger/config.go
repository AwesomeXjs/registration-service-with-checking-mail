package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetCore creates and returns a zapcore.Core that writes logs to both stdout and a file.
// It configures log formats for production and development environments.
func GetCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
	})

	// Set up encoder configurations for production and development.
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // Colorizes log levels in console

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}

// GetAtomicLevel returns a zap.AtomicLevel based on the provided log level string.
// It sets the log level and logs an error if the level is invalid.
func GetAtomicLevel(logLevel *string) zap.AtomicLevel {
	var level zapcore.Level
	fmt.Println("logger level: ", *logLevel)
	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
