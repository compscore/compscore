package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "compscore",
	Short: "Cyber Competition Scoring Platform",
	Long:  "Compscore is a full-stack Cyber Competition Platform. Compscore at it's most basic form is able to preform service checks on designated teams while maintaining uptime metrics to be represented as scores.",
	Args:  cobra.ArbitraryArgs,
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		logrus.WithError(err).Fatal("failed to load rootCmd")
	}
}
