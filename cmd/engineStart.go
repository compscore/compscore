package cmd

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	socketPath string = "compscore.sock"
)

var engineStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStart,
}

func init() {
	engineStartCmd.Flags().BoolP("control", "", false, "Runs control process that is detached from terminal")
	engineStartCmd.Flags().BoolP("worker", "", false, "Runs worker process that is control by control process")
	engineCmd.AddCommand(engineStartCmd)
}

func engineStart(cmd *cobra.Command, args []string) {
	controlFlag, _ := cmd.Flags().GetBool("control")
	workerFlag, _ := cmd.Flags().GetBool("worker")

	if controlFlag {
		control()
	} else if workerFlag {
		worker()
	} else {
		start()
	}
}

func socketAlive() bool {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func worker() {
	// Open the file for writing
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := "x"
	for {
		// Write the string to the file
		if _, err := file.WriteString(s); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func control() {
	var (
		stdout    bytes.Buffer
		stderr    bytes.Buffer
		stopChan  = make(chan struct{})
		startChan = make(chan struct{})
		killChan  = make(chan struct{})
		paused    = false
		cmd       *exec.Cmd
	)

	go func() {
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

		// Accept connections
		for {
			conn, err := listen.Accept()
			if err != nil {
				logrus.WithError(err).Error("Failed to accept connection")
				continue
			}

			// Read from connection
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				logrus.WithError(err).Error("Failed to read from connection")
				continue
			}

			// Handle command
			switch strings.TrimSpace(string(buf[:n])) {
			case "status":
				if paused {
					conn.Write([]byte("Paused"))
				} else {
					conn.Write([]byte("Running"))
				}
			case "start":
				startChan <- struct{}{}
				conn.Write([]byte("Started"))
			case "stop":
				stopChan <- struct{}{}
				conn.Write([]byte("Stopped"))
			case "kill":
				killChan <- struct{}{}
				conn.Write([]byte("Killed"))
			default:
				conn.Write([]byte("Invalid command"))
			}
		}
	}()

	go func() {
		for range stopChan {
			if cmd != nil {
				if cmd.Process != nil {
					cmd.Process.Kill()
				}
			}
			paused = true
		}
	}()

	go func() {
		startChan <- struct{}{}
	}()

	for {
		select {
		case <-startChan:
			paused = false

			cmd = exec.Command(os.Args[0], "engine", "start", "--worker")
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Start()
			if err != nil {
				logrus.WithError(err).Error("Failed to start worker")
				return
			}
		case <-killChan:
			cmd.Process.Kill()
			return
		}
	}
}

func start() {
	if socketAlive() {
		logrus.Info("Engine already running")
		conn, err := net.Dial("unix", "compscore.sock")
		if err != nil {
			logrus.WithError(err).Fatal("Failed to connect to unix socket")
			return
		}

		_, err = conn.Write([]byte("start"))
		if err != nil {
			logrus.WithError(err).Fatal("Failed to write to unix socket")
			return
		}

		logrus.Info("Started")

		err = conn.Close()
		if err != nil {
			logrus.WithError(err).Error("Failed to close connection")
		}
	} else {
		logrus.Info("Spawning control process...")

		control := exec.Command(os.Args[0], "engine", "start", "--control")
		control.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}

		err := control.Start()
		if err != nil {
			logrus.WithError(err).Fatal("Failed to spawn control process")
			return
		}

		logrus.Info("Spawned control process")
		logrus.Info(control)
		os.Exit(0)
	}
}
