package log

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	log, _ = zap.NewDevelopment()
	log.Info("init log")
}

func Log() *zap.Logger {
	return log
}
