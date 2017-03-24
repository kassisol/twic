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

type ClientConfig struct {
	TwicDir  string
	CertsDir string
	Name     string
}

type EngineConfig struct {
	TwicDir  string
	CertsDir string
}

type Filer interface {
	Init() error
	SetName(name string)
	DBFileName() string
	CertFilesName() (CertFiles, error)
}

func New(dtype string) Filer {
	if dtype == "client" {
		return &ClientConfig{}
	}

	return &EngineConfig{}
}

func (c *ClientConfig) Init() error {
	u := user.New()

	twicDir := path.Join(u.HomeDir, ".twic")
	certsDir := path.Join(twicDir, "certs")

	if err := filedir.CreateDirIfNotExist(twicDir, false, 0750); err != nil {
		return err
	}

	if err := filedir.CreateDirIfNotExist(certsDir, false, 0750); err != nil {
		return err
	}

	c.TwicDir = twicDir
	c.CertsDir = certsDir

	return nil
}

func (c *ClientConfig) SetName(name string) {
	c.Name = name
}

func (c *ClientConfig) DBFileName() string {
	return path.Join(c.TwicDir, "data.db")
}

func (c *ClientConfig) CertFilesName() (CertFiles, error) {
	certNameDir := path.Join(c.CertsDir, c.Name)

	if err := filedir.CreateDirIfNotExist(certNameDir, false, 0750); err != nil {
		return CertFiles{}, err
	}

	return CertFiles{
		Dir: certNameDir,
		Ca:  path.Join(certNameDir, "ca.pem"),
		Key: path.Join(certNameDir, "key.pem"),
		Crt: path.Join(certNameDir, "cert.pem"),
	}, nil
}

func (c *EngineConfig) Init() error {
	//	twicDir := "/var/lib/twic"
	certsDir := "/etc/docker/tls"

	/*	err := filedir.CreateDirIfNotExist(twicDir, 0750)
		if err != nil {
			return err
		}*/

	if err := filedir.CreateDirIfNotExist(certsDir, false, 0750); err != nil {
		return err
	}

	//	c.TwicDir = twicDir
	c.CertsDir = certsDir

	return nil
}

func (c *EngineConfig) SetName(name string) {
}

func (c *EngineConfig) DBFileName() string {
	return ""
}

func (c *EngineConfig) CertFilesName() (CertFiles, error) {
	return CertFiles{
		Dir: c.CertsDir,
		Ca:  path.Join(c.CertsDir, "ca.pem"),
		Key: path.Join(c.CertsDir, "server-key.pem"),
		Crt: path.Join(c.CertsDir, "server-cert.pem"),
	}, nil
}
