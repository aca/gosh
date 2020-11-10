package main

import (
	"os"

	"github.com/aca/gosh/cmds/gofilepath"
)

func main() {
	if err := gofilepath.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
