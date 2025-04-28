package utils

import (
	"log"
	"os"
)

func InitLogger() {
	logDir := "./logs/backend"
	logFile := "./logs/backend/backend.log"

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(f)
}
