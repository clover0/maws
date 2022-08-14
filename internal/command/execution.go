package command

import (
	"maws/internal/logger"
)

type runner struct {
	profiles   []string
	cmdArgs    []string
	cmdBuilder func(args []string, logger logger.Logger, profile string) ICommand
	logger     logger.Logger
	reporter   Reporter
}

func NewRunner(profiles, cmdArgs []string, logger logger.Logger, reporter Reporter) runner {
	return runner{
		profiles:   profiles,
		cmdArgs:    cmdArgs,
		cmdBuilder: NewAWSCommand,
		logger:     logger,
		reporter:   reporter,
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
				a.reporter.Output(a.decorateSuccess(s.profile, s.result))
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
	return nil
}

func (a *runner) decorateSuccess(profile, message string) string {
	return profile + "---\n" + message
}

func (a *runner) decorateFail(profile, message string) string {
	return profile + " fail\n" + message
}
