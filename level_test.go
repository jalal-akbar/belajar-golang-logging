package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	// Level
	// Secara default level hanya menampilkan level info ke atas
	logger.Trace("Ini Trace") // Tidak ada output
	logger.Debug("Ini Debug") // Tidak ada output
	logger.Info("Ini Info")
	logger.Warn("Ini Warn")
	logger.Error("Ini Error")
	// logger.Fatal("Ini Fatal") // memanggil os.Exit(1) setelah logging
	// logger.Panic("Ini Panic") // memanggil panic() setelah panic
}

func TestSetLevel(t *testing.T) {
	logger := logrus.New()
	// logger.SetLevel
	logger.SetLevel(logrus.TraceLevel) // logrus.TraceLevel untuk menetapkan dari level Trace sampai level di atasnya
	logger.Trace("Ini Trace")          // Muncul Karna Level di Set dari Trace Level
	logger.Debug("Ini Debug")          // Muncul Muncul Karna Level di Set dari Trace Level
	logger.Info("Ini Info")
	logger.Warn("Ini Warn")
	logger.Error("Ini Error")
}
