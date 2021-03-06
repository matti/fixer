package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/matti/fixer"
)

func main() {
	prefixer := fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "pinging   "
		},
		InfixFunc: func(s string) string {
			if len(s) > 25 {
				return "< " + s[:25] + " ... >"
			}

			return "< " + s + " >"
		},
		SuffixFunc: func(s string) string {
			padding := ""
			if amount := 36 - len(s); amount > 0 {
				padding = strings.Repeat(" ", amount)
			}
			return padding + "ponging"
		},
	}
	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = prefixer
	cmd.Run()
}
