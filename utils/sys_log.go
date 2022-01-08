package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	if os.Getenv("CLOSE_LOG") == "true" {
		logrus.SetOutput(&EmptyLogger{})
	}
}

type EmptyLogger struct{}

func (e *EmptyLogger) Write(data []byte) (int, error) {
	return 0, nil
}
