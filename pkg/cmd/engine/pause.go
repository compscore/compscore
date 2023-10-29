package engine

import (
	"context"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:     "pause",
	Short:   "Pause the engine",
	Long:    "Pause the engine",
	Aliases: []string{"p", "stop"},
	Run:     pauseRun,
}

// pauseRun pauses the engine
func pauseRun(cmd *cobra.Command, args []string) {
	config.Init()

	client.Open()
	defer client.Close()

	message, err := client.Pause(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to pause server")
	}

	logrus.Info(message)
}
