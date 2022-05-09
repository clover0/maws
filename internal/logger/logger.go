package logger

import "fmt"

//go:generate moq -out logger_moq.go . Logger

type Logger interface {
	Debug(format string, a ...any)
	Info(format string, a ...any)
}
type DefaultLogger struct {
	isDebug bool
}

func NewLogger(mode string) (l DefaultLogger) {
	if mode == "1" {
		l.isDebug = true
	}
	return l
}

func (l DefaultLogger) Debug(format string, a ...any) {
	if l.isDebug {
		fmt.Printf(format, a...)
	}
}

func (l DefaultLogger) Info(format string, a ...any) {
	fmt.Printf(format, a...)
}
