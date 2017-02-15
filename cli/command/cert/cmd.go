package cert

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cert",
		Short: "Manage Docker certificates",
		Long:  certDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newAddCommand(),
		newListCommand(),
		newRemoveCommand(),
		newRenewCommand(),
	)

	return cmd
}

var certDescription = `
The **twic cert** command has subcommands for managing Docker certificates.

To see help for a subcommand, use:

    twic cert [command] --help

For full details on using twic cert visit Harbormaster's online documentation.

`
