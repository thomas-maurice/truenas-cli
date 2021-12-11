package truenas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetUICertificateChoices() (*AvailableCertificates, error) {
	resp, err := c.Get("system/general/ui_certificate_choices")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API answered with %d", resp.StatusCode)
	}

	certsBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var certs AvailableCertificates
	err = json.Unmarshal(certsBytes, &certs)
	return &certs, err
}

func (c *Client) ListCertificates() ([]Certificate, error) {
	resp, err := c.Get("certificate")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API answered with %d", resp.StatusCode)
	}

	certsBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var certs []Certificate
	err = json.Unmarshal(certsBytes, &certs)
	return certs, err
}

func (c *Client) ImportCertificate(name string, cert string, key string) (*http.Response, error) {
	type importCertStruct struct {
		Name        string `json:"name"`
		Certificate string `json:"certificate"`
		PrivateKey  string `json:"privatekey"`
		CreateType  string `json:"create_type"`
	}

	certificate := &importCertStruct{
		Name:        name,
		Certificate: cert,
		PrivateKey:  key,
		CreateType:  "CERTIFICATE_CREATE_IMPORTED",
	}

	return c.Post("certificate", certificate)
}
