package main

import (
	"flag"
	"os"

	"maws/internal/aws"
	"maws/internal/command"
	"maws/internal/logger"
)

const AWS_DEFAULT_CONFIG_PATH = "/.aws/config"

func options() map[string]string {
	return map[string]string{
		"PROFILE_FILTER": "profile-filter",
	}
}

func main() {
	debugEnv := os.Getenv("MAWS_DEBUG")
	logger := logger.NewLogger(debugEnv)

	o := options()
	profileFilter := flag.String(o["PROFILE_FILTER"], "", "Keyword for filtering profiles")
	flag.Parse()

	a := os.Args
	if len(a) == 1 {
		command.UsageMessage()
		return
	}
	if len(a) > 1 {
		if a[1] == "version" {
			command.VersionMessage()
			return
		}
		if a[1] == "help" {
			command.HelpMessage()
			return
		}
	}

	args := flag.Args()

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	p := home + AWS_DEFAULT_CONFIG_PATH
	profiles := aws.FindProfiles(p, *profileFilter)

	logger.Debug("target profiles: %s\n", profiles)
	agg := command.NewAggregator(profiles, args, logger, command.NewConsoleOutput())
	agg.Do()
}
