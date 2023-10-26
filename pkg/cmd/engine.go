package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineCmd = &cobra.Command{
	Use:     "engine",
	Aliases: []string{"e", "eng"},
	Run:     engineRun,
}

func engineRun(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}

func init() {
	rootCmd.AddCommand(engineCmd)
}
