package command

import "fmt"

type Reporter interface {
	Output(text string)
	OutputErr(text string)
}

type ConsoleOutput struct {
}

func NewConsoleOutput() Reporter {
	return ConsoleOutput{}
}

func (c ConsoleOutput) Output(text string) {
	fmt.Println(text)
}

// TODO:
func (c ConsoleOutput) OutputErr(text string) {
	fmt.Println(text)
}
