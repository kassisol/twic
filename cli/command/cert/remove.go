package cert

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/pkix"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/user"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/twic/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/kassisol/twic/storage/driver"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove Docker certificate",
		Long:    removeDescription,
		Run:     runRemove,
	}

	flags := cmd.Flags()
	flags.StringVarP(&tsaUsername, "username", "u", "", "Username")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	var username string
	var password string

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

	if len(tsaUsername) <= 0 {
		username = readinput.ReadInput("Username")
	} else {
		username = tsaUsername
	}

	if len(tsaPassword) <= 0 {
		password = readinput.ReadPassword("Password")
	} else {
		password = tsaPassword
	}

	cert := s.GetCert(args[0])

	if cert.Type == "client" {
		username = cert.CN
	}

	// Input validations
	// IV - Check if name already exists
	if cert == (driver.CertResult{}) {
		log.Fatal("Name, ", args[0], ", does not exist")
	}

	// IV - Type
	if cert.Type == "engine" {
		user, err := user.New()
		if err != nil {
			log.Fatal(err)
		}

		if !user.IsRoot() {
			log.Fatal("You must be root to remove an engine certificate type")
		}
	}

	// IV - Username
	if len(username) <= 0 {
		log.Fatal("Empty username is not allowed")
	}

	// IV - Password
	if len(password) <= 0 {
		log.Fatal("Empty password is not allowed")
	}

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
	token, err := clt.GetToken(username, password)
	if err != nil {
		log.Fatal(err)
	}

	// Send Revocation Request
	cf, err := config.CertFilesName(args[0])
	if err != nil {
		log.Fatal(err)
	}

	certificate, err := pkix.NewCertificateFromPEMFile(cf.Crt)
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

	if err = os.RemoveAll(cf.Dir); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove Docker certificate

`
