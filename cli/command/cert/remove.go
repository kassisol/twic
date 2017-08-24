package cert

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/pkix"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/user"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/tsa/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/kassisol/twic/storage/driver"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove Docker client certificate",
		Long:    removeDescription,
		Run:     runRemove,
	}

	flags := cmd.Flags()

	flags.StringVarP(&tsaToken, "token", "t", "", "Token")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	var password string

	u := user.New()

	if u.IsRoot() {
		log.Fatal("You must not be root to add an client certificate type")
	}

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	cfg := adf.NewClient()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	cfg.SetName(args[0])

	s, err := storage.NewDriver("sqlite", cfg.App.Dir.Root)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if len(tsaToken) == 0 {
		if len(tsaPassword) <= 0 {
			password = readinput.ReadPassword("Password")
		} else {
			password = tsaPassword
		}
	}

	cert := s.GetCert(args[0])

	// Input validations
	// IV - Check if name already exists
	if cert == (driver.CertResult{}) {
		log.Fatal("Name, ", args[0], ", does not exist")
	}

	// IV - Password
	if len(tsaToken) == 0 {
		if len(password) <= 0 {
			log.Fatal("Empty password is not allowed")
		}
	}

	username := cert.CN

	clt, err := client.New(cert.TSAURL)
	if err != nil {
		log.Fatal(err)
	}

	// Get TSA URL directories
	err = clt.GetDirectory()
	if err != nil {
		log.Fatal(err)
	}

	// Authz
	token := tsaToken
	if len(tsaToken) == 0 {
		token, err = clt.GetToken(username, password, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Send Revocation Request
	certificate, err := pkix.NewCertificateFromPEMFile(cfg.TLS.CrtFile)
	if err != nil {
		log.Fatal(err)
	}

	err = clt.RevokeCertificate(token, int(certificate.Crt.SerialNumber.Int64()))
	if err != nil {
		log.Fatal(err)
	}

	// Once done remove entry from DB
	if err = s.RemoveCert(args[0]); err != nil {
		log.Fatal(err)
	}

	if err = os.RemoveAll(cfg.Profile.CertDir); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove Docker client certificate

`
