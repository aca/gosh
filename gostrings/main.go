package main

import (
	"os"

	"github.com/aca/gosh/cmds/gostrings"
)

func main() {
	if err := gostrings.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
