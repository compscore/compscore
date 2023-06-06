package cmd

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	socketPath string = "compscore.sock"
)

// Statuses of the scoring engine
var (
	STOPPED  string = "stopped"
	RUNNING  string = "running"
	STOPPING string = "stopping"
	KILLED   string = "killed"
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
	// Handle SIGTERM
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)

	// Open the file for writing
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ticker := time.NewTicker(time.Second)

	s := "x"
	for {
		select {
		case <-signalChan:
			return
		case <-ticker.C:
			_, err := file.WriteString(s)
			if err != nil {
				logrus.WithError(err).Error("Failed to write to file")
			}
		}
	}
}

func control() {
	var (
		stdout    bytes.Buffer
		stderr    bytes.Buffer
		stopChan  = make(chan struct{})
		startChan = make(chan struct{})
		killChan  = make(chan struct{})
		status    string
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
				conn.Write([]byte(status))
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
			err := syscall.Kill(cmd.Process.Pid, syscall.SIGTERM)
			if err != nil {
				logrus.WithError(err).Error("Failed to kill worker")
			}

			status = STOPPING

			err = cmd.Wait()
			if err != nil {
				logrus.WithError(err).Error("Failed to wait for worker")
			}

			status = STOPPED
		}
	}()

	go func() {
		startChan <- struct{}{}
	}()

	for {
		select {
		case <-startChan:
			status = RUNNING

			cmd = exec.Command(os.Args[0], "engine", "start", "--worker")
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Start()
			if err != nil {
				logrus.WithError(err).Error("Failed to start worker")
				return
			}
		case <-killChan:
			status = KILLED
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
