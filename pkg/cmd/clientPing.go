package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientPingCmd = &cobra.Command{
	Use: "ping",
	Run: clientPingRun,
}

func clientPingRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	message, err := client.Ping.Ping(context.Background(), "Hello World!")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to ping server")
	}

	logrus.Info(message)
}

func init() {
	clientCmd.AddCommand(clientPingCmd)
}
