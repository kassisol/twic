package engine

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/helpers"
	"github.com/juliengk/go-cert/pkix"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/filedir"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/user"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/tsa/pkg/adf"
	"github.com/kassisol/twic/pkg/cert"
	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Docker engine certificate",
		Long:  createDescription,
		Run:   runCreate,
	}

	flags := cmd.Flags()

	flags.StringVarP(&certCN, "common-name", "n", "", "Certificate Common Name")
	flags.StringVarP(&certAltNames, "alt-names", "a", "", "Certificate Alternative Names")

	flags.StringVarP(&tsaURL, "tsa-url", "c", "", "TSA URL")
	flags.StringVarP(&tsaToken, "token", "t", "", "Token")
	flags.StringVarP(&tsaUsername, "username", "u", "", "Username")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) {
	var certcn string
	var certaltnames string

	var tsaurl string
	var username string
	var password string

	go utils.RecoverFunc()

	u := user.New()
	if !u.IsRoot() {
		log.Fatal("You must be root to run engine subcommand")
	}

	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	certtype := "engine"

	if len(certCN) <= 0 {
		certcn = readinput.ReadInput("Common Name (CN)")
	} else {
		certcn = certCN
	}

	if len(certAltNames) <= 0 {
		certaltnames = readinput.ReadInput("Alt Names")
	} else {
		certaltnames = certAltNames
	}

	if len(tsaURL) <= 0 {
		tsaurl = readinput.ReadInput("TSA URL")
	} else {
		tsaurl = tsaURL
	}

	if len(tsaToken) == 0 {
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
	}

	cfg := adf.NewEngine()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	// Input validations
	if len(tsaToken) == 0 {
		// IV - Username
		if len(username) <= 0 {
			log.Fatal("Empty username is not allowed")
		}

		// IV - Password
		if len(password) <= 0 {
			log.Fatal("Empty password is not allowed")
		}
	}

	// Create cert name directory
	defer func() {
		if r := recover(); r != nil {
			if filedir.FileExists(cfg.TLS.CaFile) {
				if err := os.Remove(cfg.TLS.CaFile); err != nil {
					log.Fatal(err)
				}
			}
			if filedir.FileExists(cfg.TLS.KeyFile) {
				if err := os.Remove(cfg.TLS.KeyFile); err != nil {
					log.Fatal(err)
				}
			}

			log.Fatal(r)
		}
	}()

	clt, err := client.New(tsaurl)
	if err != nil {
		panic(err)
	}

	// Get TSA URL directories
	err = clt.GetDirectory()
	if err != nil {
		panic(err)
	}

	// Authz
	token := tsaToken
	if len(tsaToken) == 0 {
		token, err = clt.GetToken(username, password, 5)
		if err != nil {
			panic(err)
		}
	}

	// Get CA public Key
	caCrt, err := clt.GetCACertificate()
	if err != nil {
		panic(err)
	}

	err = pkix.ToPEMFile(cfg.TLS.CaFile, []byte(caCrt), 0444)
	if err != nil {
		panic(err)
	}

	// Certificate
	// -- Client Part --
	// Key pair
	key, err := helpers.CreateKey(4096, cfg.TLS.KeyFile)
	if err != nil {
		panic(err)
	}

	// CSR
	caCertificate, err := pkix.NewCertificateFromPEM([]byte(caCrt))
	if err != nil {
		panic(err)
	}

	ca := caCertificate.Crt.Subject
	ou := cert.GetOU(ca.OrganizationalUnit[0])

	ans := utils.CreateSlice(certaltnames, ",")

	if !utils.StringInSlice(certcn, ans, true) {
		ans = append(ans, certcn)
	}

	csr, err := helpers.CreateCSR(ca.Country[0], ca.Province[0], ca.Locality[0], ca.Organization[0], ou, certcn, "", ans, key)
	if err != nil {
		panic(err)
	}

	// Send CSR
	cert, err := clt.GetCertificate(token, certtype, csr.Bytes, 12)
	if err != nil {
		panic(err)
	}

	// Save Certificate
	err = pkix.ToPEMFile(cfg.TLS.CrtFile, []byte(cert), 0444)
	if err != nil {
		panic(err)
	}

	fmt.Println("Docker engine certificates created in the directory", cfg.CertsDir, ".")

	fmt.Printf("\nTo configure the Docker Daemon, add the following parameters:\n\n--tlsverify --tlscacert=%s --tlskey=%s --tlscert=%s -H tcp://0.0.0.0:2376\n\n", cfg.TLS.CaFile, cfg.TLS.KeyFile, cfg.TLS.CrtFile)
}

var createDescription = `
Create Docker engine certificate

`
