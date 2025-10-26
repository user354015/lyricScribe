package shared

import (
	"log"
	"os"
)

var (
	DebugEnabled bool
	logger       *log.Logger
	logFile      *os.File
)

func InitLogger(debug bool) {
	DebugEnabled = debug
	logger = log.New(os.Stdout, "", log.Ltime)

	l, _ := os.UserConfigDir()
	l += "/muse/muse.log"

	f, err := os.OpenFile(l, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logFile = f

	logger.SetOutput(f)

}

func StopLogger() {
	logFile.Close()
}

func Debug(format string, args ...interface{}) {
	if DebugEnabled {
		logger.Printf("[DEBUG] "+format, args...)
	}
}

func Info(format string, args ...interface{}) {
	logger.Printf("[INFO] "+format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Printf("[WARN] "+format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Printf("[ERROR] "+format, args...)
}
