package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	appName   string
	logger    *log.Logger
	withDebug bool
	startTime time.Time
)

func InitLogger(name, logFile string, debug bool) {
	appName = name
	withDebug = debug
	startTime = time.Now()

	if logFile == "" {
		// will use only stdout
		return
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger = log.New(file, "", log.LstdFlags)
}

func Debugf(format string, args ...interface{}) {
	if !withDebug {
		return
	}
	format = "[debug] " + format
	if appName != "" {
		format = appName + " | " + format
	}
	format = formattedTimestamp() + format
	fmt.Printf(format+"\n", args...)
	if logger != nil {
		logger.Printf(format, args...)
	}
}

func Printf(format string, args ...interface{}) {
	if appName != "" {
		format = appName + " | " + format
	}
	format = formattedTimestamp() + format
	fmt.Printf(format+"\n", args...)
	if logger != nil {
		logger.Printf(format, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	format = "[fatal] " + format
	if appName != "" {
		format = appName + " | " + format
	}
	format = formattedTimestamp() + format
	Printf(format, args...)
	if logger != nil {
		logger.Fatalf(format, args...)
	}
	os.Exit(1)
}

func formattedTimestamp() string {
	elapsedTimeStr := ""
	if !startTime.IsZero() {
		elapsedTimeStr = fmt.Sprintf(" (%0.3f)", time.Since(startTime).Seconds())
	}
	return fmt.Sprintf("%s%s ", time.Now().Format("2006-01-02 15:04:05.000"), elapsedTimeStr)
}
