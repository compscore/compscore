package engine

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "engine",
	Aliases: []string{"e", "eng"},
	Run:     cmdRun,
}

func cmdRun(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		logrus.WithError(err).Fatal("failed to print help")
	}
}

func init() {
	Cmd.AddCommand(
		heartbeatCmd,
		killCmd,
		pauseCmd,
		startCmd,
		statusCmd,
	)
}
