package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	truenas "github.com/thomas-maurice/truenas-cli/pkg"
)

var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "System configuration",
	Long:  ``,
}

var SystemGeneralCmd = &cobra.Command{
	Use:   "general",
	Short: "General system configuration",
	Long:  ``,
}

var SystemGeneralGetUICerts = &cobra.Command{
	Use:   "get-ui-certs",
	Short: "Gets the available certificates for the UI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		result, err := client.GetUICertificateChoices()
		if err != nil {
			logrus.Fatalf("could not get ui cert choices: %s", err)
		}

		output(result)
	},
}

var SystemGeneralReboot = &cobra.Command{
	Use:   "reboot",
	Short: "Reboots the NAS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		resp, err := client.Post("system/reboot", nil)
		if err != nil {
			logrus.Fatalf("could not reboot: %s", err)
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Fatal("could not read http response")
		}
		defer resp.Body.Close()

		fmt.Println(string(resp.Status), string(respBytes))
	},
}

var SystemGeneralShutdown = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdowns the NAS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		resp, err := client.Post("system/shutdown", nil)
		if err != nil {
			logrus.Fatalf("could not shutdown: %s", err)
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Fatal("could not read http response")
		}
		defer resp.Body.Close()

		fmt.Println(string(resp.Status), string(respBytes))
	},
}

var SystemGeneralSetUICert = &cobra.Command{
	Use:   "set-ui-cert [cert name]",
	Short: "Updates the general systems config",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		i, err := strconv.Atoi(args[0])
		if err != nil {
			logrus.Fatal("this isnt an integer")
		}

		cfg := truenas.SystemGeneralPutConfig{
			UICertificate: truenas.Int(i),
		}

		resp, err := client.Put("/system/general", &cfg)
		if err != nil {
			logrus.Fatalf("could not set general config: %s", err)
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Fatal("could not read http response")
		}
		defer resp.Body.Close()

		fmt.Println(string(respBytes))
	},
}

func InitSystemCmd() {
	SystemGeneralCmd.AddCommand(SystemGeneralGetUICerts)
	SystemGeneralCmd.AddCommand(SystemGeneralSetUICert)
	SystemGeneralCmd.AddCommand(SystemGeneralReboot)
	SystemGeneralCmd.AddCommand(SystemGeneralShutdown)

	SystemCmd.AddCommand(SystemGeneralCmd)
}
