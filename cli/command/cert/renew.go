package cert

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-cert/helpers"
	"github.com/juliengk/go-cert/pkix"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/tsa/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/kassisol/twic/storage/driver"
	"github.com/spf13/cobra"
)

func newRenewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew [name]",
		Short: "Renew Docker client certificate",
		Long:  renewDescription,
		Run:   runRenew,
	}

	flags := cmd.Flags()

	flags.StringVarP(&tsaToken, "token", "t", "", "Token")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runRenew(cmd *cobra.Command, args []string) {
	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	name := args[0]

	cfg := adf.NewClient()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	cfg.SetName(name)

	s, err := storage.NewDriver("sqlite", cfg.App.Dir.Root)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	crt := s.GetCert(name)
	if crt == (driver.CertResult{}) {
		log.Fatal("Name, ", name, ", does not exist")
	}

	clt, err := client.New(crt.TSAURL)
	if err != nil {
		log.Fatal(err)
	}

	err = clt.GetDirectory()
	if err != nil {
		log.Fatal(err)
	}

	oldcert, err := pkix.NewCertificateFromPEMFile(cfg.TLS.CrtFile)
	if err != nil {
		log.Fatal(err)
	}

	key, err := pkix.NewKey(4096)
	if err != nil {
		log.Fatal(err)
	}

	keyBytes, err := key.ToPEM()
	if err != nil {
		log.Fatal(err)
	}

	csr, err := helpers.CreateCSR(oldcert.Crt.Subject.Country[0], oldcert.Crt.Subject.Province[0], oldcert.Crt.Subject.Locality[0], oldcert.Crt.Subject.Organization[0], oldcert.Crt.Subject.OrganizationalUnit[0], oldcert.Crt.Subject.CommonName, "", []string{}, key)
	if err != nil {
		log.Fatal(err)
	}

	token := tsaToken
	if len(tsaToken) == 0 {
		password := tsaPassword
		if len(tsaPassword) == 0 {
			password = readinput.ReadPassword("Password")
		}
		token, err = clt.GetToken(crt.CN, password, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = clt.RevokeCertificate(token, int(oldcert.Crt.SerialNumber.Int64()))
	if err != nil {
		log.Fatal(err)
	}

	newcert, err := clt.GetCertificate(token, "client", csr.Bytes, 12)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(cfg.TLS.CrtFile)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(cfg.TLS.KeyFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	err = pkix.ToPEMFile(cfg.TLS.CrtFile, []byte(newcert), 0444)
	if err != nil {
		log.Fatal(err)
	}

	err = pkix.ToPEMFile(cfg.TLS.KeyFile, keyBytes, 0400)
	if err != nil {
		log.Fatal(err)
	}
}

var renewDescription = `
Renew Docker client certificate

`
