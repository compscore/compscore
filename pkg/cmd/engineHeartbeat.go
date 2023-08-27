package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineHeartbeatCmd = &cobra.Command{
	Use:     "heartbeat",
	Short:   "Send a heartbeat to the engine",
	Long:    "Send a heartbeat to the engine",
	Aliases: []string{"ping"},
	Run:     enginePingRun,
}

func enginePingRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	err := client.Heartbeat(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Heartbeat failed")
	}

	logrus.Info("Heartbeat successful")
}

func init() {
	engineCmd.AddCommand(engineHeartbeatCmd)
}
