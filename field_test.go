package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// With Field
func TestWithField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "deefeeL96").Info("Hello World")
	logger.WithField("name", "jalal").
		WithField("username", "defeeL96").Info("Hello World")
}

// With Fielss
func TestWithFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "defeeL96",
		"name":     "jalal",
	}).Info("Hello World")
}
