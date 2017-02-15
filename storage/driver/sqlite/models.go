package sqlite

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"created_at"`
}

type Cert struct {
	Model

	Name     string `gorm:"unique;"`
	Type     string
	CN       string
	AltNames string

	TSAURL string
}

type Profile struct {
	Model

	Name       string `gorm:"unique;"`
	Cert       Cert
	CertID     uint
	DockerHost string
}
