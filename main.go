package main

import (
	"os"

	"github.com/aca/gosh/cmds/gofilepath"
	"github.com/aca/gosh/cmds/gonet"
	"github.com/aca/gosh/cmds/gostrings"
	"github.com/aca/gosh/cmds/gourl"
	"github.com/aca/gosh/utils"
	"github.com/spf13/cobra"
)

func main() {
	cmdRoot := &cobra.Command{
		Use:          "gosh",
		SilenceUsage: true,
	}
	cmdRoot.AddCommand(gostrings.Cmd)
	cmdRoot.AddCommand(gofilepath.Cmd)
	cmdRoot.AddCommand(gonet.Cmd)
	cmdRoot.AddCommand(gourl.Cmd)

	cmdRoot.AddCommand(utils.NewCompletionCommand("gosh"))

	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}
