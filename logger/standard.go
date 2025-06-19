package logger

import (
	"log"
)

type StandardLogger struct{}

func (l *StandardLogger) Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

func (l *StandardLogger) Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}

func (l *StandardLogger) Debug(msg string) {
	log.Printf("[DEBUG] %s", msg)
}

func (l *StandardLogger) Fatal(msg string) {
	log.Fatalf("[FATAL] %s", msg)
}
