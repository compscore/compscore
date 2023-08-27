package cmd

import (
	"context"
	"os"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/compscore/compscore/pkg/helpers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the compscore server",
	Long:  `Start the compscore server`,
	Run:   engineStartRun,
}

func engineStartRun(cmd *cobra.Command, args []string) {
	exists, err := helpers.UnixSocketExists()
	if err != nil {
		logrus.WithError(err).Fatal("Error checking if unix socket exists")
	}

	if !exists {
		// Socket Does Not Exist
		// Spawn Engine

		err := helpers.SpawnCompscoreEngine()
		if err != nil {
			logrus.WithError(err).Fatal("Error spawning compscore engine")
		}

		logrus.Info("Compscore engine spawned")
	} else {
		active, err := helpers.UnixSocketActive()
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

			err := os.Remove(config.UnixSocket)
			if err != nil {
				logrus.WithError(err).Fatal("Error removing unix socket")
			}

			err = helpers.SpawnCompscoreEngine()
			if err != nil {
				logrus.WithError(err).Fatal("Error spawning compscore engine")
			}

			logrus.Info("Compscore engine spawned")
		}
	}
}

func init() {
	engineCmd.AddCommand(engineStartCmd)
}
