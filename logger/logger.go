package logger

import (
	"log"

	"go.uber.org/zap"
)

var (
	sugar *zap.SugaredLogger
)

func Init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	sugar = logger.Sugar()
}

func Debugf(format string, args ...interface{}) {
	sugar.Debugf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	sugar.Errorf(format, args...)
}
