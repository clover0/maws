package command

import "flag"

var UsageMessage = HelpMessage

func HelpMessage() {
	flag.Usage()
}
