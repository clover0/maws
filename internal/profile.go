package internal

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const AWS_CONFIG_FILE_NAME = "config"
const AWS_PROFILE_NAME_FORMAT_REGEXP = `\[profile (.*)\]`

func FindProfiles(filter string) []string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(home + "/.aws/" + AWS_CONFIG_FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	profiles := make([]string, 0)
	re := regexp.MustCompile(AWS_PROFILE_NAME_FORMAT_REGEXP)
	for scanner.Scan() {
		l := scanner.Text()
		matched := re.FindStringSubmatch(l)
		if len(matched) > 1 {
			profiles = append(profiles, matched[1])
		}

	}

	filtered := make([]string, 0)
	for _, l := range profiles {
		if strings.Contains(l, filter) {
			filtered = append(filtered, l)
		}
	}
	return filtered
}
