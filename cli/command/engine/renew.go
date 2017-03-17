package engine

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/user"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRenewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew [name]",
		Short: "Renew Docker engine certificate",
		Long:  renewDescription,
		Run:   runRenew,
	}

	return cmd
}

func runRenew(cmd *cobra.Command, args []string) {
	user, err := user.New()
	if err != nil {
		log.Fatal(err)
	}

	if !user.IsRoot() {
		log.Fatal("You must be root to run engine subcommand")
	}

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	utils.Exit(fmt.Errorf("Not implemented yet"))
}

var renewDescription = `
Renew Docker engine certificate

`
