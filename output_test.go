package belajar_golang_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()
	// output logger adalah io.Writer
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("Hello Logging")
	logger.Debug("Hello Logging")
	logger.Info("Hello Logging")
	logger.Warn("Hello Logging")
	logger.Error("Hello Logging")
}
