package main

import (
	"fixer"
	"os"
	"os/exec"
)

func main() {
	prefixer := fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "pinging --> "
		},
	}
	cmd := exec.Command("ping", "google.com")
	cmd.Stdout = prefixer
	cmd.Run()
}
