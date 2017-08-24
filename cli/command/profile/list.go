package profile

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/kassisol/tsa/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List Docker profiles",
		Long:    listDescription,
		Run:     runList,
	}

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	cfg := adf.NewClient()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewDriver("sqlite", cfg.App.Dir.Root)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	profiles := s.ListProfiles()

	if len(profiles) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tCERTIFICATE NAME\tDOCKER HOST")

		for _, p := range profiles {
			fmt.Fprintf(w, "%s\t%s\t%s\n", p.Name, p.Cert.Name, p.DockerHost)
		}

		w.Flush()
	}
}

var listDescription = `
List Docker profiles

`
