package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compscore",
	Short: "Compscore is a scoring engine for Red/Blue Competitions",
	Long:  "Compscore is a scoring engine for Red/Blue Competitions",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logrus.WithError(err).Fatal("failed to execute root command")
	}
}
