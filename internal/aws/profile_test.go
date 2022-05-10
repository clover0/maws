package aws_test

import (
	"os"
	"testing"

	"maws/internal/aws"
)

func createTempProfileFile() string {
	profile := `
[default]
aws_key=xxx

[profile account.Admin]
key=xxx

[profile account.admin]
key=xxx

[profile account.Admin.45]
key=xxx

`
	file, err := os.CreateTemp("", "test")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(profile)
	return file.Name()
}

func TestFindProfiles(t *testing.T) {

	tests := []struct {
		Name           string
		FileName       string
		Filter         string
		ExpectProfiles int
	}{
		{
			Name:           "match partial",
			FileName:       createTempProfileFile(),
			Filter:         "Admin",
			ExpectProfiles: 2,
		},
		{
			Name:           "match prefix",
			FileName:       createTempProfileFile(),
			Filter:         "account",
			ExpectProfiles: 3,
		},
		{
			Name:           "not match",
			FileName:       createTempProfileFile(),
			Filter:         "not-exist",
			ExpectProfiles: 0,
		},
	}

	for _, c := range tests {
		t.Run(c.Name, func(t *testing.T) {
			target := aws.FindProfiles(c.FileName, c.Filter)
			if len(target) != c.ExpectProfiles {
				t.Errorf("shoud be %d but actual %d", c.ExpectProfiles, len(target))
			}
		})
	}

}
