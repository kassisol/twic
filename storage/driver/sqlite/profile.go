package sqlite

import (
	"github.com/kassisol/twic/storage/driver"
)

func (c *Config) AddProfile(name, certName, dockerHost string) {
	var cert Cert
	c.DB.Where("name = ?", certName).First(&cert)

	c.DB.Create(&Profile{
		Name:       name,
		Cert:       cert,
		DockerHost: dockerHost,
	})
}

func (c *Config) RemoveProfile(name string) {
	c.DB.Where("name = ?", name).Delete(Profile{})
}

func (c *Config) GetProfile(name string) driver.ProfileResult {
	var profile Profile

	c.DB.Preload("Cert").Where("name = ?", name).First(&profile)

	cert := profile.Cert
	cr := driver.CertResult{
		Name:     cert.Name,
		Type:     cert.Type,
		CN:       cert.CN,
		AltNames: cert.AltNames,
		TSAURL:   cert.TSAURL,
	}

	return driver.ProfileResult{
		Name:       profile.Name,
		Cert:       cr,
		DockerHost: profile.DockerHost,
	}
}

func (c *Config) ListProfiles() []driver.ProfileResult {
	var profiles []Profile
	var result []driver.ProfileResult

	c.DB.Preload("Cert").Find(&profiles)

	for _, profile := range profiles {
		cert := profile.Cert

		cr := driver.CertResult{
			Name:     cert.Name,
			Type:     cert.Type,
			CN:       cert.CN,
			AltNames: cert.AltNames,
			TSAURL:   cert.TSAURL,
		}

		r := driver.ProfileResult{
			Name:       profile.Name,
			Cert:       cr,
			DockerHost: profile.DockerHost,
		}

		result = append(result, r)
	}

	return result
}
