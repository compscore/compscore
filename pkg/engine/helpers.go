package engine

import (
	"context"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
)

// UnixSocketExists checks if the unix socket exists
func UnixSocketExists() (bool, error) {
	stat, err := os.Stat(config.RunningConfig.Engine.Socket)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if stat.Mode()&os.ModeSocket != 0 {
		return true, nil
	}

	return false, nil
}

// UnixSocketActive checks if the unix socket is active
func UnixSocketActive() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client.Open()
	defer client.Close()

	err := client.Heartbeat(ctx)
	if err != nil && err != context.DeadlineExceeded {
		return false, err
	} else if err == context.DeadlineExceeded {
		return false, nil
	}

	return true, nil
}

// Spawn Compscore Engine Process
func SpawnCompscoreEngine() error {
	logrus.Info("Spawning compscore engine process")

	engine := exec.Command(os.Args[0], "server")
	engine.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	err := engine.Start()
	if err != nil {
		return err
	}

	logrus.Info("Compscore engine process spawned")
	return nil
}
