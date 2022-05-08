package internal

import "fmt"

type Aggregator struct {
	profiles []string
	cmdArgs  []string
	logger   Logger
}

func NewAggregator(profiles, cmdArgs []string, logger Logger) Aggregator {
	return Aggregator{profiles: profiles, cmdArgs: cmdArgs, logger: logger}
}

type status int

const (
	SUCCESS = status(iota)
	FAIL
)

type message struct {
	status
	result string
}

func (a *Aggregator) Do() {

	stream := make(chan message)
	for _, p := range a.profiles {
		go func(args []string, prof string, logger Logger) {
			c := NewAWSCommand(args, logger, prof)
			r := prof + "---\n"
			if err := c.Exec(); err != nil {
				stream <- message{
					status: FAIL,
					result: r + "(fail)" + c.Output(),
				}
				return
			}
			stream <- message{
				status: SUCCESS,
				result: r + c.Output(),
			}
		}(a.cmdArgs, p, a.logger)
	}

	finished := 0
	for {
		select {
		case s := <-stream:
			switch s.status {
			case SUCCESS:
				fmt.Println(s.result)
				finished += 1
			case FAIL:
				// TODO:
				fmt.Println(s.result)
				finished += 1
			}
		}
		if finished >= len(a.profiles) {
			break
		}
	}
	return
}
