package profile

import (
	"os"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/kassisol/twic/pkg/adf"
	"github.com/kassisol/twic/storage"
	"github.com/spf13/cobra"
)

var (
	profileEnvShell string
	profileEnvUnset bool
)

type Data struct {
	Shell     string
	Unset     bool
	TLSVerify string
	Host      string
	CertPath  string
}

func newEnvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env [name]",
		Short: "Set / Unset Docker environment variables",
		Long:  envDescription,
		Run:   runEnv,
	}

	flags := cmd.Flags()
	flags.StringVarP(&profileEnvShell, "shell", "s", "bash", "Force environment to be configured for a specified shell: (tcsh, bash)")
	flags.BoolVarP(&profileEnvUnset, "unset", "u", false, "Unset variables instead of setting them")

	return cmd
}

func runEnv(cmd *cobra.Command, args []string) {
	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	// Data validations
	if profileEnvShell != "bash" && profileEnvShell != "tcsh" {
		log.Fatal("Shell is not correct")
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

	profile := s.GetProfile(args[0])

	config.SetName(profile.Cert.Name)

	cf, err := config.CertFilesName()
	if err != nil {
		log.Fatal(err)
	}

	data := Data{
		Shell:     profileEnvShell,
		Unset:     profileEnvUnset,
		TLSVerify: "1",
		Host:      profile.DockerHost,
		CertPath:  cf.Dir,
	}

	t := template.New("Shell commands template")
	t = template.Must(t.Parse(envTpl))

	t.Execute(os.Stdout, data)
}

var envTpl = `
{{- if eq .Shell "bash" }}
{{- if .Unset }}
unset DOCKER_HOST DOCKER_TLS_VERIFY DOCKER_CERT_PATH
{{- else }}
export DOCKER_HOST={{ .Host }}
export DOCKER_TLS_VERIFY={{ .TLSVerify }}
export DOCKER_CERT_PATH={{ .CertPath }}/
{{- end }}
{{- end }}
{{- if eq .Shell "tcsh" }}
{{- if .Unset }}
unsetenv DOCKER_HOST DOCKER_TLS_VERIFY DOCKER_CERT_PATH
{{- else }}
setenv DOCKER_HOST {{ .Host }}
setenv DOCKER_TLS_VERIFY {{ .TLSVerify }}
setenv DOCKER_CERT_PATH {{ .CertPath }}/
{{- end }}
{{- end }}
`

var envDescription = `
Set / Unset Docker environment variables

`
