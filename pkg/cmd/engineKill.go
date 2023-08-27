package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineKillCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill the engine",
	Long:  "Kill the engine",
	Run:   engineKillRun,
}

func engineKillRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	message, err := client.Kill(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to kill server")
	}

	logrus.Info(message)
}

func init() {
	engineCmd.AddCommand(engineKillCmd)
}
