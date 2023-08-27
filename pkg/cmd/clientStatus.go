package cmd

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientStatusCmd = &cobra.Command{
	Use: "status",
	Run: clientStatusRun,
}

func clientStatusRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	status, message, err := client.Status(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to Status server")
	}

	logrus.WithField("status", status).Info(message)
}

func init() {
	clientCmd.AddCommand(clientStatusCmd)
}
