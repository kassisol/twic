package profile

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils/readinput"
	"github.com/juliengk/go-utils/validation"
	"github.com/juliengk/stack/client"
	"github.com/kassisol/tsa/pkg/adf"
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

	cfg := adf.NewClient()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewDriver("sqlite", cfg.App.Dir.Root)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	// Input Validations
	// IV - Profile name
	if err = validation.IsValidName(args[0]); err != nil {
		log.Fatal(err)
	}

	// IV - Cert
	cert := s.GetCert(certname)

	if len(cert.Name) == 0 {
		log.Fatal("Certificate name is not valid")
	}

	if cert.Type == "engine" {
		log.Fatal("Engine certificate type cannot be added to profile")
	}

	// IV - Docker scheme
	if dockerscheme != "tcp" {
		log.Fatal("Docker host scheme should be tcp")
	}

	// IV - Docker host
	if len(dockerhost) == 0 {
		log.Fatal("Docker host cannot be empty")
	}

	// IV - Docker port
	p,  err := strconv.Atoi(dockerport)
	if err != nil {
		log.Fatal(err)
	}
	if err = validation.IsValidPort(p); err != nil {
		log.Fatal(err)
	}

	dockerurl := fmt.Sprintf("%s://%s:%s", dockerscheme, dockerhost, dockerport)
	if _, err := client.ParseUrl(dockerurl); err != nil {
		log.Fatal(err)
	}

	s.AddProfile(args[0], certname, dockerurl)
}

var addDescription = `
Add Docker profile

`
