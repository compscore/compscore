package cmd

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStop,
}

func init() {
	engineCmd.AddCommand(engineStopCmd)
}

func engineStop(cmd *cobra.Command, args []string) {
	logrus.Info("Stopping engine")

	conn, err := net.Dial("unix", "compscore.sock")
	if err != nil {
		logrus.Info("Already stopped")
		return
	}

	_, err = conn.Write([]byte("stop"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed to write to unix socket")
		return
	}

	logrus.Info("Stopped")

	err = conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Failed to close connection")
	}
}
