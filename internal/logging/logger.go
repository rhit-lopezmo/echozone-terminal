package logging

import (
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger

var isFirstLog = true

func InitializeLogger() {
	logFile, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open logger file.", err)
	}

	Logger = log.New(logFile, "", log.Ltime)

	Info("Logger initialized.")
}

func logWithFirstCheck(level, msg string) {
	if isFirstLog {
		Logger.Println("=== NEW SESSION STARTED ===")
		isFirstLog = false
	}
	Logger.Printf("[%s] %s", level, msg)
}

func Debug(msg string) {
	logWithFirstCheck("DEBUG", msg)
}

func Info(msg string) {
	logWithFirstCheck("INFO", msg)
}

func Error(msg string) {
	logWithFirstCheck("ERROR", msg)
}
