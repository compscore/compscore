package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientPauseCmd = &cobra.Command{
	Use: "pause",
	Run: clientPauseRun,
}

func clientPauseRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	message, err := client.Pause(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to pause server")
	}

	logrus.Info(message)
}

func init() {
	clientCmd.AddCommand(clientPauseCmd)
}
