package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientKillCmd = &cobra.Command{
	Use: "kill",
	Run: clientKillRun,
}

func clientKillRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	message, err := client.Kill(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to kill server")
	}

	logrus.Info(message)
}

func init() {
	clientCmd.AddCommand(clientKillCmd)
}
