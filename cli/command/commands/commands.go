package commands

import (
	"github.com/kassisol/twic/cli/command/cert"
	"github.com/kassisol/twic/cli/command/profile"
	"github.com/kassisol/twic/cli/command/system"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		cert.NewCommand(),
		profile.NewCommand(),
		system.NewVersionCommand(),
	)
}
