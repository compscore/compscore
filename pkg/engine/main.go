package engine

import (
	"os"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/compscore/compscore/pkg/grpc/server"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/sirupsen/logrus"
)

func Run() {
	interval := config.RunningConfig.Scoring.Interval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			logrus.Info("Scoring interval")
			if interval != config.RunningConfig.Scoring.Interval {
				interval = config.RunningConfig.Scoring.Interval
				ticker = time.NewTicker(time.Duration(config.RunningConfig.Scoring.Interval) * time.Second)
			}

			if server.Status != proto.StatusEnum_RUNNING {
				return
			}

			f, err := os.OpenFile("text.log",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logrus.WithError(err).Error("Failed to open log file")
				continue
			}
			defer f.Close()
			if _, err := f.WriteString("ballz\n"); err != nil {
				logrus.WithError(err).Error("Failed to write to log file")
				continue
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func GetAllGitRemotes() []structs.Release_s {
	remoteMap := make(map[structs.Release_s]bool)

	for _, team := range config.RunningConfig.Teams {
		for _, check := range team.Checks {
			remoteMap[check.Release] = true
		}
	}

	remoteSlice := make([]structs.Release_s, len(remoteMap))
	i := 0
	for remote := range remoteMap {
		remoteSlice[i] = remote
		i++
	}

	return remoteSlice
}
