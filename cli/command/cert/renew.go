package cert

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/spf13/cobra"
)

func newRenewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew [name]",
		Short: "Renew Docker certificate",
		Long:  renewDescription,
		Run:   runRenew,
	}

	return cmd
}

func runRenew(cmd *cobra.Command, args []string) {
	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	utils.Exit(fmt.Errorf("Not implemented yet"))
}

var renewDescription = `
Renew Docker certificate

`
