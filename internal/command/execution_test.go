package command_test

import (
	"testing"

	"maws/internal/command"
	"maws/internal/logger"
)

func TestAggregator_Do(t *testing.T) {
	loggerMock := &logger.LoggerMock{
		DebugFunc: func(format string, a ...any) {},
		InfoFunc:  func(format string, a ...any) {},
	}
	reporterMock := &command.ReporterMock{
		OutputFunc:    func(text string) {},
		OutputErrFunc: func(text string) {},
	}

	tests := []struct {
		Name     string
		Profiles []string
		CmdArgs  []string
	}{
		{
			Name:     "args length 0",
			Profiles: []string{"test"},
			CmdArgs:  []string{},
		},
		{
			Name:     "profile length 0",
			Profiles: []string{"test"},
			CmdArgs:  []string{"ec2"},
		},
	}

	for _, c := range tests {
		t.Run(c.Name, func(t *testing.T) {
			target := command.NewAggregator(c.Profiles, c.CmdArgs, loggerMock, reporterMock)
			result := target.Do()
			if result != nil {
				t.Errorf("error")
			}
		})
	}

}
