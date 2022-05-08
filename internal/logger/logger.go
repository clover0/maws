package logger

import "fmt"

type Logger struct {
	isDebug bool
}

func NewLogger(mode string) (l Logger) {
	if mode == "1" {
		l.isDebug = true
	}
	return l
}

func (l Logger) Debug(format string, a ...any) {
	if l.isDebug {
		fmt.Printf(format, a...)
	}
}

func (l Logger) Info(format string, a ...any) {
	fmt.Printf(format, a...)
}
