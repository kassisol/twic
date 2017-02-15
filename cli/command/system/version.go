package system

import (
	"fmt"

	"github.com/kassisol/twic/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the TWIC version information",
		Long:  versionDescription,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("TWIC", version.Version, "-- HEAD")
		},
	}

	return cmd
}

var versionDescription = `
All software has versions. This is TWIC's

`
