package cmd

import (
	"github.com/sirupsen/logrus"
	truenas "github.com/thomas-maurice/truenas-cli/pkg"
)

func mustGetClient(cfgFile string) *truenas.Client {
	cfg, err := truenas.LoadConfigFromFile("config.yaml")
	if err != nil {
		logrus.Fatalf("could not load config: %s", err)
	}

	client, err := truenas.NewClient(cfg)
	if err != nil {
		logrus.Fatalf("could not create client: %s", err)
	}

	return client
}
