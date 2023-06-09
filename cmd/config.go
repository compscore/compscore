package cmd

import (
	"fmt"

	"github.com/compscore/compscore/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Subcommand for config management",
	Args:  cobra.ArbitraryArgs,
	Run:   func(cmd *cobra.Command, args []string) { printConfig() },
}

func printConfig() {
	configData, err := yaml.Marshal(config.Competition)
	if err != nil {
		logrus.WithError(err).Fatal("Error marshalling config")
	}

	fmt.Println(string(configData))
}

func init() {
	RootCmd.AddCommand(configCmd)
}
