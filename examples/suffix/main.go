package main

import (
	"fixer"
	"os"
	"os/exec"
)

func main() {
	suffixer := fixer.Fixer{
		Writer: os.Stdout,
		SuffixFunc: func(s string) string {
			return " <-- pinging"
		},
	}
	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = suffixer
	cmd.Run()
}
