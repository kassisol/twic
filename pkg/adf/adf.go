package adf

import (
	"path"

	"github.com/juliengk/go-utils/filedir"
	"github.com/juliengk/go-utils/user"
)

type CertFiles struct {
	Dir string
	Ca  string
	Key string
	Crt string
}

type Config struct {
	TwicDir  string
	CertsDir string
}

func New() (*Config, error) {
	user, err := user.New()
	if err != nil {
		return &Config{}, err
	}

	twicDir := path.Join(user.HomeDir, ".twic")
	certsDir := path.Join(twicDir, "certs")

	// Create directories
	err = filedir.CreateDirIfNotExist(twicDir, 0750)
	if err != nil {
		return &Config{}, err
	}

	err = filedir.CreateDirIfNotExist(certsDir, 0750)
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		TwicDir:  twicDir,
		CertsDir: certsDir,
	}, nil
}

func (c *Config) DBFileName() string {
	return path.Join(c.TwicDir, "data.db")
}

func (c *Config) CertFilesName(name string) (CertFiles, error) {
	certNameDir := path.Join(c.CertsDir, name)

	err := filedir.CreateDirIfNotExist(certNameDir, 0750)
	if err != nil {
		return CertFiles{}, err
	}

	return CertFiles{
		Dir: certNameDir,
		Ca:  path.Join(certNameDir, "ca.pem"),
		Key: path.Join(certNameDir, "key.pem"),
		Crt: path.Join(certNameDir, "cert.pem"),
	}, nil
}
