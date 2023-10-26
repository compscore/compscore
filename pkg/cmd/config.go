package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf", "c"},
	Short:   "Manage compscore configuration",
	Long:    "Manage compscore configuration",
	Run:     configRun,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func configRun(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}
