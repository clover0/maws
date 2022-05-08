package main

import (
	"flag"
	"fmt"
	"maws/internal"
	"os"
)

func options() map[string]string {
	return map[string]string{
		"PROFILE_FILTER": "profile-filter",
	}
}

func main() {
	debugEnv := os.Getenv("MAWS_DEBUG")
	logger := internal.NewLogger(debugEnv)

	o := options()
	profileFilter := flag.String(o["PROFILE_FILTER"], "", "Keyword for filtering profiles")
	flag.Parse()

	args := flag.Args()

	profiles := internal.FindProfiles(*profileFilter)

	logger.Debug("target profiles: %s\n", profiles)
	for _, p := range profiles {

		c := internal.NewAWSCommand(args, logger, p)
		c.Exec()
		fmt.Println(c.Output())
	}
}
