package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	rawLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Logger = rawLogger.Sugar()
}
