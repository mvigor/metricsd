package utils

import (
	"go.uber.org/zap"
)

var logger *zap.Logger = nil

func GetLogger() *zap.Logger {
	var err error

	if logger == nil {
		logger, err = zap.NewProduction()
		if err != nil {
			// вызываем панику, если ошибка
			panic("cannot initialize zap")
		}
	}

	return logger
}
