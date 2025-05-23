package logger

type logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}
