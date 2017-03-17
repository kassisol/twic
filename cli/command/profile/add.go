package profile

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/validation"
	"github.com/juliengk/stack/client"
	"github.com/kassisol/twic/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/spf13/cobra"
)

var (
	certName     string
	dockerScheme string
	dockerHost   string
	dockerPort   string
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add Docker profile",
		Long:  addDescription,
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&certName, "cert-name", "c", "", "Certificate Name")
	flags.StringVarP(&dockerScheme, "docker-scheme", "s", "tcp", "Docker Scheme")
	flags.StringVarP(&dockerHost, "docker-host", "a", "", "Docker Host")
	flags.StringVarP(&dockerPort, "docker-port", "p", "2376", "Docker Port")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	var certname string
	var dockerscheme string
	var dockerhost string
	var dockerport string

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if len(certName) <= 0 {
		certname = readinput.ReadInput("Certificate Name")
	} else {
		certname = certName
	}

	if len(dockerScheme) <= 0 {
		dockerscheme = readinput.ReadInput("Docker Scheme")
	} else {
		dockerscheme = dockerScheme
	}

	if len(dockerHost) <= 0 {
		dockerhost = readinput.ReadInput("Docker Host")
	} else {
		dockerhost = dockerHost
	}

	if len(dockerPort) <= 0 {
		dockerport = readinput.ReadInput("Docker Port")
	} else {
		dockerport = dockerPort
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

	if err = validation.IsValidName(args[0]); err != nil {
		log.Fatal(err)
	}

	// Get Cert Info
	cert := s.GetCert(certname)

	if cert.Name == "" {
		log.Fatal("Certificate name is not valid")
	}

	if cert.Type == "engine" {
		log.Fatal("Engine certificate type cannot be added to profile")
	}

	dockerurl := fmt.Sprintf("%s://%s:%s", dockerscheme, dockerhost, dockerport)
	u, err := client.ParseUrl(dockerurl)
	if err != nil {
		log.Fatal(err)
	}
	if u.Scheme != "tcp" {
		log.Fatal("Docker host scheme should be tcp")
	}

	s.AddProfile(args[0], certname, dockerurl)
}

var addDescription = `
Add Docker profile

`
