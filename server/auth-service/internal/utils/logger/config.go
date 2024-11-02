package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func GetCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	// настраиваем куда мы будем писать логи
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
	})

	// берем за основу дефолтный прод конфиг и его настраиваем
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// делаем еще один конфиг
	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // делает цветным уровень логирования в логе

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}

// getAtomicLevel - парсим с входных параметров при запуске лог левел и приводим его к нужному значению
func GetAtomicLevel(logLevel *string) zap.AtomicLevel {
	var level zapcore.Level
	fmt.Println("logger level: ", *logLevel)
	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
