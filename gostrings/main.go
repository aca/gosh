package main

import (
	"os"

	"github.com/aca/gosh/cmds/gostrings"
)

func main() {
	if err := gostrings.CmdGostrings.Execute(); err != nil {
		os.Exit(1)
	}
}
