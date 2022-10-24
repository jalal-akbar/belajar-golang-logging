package belajar_golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Message, entry.Level)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Info")
	logger.Debug("Debug")
	logger.Warn("Warn")
	logger.Error("Error")
}
