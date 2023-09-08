package engine

import (
	"fmt"
	"os"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/sirupsen/logrus"
)

var (
	Status proto.StatusEnum = proto.StatusEnum_PAUSED

	runLock *structs.Lock = &structs.Lock{}
	quit    chan struct{} = make(chan struct{})
)

func Pause() error {
	if !runLock.IsLocked() {
		return fmt.Errorf("engine is not running")
	}

	Status = proto.StatusEnum_PAUSED
	quit <- struct{}{}

	return nil
}

func Run() error {
	err := runLock.Lock()
	if err != nil {
		return fmt.Errorf("failed to lock run lock: %w", err)
	}

	Status = proto.StatusEnum_RUNNING
	go runEngine()

	return nil
}

func runEngine() {
	interval := config.RunningConfig.Scoring.Interval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	defer runLock.Unlock()
	defer func() {
		Status = proto.StatusEnum_PAUSED
	}()

	for {
		select {
		case <-ticker.C:
			logrus.Info("Scoring interval")
			if interval != config.RunningConfig.Scoring.Interval {
				interval = config.RunningConfig.Scoring.Interval
				ticker = time.NewTicker(time.Duration(config.RunningConfig.Scoring.Interval) * time.Second)
			}

			if Status != proto.StatusEnum_RUNNING {
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
			return
		}
	}
}
