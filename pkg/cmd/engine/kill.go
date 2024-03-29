package engine

import (
	"context"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill the engine",
	Long:  "Kill the engine",
	Run:   killRun,
}

// killRun kills the engine
func killRun(cmd *cobra.Command, args []string) {
	config.Init()

	client.Open()
	defer client.Close()

	message, err := client.Kill(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to kill server")
	}

	logrus.Info(message)
}
