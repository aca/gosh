package main

import (
	"os"

	"github.com/aca/gosh/cmds/gonet"
)

func main() {
	if err := gonet.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
