package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
