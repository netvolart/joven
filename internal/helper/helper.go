package helper

import "fmt"

const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
)

func HumanReadableOutdated(outdated bool) string {
	if outdated {
		return "Outdated"
	}
	return ""

}

func AddColor(outdated bool, s string) string {
	if outdated {
		return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
	}
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)

}
