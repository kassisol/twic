package cert

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-cert/helpers"
	"github.com/juliengk/go-cert/pkix"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/user"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/tsa/client"
	"github.com/kassisol/twic/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/kassisol/twic/storage/driver"
	"github.com/spf13/cobra"
)

var (
	certType     string
	certCN       string
	certAltNames string

	tsaURL      string
	tsaUsername string
	tsaPassword string
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add Docker certificate",
		Long:  addDescription,
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&certType, "type", "t", "client", "Certificate type")
	flags.StringVarP(&certCN, "common-name", "n", "", "Certificate Common Name")
	flags.StringVarP(&certAltNames, "alt-names", "a", "", "Certificate Alternative Names")

	flags.StringVarP(&tsaURL, "tsa-url", "c", "", "TSA URL")
	flags.StringVarP(&tsaUsername, "username", "u", "", "Username")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func getOU(ou string) string {
	words := []string{
		"Certificate",
		"Authority",
	}

	oldou := strings.Split(ou, " ")

	if len(oldou) > 1 {
		newou := []string{}

		for _, word := range oldou {
			if !utils.StringInSlice(word, words, true) {
				newou = append(newou, word)
			}
		}

		if len(newou) > 0 {
			return strings.Join(newou, " ")
		}
	}

	return ou
}

func runAdd(cmd *cobra.Command, args []string) {
	var certtype string
	var certcn string
	var certaltnames string

	var tsaurl string
	var username string
	var password string

	go utils.RecoverFunc()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	name := args[0]

	if len(certType) <= 0 {
		certtype = readinput.ReadInput("Type")
	} else {
		certtype = certType
	}

	if certtype == "engine" {
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
	}

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

	if len(tsaPassword) <= 0 {
		password = readinput.ReadPassword("Password")
	} else {
		password = tsaPassword
	}

	if certtype == "client" {
		certcn = username
	}

	config, err := adf.New()
	if err != nil {
		log.Fatal(err)
	}

	// DB
	s, err := storage.NewDriver("sqlite", config.DBFileName())
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

	// IV - type
	if certtype != "engine" && certtype != "client" {
		log.Fatal("Type is not correct")
	}

	if certtype == "engine" {
		user, err := user.New()
		if err != nil {
			log.Fatal(err)
		}

		if !user.IsRoot() {
			log.Fatal("You must be root to add an engine certificate type")
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

	// Create cert name directory
	cf, err := config.CertFilesName(name)
	if err != nil {
		log.Fatal(err)
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
	token, err := clt.GetToken(username, password)
	if err != nil {
		log.Fatal(err)
	}

	// Get CA public Key
	caCrt, err := clt.GetCACertificate()
	if err != nil {
		log.Fatal(err)
	}

	err = pkix.ToPEMFile(cf.Ca, []byte(caCrt), 0444)
	if err != nil {
		log.Fatal(err)
	}

	// Certificate
	// -- Client Part --
	// Key pair
	key, err := helpers.CreateKey(4096, cf.Key)
	if err != nil {
		log.Fatal(err)
	}

	// CSR
	caCertificate, err := pkix.NewCertificateFromPEM([]byte(caCrt))
	if err != nil {
		log.Fatal(err)
	}

	ca := caCertificate.Crt.Subject
	ou := getOU(ca.OrganizationalUnit[0])

	ans := utils.CreateSlice(certaltnames, ",")

	csr, err := helpers.CreateCSR(ca.Country[0], ca.Province[0], ca.Locality[0], ca.Organization[0], ou, certcn, "", ans, key)
	if err != nil {
		log.Fatal(err)
	}

	// Send CSR
	cert, err := clt.GetCertificate(token, certtype, csr.Bytes, 12)
	if err != nil {
		if err = os.RemoveAll(cf.Dir); err != nil {
			log.Fatal(err)
		}

		log.Fatal(err)
	}

	// Save Certificate
	err = pkix.ToPEMFile(cf.Crt, []byte(cert), 0444)
	if err != nil {
		log.Fatal(err)
	}

	// Add data to DB
	s.AddCert(name, certtype, certcn, certaltnames, tsaurl)
}

var addDescription = `
Add Docker certificate

`
