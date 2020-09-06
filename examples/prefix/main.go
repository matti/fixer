package main

import (
	"os"
	"os/exec"

	"github.com/matti/fixer"
)

func main() {
	prefixer := fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "pinging --> "
		},
	}
	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = prefixer
	cmd.Run()
}
