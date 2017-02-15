package driver

type CertResult struct {
	Name     string
	Type     string
	CN       string
	AltNames string
	TSAURL   string
}

type ProfileResult struct {
	Name       string
	Cert       CertResult
	DockerHost string
}
