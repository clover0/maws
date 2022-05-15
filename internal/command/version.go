package command

import "fmt"

const VERSION = "0.0.3"

func VersionMessage() {
	out := fmt.Sprintf(`
MAWS version is %s
`, VERSION)
	fmt.Printf(out)
}
