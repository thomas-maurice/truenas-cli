package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thomas-maurice/truenas-cli/pkg/version"
)

var (
	configFile string
	marshaller string
)

var rootCmd = &cobra.Command{
	Use:   "truenas-cli",
	Short: "TrueNAS command line interface",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	InitRootCmd()
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Git Hash: %s\nBuild Host: %s\nBuild Time: %s\nBuild Tag: %s\n", version.BuildHash, version.BuildHost, version.BuildTime, version.Version)
	},
}

func InitRootCmd() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "Configuration file")
	rootCmd.PersistentFlags().StringVarP(&marshaller, "output", "o", "json", "Output format for server responses (client mode), must be yaml or json")

	InitSystemCmd()
	InitCertificateCmd()
	InitRawCmd()

	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(SystemCmd)
	rootCmd.AddCommand(RawCmd)
	rootCmd.AddCommand(CertificateCmd)
}
