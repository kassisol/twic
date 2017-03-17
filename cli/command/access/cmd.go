package access

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/readinput"
	"github.com/kassisol/tsa/client"
	"github.com/spf13/cobra"
)

var (
	tsaURL      string
	tsaTTL      int
	tsaUsername string
	tsaPassword string
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access",
		Short: "Get TSA access token",
		Long:  accessDescription,
		Run:   runAccess,
	}

	flags := cmd.Flags()

	flags.StringVarP(&tsaURL, "tsa-url", "c", "", "TSA URL")
	flags.IntVarP(&tsaTTL, "ttl", "t", 1440, "Token TTL")
	flags.StringVarP(&tsaUsername, "username", "u", "", "Username")
	flags.StringVarP(&tsaPassword, "password", "p", "", "Password")

	return cmd
}

func runAccess(cmd *cobra.Command, args []string) {
	var tsaurl string
	var tsattl int
	var username string
	var password string

	go utils.RecoverFunc()

	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	if len(tsaURL) <= 0 {
		tsaurl = readinput.ReadInput("TSA URL")
	} else {
		tsaurl = tsaURL
	}

	tsattl = tsaTTL

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

	// Input validations
	// IV - Username
	if len(username) <= 0 {
		log.Fatal("Empty username is not allowed")
	}

	// IV - Password
	if len(password) <= 0 {
		log.Fatal("Empty password is not allowed")
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
	token, err := clt.GetToken(username, password, tsattl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
}

var accessDescription = `
The **twic access** command has subcommands for Getting TSA access token.

To see help for a subcommand, use:

    twic access [command] --help

For full details on using twic access visit Harbormaster's online documentation.

`
