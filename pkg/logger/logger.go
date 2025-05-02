package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func InitLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
	})
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)

	switch strings.ToLower(logLevel) {
	case "warn", "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}
