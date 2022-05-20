package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var logger *zap.Logger

func InitLoggerServer(appName string, isDevelopment bool, level string) (err error) {
	zLevel := new(zapcore.Level)
	if err = zLevel.UnmarshalText([]byte(level)); err != nil {
		return
	}
	logger = NewLogger(
		SetAppName(appName),
		SetDevelopment(isDevelopment),
		SetLevel(*zLevel),
	)
	return
}

func GetLogger() *zap.Logger {
	return logger
}