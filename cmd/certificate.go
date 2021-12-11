package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CertificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Manages certificates on the TrueNAS system",
	Long:  ``,
}

var CertificateListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists certificates",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		result, err := client.ListCertificates()
		if err != nil {
			logrus.Fatalf("could not get certificates: %s", err)
		}

		output(result)
	},
}

var CertificateAddCmd = &cobra.Command{
	Use:   "add [name] [certificate file] [certificate key]",
	Short: "Adds a certificate",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		certBytes, err := ioutil.ReadFile(args[1])
		if err != nil {
			logrus.Fatalf("could not read certificate: %s", err)
		}

		certKey, err := ioutil.ReadFile(args[2])
		if err != nil {
			logrus.Fatalf("could not read key: %s", err)
		}

		resp, err := client.ImportCertificate(args[0], string(certBytes), string(certKey))
		if err != nil {
			logrus.Fatalf("could not upoad cert: %s", err)
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Fatal("could not read http response")
		}
		defer resp.Body.Close()

		fmt.Println(string(resp.Status), string(respBytes))
	},
}

func InitCertificateCmd() {
	CertificateCmd.AddCommand(CertificateListCmd)
	CertificateCmd.AddCommand(CertificateAddCmd)
}
