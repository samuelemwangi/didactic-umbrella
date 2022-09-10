package logging

import "go.uber.org/zap"

func SetUpLogger() *zap.Logger {

	logger, _ := zap.NewProduction()

	return logger
}
