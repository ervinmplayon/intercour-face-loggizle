package logger

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
	l *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	logger := logrus.New()
	return &LogrusLogger{l: logger}
}

func (l *LogrusLogger) Info(msg string) {
	l.l.Info(msg)
}

func (l *LogrusLogger) Error(msg string) {
	l.l.Info(msg)
}

func (l *LogrusLogger) Debug(msg string) {
	l.l.Info(msg)
}
