package engine

import (
	"context"
	"os"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/engine"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run", "r"},
	Short:   "Start the compscore server",
	Long:    `Start the compscore server`,
	Run:     startRun,
}

// startRun starts the compscore server
func startRun(cmd *cobra.Command, args []string) {
	config.Init()

	exists, err := engine.UnixSocketExists()
	if err != nil {
		logrus.WithError(err).Fatal("Error checking if unix socket exists")
	}

	if !exists {
		// Socket Does Not Exist
		// Spawn Engine

		err := engine.SpawnCompscoreEngine()
		if err != nil {
			logrus.WithError(err).Fatal("Error spawning compscore engine")
		}

		logrus.Info("Compscore engine spawned")
	} else {
		active, err := engine.UnixSocketActive()
		if err != nil {
			logrus.WithError(err).Fatal("Error checking if unix socket is active")
		}

		if active {
			// Socket Exist and is Active
			// Send Start Command to Engine

			client.Open()
			defer client.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			message, err := client.Start(ctx)
			cancel()

			if err != nil {
				logrus.WithError(err).Fatal("Error starting compscore server")
			}

			logrus.WithField("message", message).Info("Compscore server started")
		} else {
			// Socket Exist and is Inactive
			// Remove Socket and Spawn Engine

			err := os.Remove(config.Engine.Socket)
			if err != nil {
				logrus.WithError(err).Fatal("Error removing unix socket")
			}

			err = engine.SpawnCompscoreEngine()
			if err != nil {
				logrus.WithError(err).Fatal("Error spawning compscore engine")
			}

			logrus.Info("Compscore engine spawned")
		}
	}
}
