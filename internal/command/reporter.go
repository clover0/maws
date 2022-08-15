package command

import "fmt"

//go:generate moq -out reporter_moq.go . Reporter

type Reporter interface {
	Output(text string)
	OutputErr(text string)
}

func NewConsoleOutput() Reporter {
	return consoleOutput{}
}

func (c consoleOutput) Output(s string) {
	fmt.Println(s)
}

func (c consoleOutput) OutputErr(s string) {
	fmt.Println(s)
}

type consoleOutput struct {
}
