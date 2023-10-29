package engine

import (
	"context"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Get the status of the engine",
	Long:    "Get the status of the engine",
	Aliases: []string{"stat"},
	Run:     statusRun,
}

// statusRun gets the status of the engine
func statusRun(cmd *cobra.Command, args []string) {
	config.Init()

	client.Open()
	defer client.Close()

	status, message, err := client.Status(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to Status server")
	}

	logrus.WithField("status", status).Info(message)
}
