package cert

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/pkix"
	"github.com/kassisol/twic/pkg/adf"
	"github.com/kassisol/twic/pkg/date"
	"github.com/kassisol/twic/storage"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List Docker client certificates",
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

	config := adf.New("client")

	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewDriver("sqlite", config.DBFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	certs := s.ListCerts()

	if len(certs) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tTYPE\tCN\tTSA URL\tEXPIRE")

		for _, c := range certs {
			var expire string

			config.SetName(c.Name)

			cf, err := config.CertFilesName()
			if err != nil {
				log.Fatal(err)
			}

			certificate, err := pkix.NewCertificateFromPEMFile(cf.Crt)
			if err == nil {
				expire = date.ExpireDateString(certificate.Crt.NotAfter)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", c.Name, c.Type, c.CN, c.TSAURL, expire)
		}

		w.Flush()
	}
}

var listDescription = `
List Docker client certificates

`
