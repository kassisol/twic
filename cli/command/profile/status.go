package profile

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Display Docker environment variables if set",
		Long:  statusDescription,
		Run:   runStatus,
	}

	return cmd
}

func runStatus(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Usage()
		os.Exit(-1)
	}

	dockerHost := os.Getenv("DOCKER_HOST")
	dockerTLSVerify := os.Getenv("DOCKER_TLS_VERIFY")
	dockerCertPath := os.Getenv("DOCKER_CERT_PATH")

	if len(dockerHost) > 0 && len(dockerTLSVerify) > 0 && len(dockerCertPath) > 0 {
		fmt.Printf("DOCKER_HOST=%s\n", dockerHost)
		fmt.Printf("DOCKER_TLS_VERIFY=%s\n", dockerTLSVerify)
		fmt.Printf("DOCKER_CERT_PATH=%s\n", dockerCertPath)
	} else {
		fmt.Println("Docker variables are not set")
	}
}

var statusDescription = `
Display Docker environment variables if set

`
