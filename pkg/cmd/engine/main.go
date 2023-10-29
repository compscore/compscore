package engine

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the engine command
var Cmd = &cobra.Command{
	Use:     "engine",
	Aliases: []string{"e", "eng"},
	Run:     cmdRun,
}

// cmdRun prints the help for the engine command if no subcommand is specified
func cmdRun(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}

// registers all engine commands
func init() {
	Cmd.AddCommand(
		heartbeatCmd,
		killCmd,
		pauseCmd,
		startCmd,
		statusCmd,
	)
}
