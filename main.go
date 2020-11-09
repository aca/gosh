package main

import (
	"os"

	"github.com/aca/gosh/cmds/gostrings"
	"github.com/spf13/cobra"
)

func main() {
	cmdRoot := &cobra.Command{
		Use:          "gosh",
		SilenceUsage: true,
	}
	cmdRoot.AddCommand(gostrings.CmdGostrings)
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}
