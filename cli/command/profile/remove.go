package profile

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/kassisol/twic/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove Docker profile",
		Long:    removeDescription,
		Run:     runRemove,
	}

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	config, err := adf.New()
	if err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewDriver("sqlite", config.DBFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	s.RemoveProfile(args[0])
}

var removeDescription = `
Remove Docker profile

`
