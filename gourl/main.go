package main

import (
	"os"

	"github.com/aca/gosh/cmds/gourl"
)

func main() {
	if err := gourl.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
