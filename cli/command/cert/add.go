package cert

import (
	"net/url"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/helpers"
	"github.com/juliengk/go-cert/pkix"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/user"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/tsa/pkg/adf"
	"github.com/kassisol/twic/pkg/cert"
	"github.com/kassisol/twic/storage"
	"github.com/kassisol/twic/storage/driver"
	"github.com/spf13/cobra"
)

var (
	certType     string
	certCN       string
	certAltNames string

	tsaURL      string
	tsaToken    string
	tsaUsername string
	tsaPassword string
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add Docker client certificate",
		Long:  addDescription,
		Run:   runAdd,
	}

	flags := cmd.Flags()

	flags.StringVarP(&tsaURL, "tsa-url", "c", "", "TSA URL")
	flags.StringVarP(&tsaToken, "token", "t", "", "Token")
	flags.StringVarP(&tsaUsername, "username", "u", "", "Username")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	var tsaurl string
	var username string
	var password string

	go utils.RecoverFunc()

	u := user.New()
	if u.IsRoot() {
		log.Fatal("You must not be root to add a client certificate type")
	}

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	name := args[0]

	certtype := "client"

	if len(tsaURL) <= 0 {
		tsaurl = readinput.ReadInput("TSA URL")
	} else {
		tsaurl = tsaURL
	}

	if len(tsaUsername) <= 0 {
		username = readinput.ReadInput("Username")
	} else {
		username = tsaUsername
	}

	if len(tsaToken) == 0 {
		if len(tsaPassword) <= 0 {
			password = readinput.ReadPassword("Password")
		} else {
			password = tsaPassword
		}
	}

	certcn := username

	cfg := adf.NewClient()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	cfg.SetName(name)

	// DB
	s, err := storage.NewDriver("sqlite", cfg.App.Dir.Root)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	// Input validations
	// IV - Name
	if err = validation.IsValidName(args[0]); err != nil {
		log.Fatal(err)
	}

	// Check if name already exists
	if s.GetCert(name) != (driver.CertResult{}) {
		log.Fatal("Name, ", name, ", already exists")
	}

	// IV - TSA URL
	tu, err := url.Parse(tsaurl)
	if err != nil {
		log.Fatal(err)
	}

	if tu.Scheme != "https" {
		log.Fatal("TSA URL scheme should be https")
	}

	// IV - Username
	if len(username) <= 0 {
		log.Fatal("Empty username is not allowed")
	}

	// IV - Password
	if len(tsaToken) == 0 {
		if len(password) <= 0 {
			log.Fatal("Empty password is not allowed")
		}
	}

	// Create cert name directory
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

	// Get CA public Key
	caCrt, err := clt.GetCACertificate()
	if err != nil {
		log.Fatal(err)
	}

	err = pkix.ToPEMFile(cfg.TLS.CaFile, []byte(caCrt), 0444)
	if err != nil {
		log.Fatal(err)
	}

	// Certificate
	// -- Client Part --
	// Key pair
	key, err := helpers.CreateKey(4096, cfg.TLS.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// CSR
	caCertificate, err := pkix.NewCertificateFromPEM([]byte(caCrt))
	if err != nil {
		log.Fatal(err)
	}

	ca := caCertificate.Crt.Subject
	ou := cert.GetOU(ca.OrganizationalUnit[0])

	ans := []string{}

	csr, err := helpers.CreateCSR(ca.Country[0], ca.Province[0], ca.Locality[0], ca.Organization[0], ou, certcn, "", ans, key)
	if err != nil {
		log.Fatal(err)
	}

	// Send CSR
	cert, err := clt.GetCertificate(token, certtype, csr.Bytes, 12)
	if err != nil {
		if err1 := os.RemoveAll(cfg.Profile.CertDir); err1 != nil {
			log.Fatal(err)
		}

		log.Fatal(err)
	}

	// Save Certificate
	err = pkix.ToPEMFile(cfg.TLS.CrtFile, []byte(cert), 0444)
	if err != nil {
		log.Fatal(err)
	}

	// Add data to DB
	s.AddCert(name, certtype, certcn, "", tsaurl)
}

var addDescription = `
Add Docker client certificate

`
