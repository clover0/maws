package aws

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const AWS_PROFILE_NAME_FORMAT_REGEXP = `\[profile (.*)\]`

func FindProfiles(path string, filter string) []string {
	f, err := os.Open(path)
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

	if filter == "" {
		return profiles
	}

	filtered := make([]string, 0)
	for _, l := range profiles {
		if strings.Contains(l, filter) {
			filtered = append(filtered, l)
		}
	}
	return filtered
}
