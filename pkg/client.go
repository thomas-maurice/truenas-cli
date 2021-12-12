package truenas

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	yaml "gopkg.in/yaml.v3"
)

type Client struct {
	client   *http.Client
	apiKey   string
	endpoint string
}

type Config struct {
	Endpoint           string `yaml:"endpoint"`
	ApiKey             string `yaml:"apiKey"`
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
}

func LoadConfigFromFile(fileName string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(configBytes, &cfg)
	return &cfg, err
}

func NewClient(config *Config) (*Client, error) {
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: config.InsecureSkipVerify,
				},
			},
		},
		apiKey:   config.ApiKey,
		endpoint: config.Endpoint,
	}, nil
}

func (c *Client) Get(apiPath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2.0/%s", c.endpoint, apiPath), nil)
	if c.apiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) Post(apiPath string, body interface{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2.0/%s", c.endpoint, apiPath), bytes.NewReader(bodyBytes))
	req.Header.Add("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) Put(apiPath string, body interface{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2.0/%s", c.endpoint, apiPath), bytes.NewReader(bodyBytes))
	req.Header.Add("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) Delete(apiPath string) (*http.Response, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2.0/%s", c.endpoint, apiPath), nil)
	if c.apiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}
