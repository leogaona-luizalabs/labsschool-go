package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func initLogger() {
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
}

// GetLogger ...
func GetLogger() *logrus.Entry {
	if logger == nil {
		initLogger()
	}

	hostname, _ := os.Hostname()

	return logger.WithFields(map[string]interface{}{
		"environment": "local",
		"version":     "1.0.0",
		"hostname":    hostname,
	})
}
