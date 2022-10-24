package belajar_golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logging") //print time,level,msg
	fmt.Println("Hello Logging")    // print Hello Logging
}
