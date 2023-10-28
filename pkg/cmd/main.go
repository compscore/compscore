package cmd

import (
	"github.com/compscore/compscore/pkg/cmd/engine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "compscore",
	Short: "Compscore is a scoring engine for Red/Blue Competitions",
	Long:  "Compscore is a scoring engine for Red/Blue Competitions",
	Run:   run,
}

// print help if no subcommand is given
func run(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}

// Entrypoint for all commands
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logrus.WithError(err).Fatal("failed to execute root command")
	}
}

func init() {
	rootCmd.AddCommand(
		engine.Cmd,
		generateCmd,
		serverCmd,
		versionCmd,
	)
}
