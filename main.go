package main

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/thomas-maurice/truenas-cli/cmd"
	truenas "github.com/thomas-maurice/truenas-cli/pkg"
)

func main() {
	cmd.Execute()

	return

	cfg, err := truenas.LoadConfigFromFile("config.yaml")
	if err != nil {
		logrus.Fatalf("could not load config: %s", err)
	}

	client, err := truenas.NewClient(cfg)
	if err != nil {
		logrus.Fatalf("could not create client: %s", err)
	}

	res, err := client.Get("certificate")
	if err != nil {
		logrus.Fatalf("could not get certs: %s", err)
	}

	defer res.Body.Close()

	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBytes))
}
