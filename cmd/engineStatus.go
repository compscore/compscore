package cmd

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check status of scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStatus,
}

func init() {
	engineCmd.AddCommand(engineStatusCmd)
}

func engineStatus(cmd *cobra.Command, args []string) {
	logrus.Info("Checking engine status....\n")

	conn, err := net.Dial("unix", "compscore.sock")
	if err != nil {
		logrus.Info("dead")
		return
	}

	_, err = conn.Write([]byte("status"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed to write to unix socket")
		return
	}

	buf := make([]byte, 1024)

	response_len, err := conn.Read(buf)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read from unix socket")
		return
	}

	logrus.Info(string(buf[:response_len]))

	err = conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Failed to close connection")
	}
}
