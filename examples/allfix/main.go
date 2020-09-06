package main

import (
	"fixer"
	"os"
	"os/exec"
	"strings"
)

func main() {
	prefixer := fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "pinging   "
		},
		InfixFunc: func(s string) string {
			return "< " + s + " >"
		},
		SuffixFunc: func(s string) string {
			padding := ""
			if amount := 80 - len(s); amount > 0 {
				padding = strings.Repeat(" ", amount)
			}
			return padding + "ponging"
		},
	}
	cmd := exec.Command("ping", "google.com")
	cmd.Stdout = prefixer
	cmd.Run()
}
