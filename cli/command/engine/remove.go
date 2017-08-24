package engine

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/pkix"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/user"
	sclient "github.com/juliengk/stack/client"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/tsa/pkg/adf"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove Docker engine certificate",
		Long:    removeDescription,
		Run:     runRemove,
	}

	flags := cmd.Flags()

	flags.StringVarP(&tsaToken, "token", "t", "", "Token")
	flags.StringVarP(&tsaUsername, "username", "u", "", "Username")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	var username string
	var password string

	u := user.New()

	if !u.IsRoot() {
		log.Fatal("You must be root to run engine subcommand")
	}

	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	if len(tsaToken) == 0 {
		if len(tsaPassword) <= 0 {
			password = readinput.ReadPassword("Password")
		} else {
			password = tsaPassword
		}
	}

	// Input validations
	// IV - Password
	if len(tsaToken) == 0 {
		if len(username) <= 0 {
			log.Fatal("Empty username is not allowed")
		}

		if len(password) <= 0 {
			log.Fatal("Empty password is not allowed")
		}
	}

	cfg := adf.NewEngine()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	certificate, err := pkix.NewCertificateFromPEMFile(cfg.TLS.CrtFile)
	if err != nil {
		log.Fatal(err)
	}

	crldp := certificate.Crt.CRLDistributionPoints[0]

	url, err := sclient.ParseUrl(crldp)
	if err != nil {
		log.Fatal(err)
	}

	tsaurl := fmt.Sprintf("%s://%s", url.Scheme, url.Host)
	if url.Port != 443 {
		tsaurl = fmt.Sprintf("%s://%s:%d", url.Scheme, url.Host, url.Port)
	}

	clt, err := client.New(tsaurl)
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
	err = clt.RevokeCertificate(token, int(certificate.Crt.SerialNumber.Int64()))
	if err != nil {
		log.Fatal(err)
	}

	// Once done remove files
	if err = os.Remove(cfg.TLS.CaFile); err != nil {
		log.Fatal(err)
	}

	if err = os.Remove(cfg.TLS.KeyFile); err != nil {
		log.Fatal(err)
	}

	if err = os.Remove(cfg.TLS.CrtFile); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove Docker engine certificate

`
