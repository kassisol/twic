package profile

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Manage Docker profiles",
		Long:  profileDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newAddCommand(),
		newEnvCommand(),
		newListCommand(),
		newRemoveCommand(),
		newStatusCommand(),
	)

	return cmd
}

var profileDescription = `
The **twic profile** command has subcommands for managing Docker profiles.

To see help for a subcommand, use:

    twic profile [command] --help

For full details on using twic profile visit Harbormaster's online documentation.

`
