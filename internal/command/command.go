package command

import (
	"context"
	"os/exec"
	"strings"

	"maws/internal/logger"
)

//go:generate moq -out command_moq.go . ICommand

const AWS_COMMAND = "aws"

type ICommand interface {
	Output() string
	Exec() error
}

type Command struct {
	command []string
	result  string
	logger  logger.Logger
}

func NewAWSCommand(command []string, logger logger.Logger, profile string) ICommand {
	awsCmd := make([]string, 0)
	awsCmd = append(awsCmd, command...)
	awsCmd = append(awsCmd, "--profile")
	awsCmd = append(awsCmd, profile)
	return &Command{command: awsCmd, logger: logger}
}

func (c *Command) Exec() (err error) {
	cm := exec.CommandContext(context.TODO(), AWS_COMMAND, c.command...)
	c.logger.Debug("command: %s\n", strings.Join(cm.Args, " "))
	o, err := cm.CombinedOutput()
	c.result = string(o)
	return err
}

func (c *Command) Output() string {
	return c.result
}
