package engine

import (
	"context"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// heartbeatCmd represents the heartbeat command
var heartbeatCmd = &cobra.Command{
	Use:     "heartbeat",
	Short:   "Send a heartbeat to the engine",
	Long:    "Send a heartbeat to the engine",
	Aliases: []string{"ping", "p", "h"},
	Run:     heartbeatRun,
}

// heartbeatRun sends a heartbeat to the engine
func heartbeatRun(cmd *cobra.Command, args []string) {
	config.Init()

	client.Open()
	defer client.Close()

	err := client.Heartbeat(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Heartbeat failed")
	}

	logrus.Info("Heartbeat successful")
}
