package sqlite

import (
	"fmt"

	"github.com/kassisol/twic/storage/driver"
)

func (c *Config) AddCert(name, ptype, cn, altNames, tsaUrl string) {
	c.DB.Create(&Cert{
		Name:     name,
		Type:     ptype,
		CN:       cn,
		AltNames: altNames,
		TSAURL:   tsaUrl,
	})
}

func (c *Config) RemoveCert(name string) error {
	if c.certUsedInProfile(name) {
		return fmt.Errorf("cert \"%s\" cannot be removed. It is being used by a profile", name)
	}

	c.DB.Where("name = ?", name).Delete(Cert{})

	return nil
}

func (c *Config) GetCert(name string) driver.CertResult {
	var cert Cert

	c.DB.Where("name = ?", name).First(&cert)

	return driver.CertResult{
		Name:     cert.Name,
		Type:     cert.Type,
		CN:       cert.CN,
		AltNames: cert.AltNames,
		TSAURL:   cert.TSAURL,
	}
}

func (c *Config) ListCerts() []driver.CertResult {
	var certs []Cert
	var result []driver.CertResult

	c.DB.Find(&certs)

	for _, cert := range certs {
		r := driver.CertResult{
			Name:     cert.Name,
			Type:     cert.Type,
			CN:       cert.CN,
			AltNames: cert.AltNames,
			TSAURL:   cert.TSAURL,
		}

		result = append(result, r)
	}

	return result
}

func (c *Config) certUsedInProfile(name string) bool {
	var count int64

	c.DB.Table("profiles").Joins("JOIN certs ON certs.id = profiles.cert_id").Where("certs.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
