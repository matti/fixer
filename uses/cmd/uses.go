package main

import (
	"os"
	"os/exec"

	"github.com/matti/fixer"
)

func main() {
	f := &fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "prefix: "
		},
	}

	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = f
	cmd.Run()
}
