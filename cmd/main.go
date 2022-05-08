package main

import (
	"flag"
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
	agg := internal.NewAggregator(profiles, args, logger)
	agg.Do()
}
