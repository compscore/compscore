package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientStartCmd = &cobra.Command{
	Use: "start",
	Run: clientStartRun,
}

func clientStartRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	message, err := client.Start(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to Start server")
	}

	logrus.Info(message)
}

func init() {
	clientCmd.AddCommand(clientStartCmd)
}
