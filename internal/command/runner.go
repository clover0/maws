package command

import (
	"fmt"

	"maws/internal/holder"
	"maws/internal/logger"
)

type Runner interface {
	Do() error
}

type runner struct {
	profiles   []string
	cmdArgs    []string
	cmdBuilder func(args []string, logger logger.Logger, profile string) ICommand
	logger     logger.Logger
	reporter   Reporter
	holder     holder.Holder
}

func NewRunner(profiles, cmdArgs []string, logger logger.Logger, reporter Reporter, out OutputFormat) Runner {
	if out == OutText {
		panic(fmt.Errorf("not implement other than json output format"))
	}
	// TODO: implement output format for text
	return &runner{
		profiles:   profiles,
		cmdArgs:    cmdArgs,
		cmdBuilder: NewAWSCommand,
		logger:     logger,
		reporter:   reporter,
		holder:     holder.NewJsonHolder(),
	}
}

type status int

const (
	SUCCESS = status(iota)
	FAIL
)

type message struct {
	status
	profile string
	result  string
}

func (a *runner) Do() error {
	if len(a.profiles) == 0 {
		return nil
	}

	stream := make(chan message)
	for _, p := range a.profiles {
		go func(args []string, prof string, logger logger.Logger) {
			c := a.cmdBuilder(args, logger, prof)
			if err := c.Exec(); err != nil {
				stream <- message{
					status:  FAIL,
					profile: prof,
					result:  c.Output(),
				}
				return
			}
			stream <- message{
				status:  SUCCESS,
				profile: prof,
				result:  c.Output(),
			}
		}(a.cmdArgs, p, a.logger)
	}

	finished := 0
	for {
		select {
		case s := <-stream:
			switch s.status {
			case SUCCESS:
				a.holder.Add(s.profile, s.result)
				finished += 1
			case FAIL:
				a.reporter.OutputErr(a.decorateFail(s.profile, s.result))
				finished += 1
			}
		}
		if finished == 0 || finished >= len(a.profiles) {
			break
		}
	}
	a.reporter.Output(a.holder.OutAll())
	return nil
}

func (a *runner) decorateSuccess(profile, message string) string {
	return profile + "---\n" + message
}

func (a *runner) decorateFail(profile, message string) string {
	return profile + " fail\n" + message
}
