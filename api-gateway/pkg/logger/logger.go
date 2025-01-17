package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

// Init initializes the global logger with the specified zapcore.Core and optional options.
// It sets up a new zap.Logger and assigns it to the global variable `globalLogger`.
func Init(core zapcore.Core, options ...zap.Option) {
	globalLogger = zap.New(core, options...)
}

// Debug logs a debug message with optional fields using the global logger.
func Debug(msg string, placeInCode string, fields ...zap.Field) {
	fields = append(fields, zap.String("placeInCode: ", placeInCode))
	globalLogger.Debug(msg, fields...)
}

// Info logs an informational message with optional fields using the global logger.
func Info(msg string, placeInCode string, fields ...zap.Field) {
	fields = append(fields, zap.String("place_in_code: ", placeInCode))
	globalLogger.Info(msg, fields...)
}

// Warn logs a warning message with optional fields using the global logger.
func Warn(msg string, placeInCode string, fields ...zap.Field) {
	fields = append(fields, zap.String("place_in_code: ", placeInCode))
	globalLogger.Warn(msg, fields...)
}

// Error logs an error message with optional fields using the global logger.
func Error(msg string, placeInCode string, fields ...zap.Field) {
	fields = append(fields, zap.String("place_in_code: ", placeInCode))
	globalLogger.Error(msg, fields...)
}

// Fatal logs a fatal error message with optional fields and then terminates the application.
func Fatal(msg string, placeInCode string, fields ...zap.Field) {
	fields = append(fields, zap.String("place_in_code: ", placeInCode))
	globalLogger.Fatal(msg, fields...)
}

// WithOptions returns a new logger with the specified options applied to the global logger.
func WithOptions(opts ...zap.Option) *zap.Logger {
	return globalLogger.WithOptions(opts...)
}
