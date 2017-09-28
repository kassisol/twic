package engine

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/user"
	"github.com/spf13/cobra"
)

func newRenewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew",
		Short: "Renew Docker engine certificate",
		Long:  renewDescription,
		Run:   runRenew,
	}

	return cmd
}

func runRenew(cmd *cobra.Command, args []string) {
	u := user.New()

	if !u.IsRoot() {
		log.Fatal("You must be root to run engine subcommand")
	}

	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	log.Fatal("Not implemented yet")
}

var renewDescription = `
Renew Docker engine certificate

`
