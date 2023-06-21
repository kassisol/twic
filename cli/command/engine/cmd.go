package engine

import (
	"github.com/spf13/cobra"
)

var (
	certType     string
	certCN       string
	certAltNames string
	duration     int

	tsaURL      string
	tsaToken    string
	tsaUsername string
	tsaPassword string
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "engine",
		Short: "Manage Docker engine certificate",
		Long:  engineDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newCreateCommand(),
		newInfoCommand(),
		newRemoveCommand(),
		newRenewCommand(),
	)

	return cmd
}

var engineDescription = `
The **twic engine** command has subcommands for managing Docker engine certificate.

To see help for a subcommand, use:

    twic engine [command] --help

For full details on using twic engine visit Harbormaster's online documentation.

`
