package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientHeartbeatCmd = &cobra.Command{
	Use:     "heartbeat",
	Aliases: []string{"ping"},
	Run:     clientPingRun,
}

func clientPingRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	err := client.Heartbeat(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Heartbeat failed")
	}

	logrus.Info("Heartbeat successful")
}

func init() {
	clientCmd.AddCommand(clientHeartbeatCmd)
}
