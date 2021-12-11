package truenas

func String(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}

type AvailableCertificates map[string]string

type Certificate struct {
	Id                 int               `json:"id"`
	Type               int               `json:"type"`
	Name               string            `json:"name"`
	Certificate        string            `json:"certificate"`
	PrivateKey         string            `json:"privatekey"`
	RootPath           string            `json:"root_path"`
	CertificatePath    string            `json:"certificate_path"`
	PrivateKeyPath     string            `json:"privatekey_path"`
	CertificateType    string            `json:"cert_type"`
	Revoked            bool              `json:"revoked"`
	Issuer             string            `json:"issuer"`
	ChainList          []string          `json:"chain_list"`
	Country            string            `json:"country"`
	State              string            `json:"state"`
	City               string            `json:"city"`
	Organization       string            `json:"organization"`
	OrganizationalUnit string            `json:"organizational_unit"`
	Common             string            `json:"common"`
	SAN                []string          `json:"san"`
	Email              string            `json:"email"`
	DN                 string            `json:"DN"`
	Extensions         map[string]string `json:"extensions"`
	From               string            `json:"from"`
	Until              string            `json:"until"`
	Chain              bool              `json:"chain"`
	Fingerprint        string            `json:"fingerprint"`
}

type SystemGeneralPutConfig struct {
	UICertificate   *int  `json:"ui_certificate,omitempty"`
	UIHttpsRedirect *bool `json:"ui_httpsredirect,omitempty"`
}
