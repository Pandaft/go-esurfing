package logger

import (
	"github.com/charmbracelet/log"
	"os"
)

var Debug bool

func GetLogger(prefix string) *log.Logger {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Prefix:          prefix,
		ReportTimestamp: true,
	})
	if Debug {
		logger.SetLevel(log.DebugLevel)
	}
	return logger
}
