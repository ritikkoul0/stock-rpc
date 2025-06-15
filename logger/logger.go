package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func NewLogger(appEnv string) *zap.SugaredLogger {
	var zapLogger *zap.Logger
	var err error

	if appEnv == "production" {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("Failed to initialize zap logger: " + err.Error())
	}

	logger = zapLogger.Sugar()
	return logger
}

func SyncLogger() {
	if logger != nil {
		logger.Sync()
	}
}

// Logging wrappers...
func Info(args ...interface{})                    { logger.Info(args...) }
func Debug(args ...interface{})                   { logger.Debug(args...) }
func Error(args ...interface{})                   { logger.Error(args...) }
func Fatal(args ...interface{})                   { logger.Fatal(args...) }
func Infof(template string, args ...interface{})  { logger.Infof(template, args...) }
func Debugf(template string, args ...interface{}) { logger.Debugf(template, args...) }
func Errorf(template string, args ...interface{}) { logger.Errorf(template, args...) }
func Fatalf(template string, args ...interface{}) { logger.Fatalf(template, args...) }

func SetLogger(customLogger *zap.SugaredLogger) {
	logger = customLogger
}
