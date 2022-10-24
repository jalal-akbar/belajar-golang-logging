package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	formatter := &logrus.JSONFormatter{}

	logger.SetFormatter(formatter)
	// Entry
	entry := logrus.NewEntry(logger)
	var fields = logrus.Fields{
		"username": "defeeL96",
	}
	entry.WithFields(fields)
}
