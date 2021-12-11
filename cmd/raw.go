package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RawCmd = &cobra.Command{
	Use:   "raw",
	Short: "Raw calls",
	Long:  ``,
}

var RawGetCmd = &cobra.Command{
	Use:   "get [path]",
	Short: "Performs a raw get",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := mustGetClient(configFile)

		result, err := client.Get(args[0])
		if err != nil {
			logrus.Fatalf("could not get %s: %s", args[0], err)
		}

		resultBytes, err := ioutil.ReadAll(result.Body)
		if err != nil {
			logrus.Fatalf("could not read result: %s", err)
		}
		defer result.Body.Close()

		fmt.Println(string(resultBytes))
	},
}

func InitRawCmd() {
	RawCmd.AddCommand(RawGetCmd)
}
