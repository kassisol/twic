package driver

type Storager interface {
	AddCert(name, ptype, cn, altNames, caUrl string)
	RemoveCert(name string) error
	GetCert(name string) CertResult
	ListCerts() []CertResult

	AddProfile(name, certName, dockerHost string)
	RemoveProfile(name string)
	GetProfile(name string) ProfileResult
	ListProfiles() []ProfileResult

	End()
}
