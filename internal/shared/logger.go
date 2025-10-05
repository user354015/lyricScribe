package shared

import (
	"log"
	"os"
)

var (
	DebugEnabled bool
	logger       *log.Logger
)

func InitLogger(debug bool) {
	DebugEnabled = debug
	logger = log.New(os.Stdout, "", log.Ltime)
}

func Debug(format string, args ...interface{}) {
	if DebugEnabled {
		logger.Printf("[DEBUG] "+format, args...)
	}
}

func Info(format string, args ...interface{}) {
	logger.Printf("[INFO] "+format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Printf("[ERROR] "+format, args...)
}
