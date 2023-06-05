package cmd

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineKillCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kills the scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineKill,
}

func init() {
	engineCmd.AddCommand(engineKillCmd)
}

func engineKill(cmd *cobra.Command, args []string) {
	logrus.Info("Stopping engine")

	conn, err := net.Dial("unix", "compscore.sock")
	if err != nil {
		logrus.Info("Already Killed")
		return
	}

	_, err = conn.Write([]byte("kill"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed to write to unix socket")
		return
	}

	logrus.Info("Kill")

	err = conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Failed to close connection")
	}
}
