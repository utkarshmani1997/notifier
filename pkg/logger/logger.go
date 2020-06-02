package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    &logrus.TextFormatter{},
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}
