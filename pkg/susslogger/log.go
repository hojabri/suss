package susslogger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var (
	logger  *logrus.Logger
)

func init() {
	newLogger()
}

// newLogger Function
func newLogger() {
	// Initialize Log as New Logrus Logger
	logger = logrus.New()
	
	// Set Text Format
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: false, FullTimestamp: true, TimestampFormat: time.RFC3339Nano})
	logger.SetReportCaller(false)
	logger.SetOutput(os.Stdout)
	
	// Set Log Level
	switch strings.ToLower(os.Getenv("CONFIG_LOG_LEVEL")) {
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	default:
		// Default INFO
		logger.SetLevel(logrus.InfoLevel)
	}
}


func Log() *logrus.Logger {
	return logger
}



