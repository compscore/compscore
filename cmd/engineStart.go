package cmd

import (
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var engineStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStart,
}

func init() {
	engineStartCmd.Flags().BoolP("dettached", "d", false, "Run command in dettached mode")
	engineCmd.AddCommand(engineStartCmd)
}

func engineStart(cmd *cobra.Command, args []string) {
	if cmd.Flag("dettached").Value.String() == "false" {
		logrus.Info("Spawning detached process...")

		fork := exec.Command(os.Args[0], "engine", "start", "--dettached")
		fork.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}

		fork.Start()
		logrus.Info("Spawned detached process")
	} else {

		logrus.Info("Starting engine")

		// Remove socket file if it exists
		err := os.RemoveAll("compscore.sock")
		if err != nil {
			logrus.WithError(err).Error("Failed to remove socket file")
		}

		// Listen on unix socket
		listen, err := net.Listen("unix", "compscore.sock")
		if err != nil {
			logrus.WithError(err).Fatal("Failed to listen on unix socket")
		}

		// Remove socket file on exit
		defer func() {
			err := listen.Close()
			if err != nil {
				logrus.WithError(err).Error("Failed to close socket")
			}

			err = os.RemoveAll("compscore.sock")
			if err != nil {
				logrus.WithError(err).Error("Failed to remove socket file")
			}
		}()

		stopChan := make(chan struct{})

		// Accept connections
		go func() {
			for {
				fd, err := listen.Accept()
				if err != nil {
					logrus.WithError(err).Error("Failed to accept connection")
					continue
				}

				logrus.Info("Accepted connection")
				go handleConnection(fd, stopChan)
			}
		}()

		// Wait for stop signal
		<-stopChan
	}
}

func handleConnection(fd net.Conn, stopChan chan struct{}) {
	buf := make([]byte, 1024)

	defer func() {
		fd.Close()
		logrus.Info("Closed connection")
	}()

	// Read from connection
	response_len, err := fd.Read(buf)
	if err != nil {
		logrus.WithError(err).Error("Failed to read from connection")
		return
	}

	// Handle command
	switch strings.TrimSpace(string(buf[:response_len])) {
	case "status":
		_, err := fd.Write([]byte("Running"))
		if err != nil {
			logrus.WithError(err).Error("Failed to write to connection")
		}
	case "stop":
		stopChan <- struct{}{}
		return
	default:
		logrus.WithField("command", string(buf)).Error("Unknown command")
	}
}
