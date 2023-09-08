package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/compscore/compscore/pkg/checks"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/sirupsen/logrus"
)

var (
	Status proto.StatusEnum = proto.StatusEnum_PAUSED

	runLock       *structs.Lock = &structs.Lock{}
	databaseMutex *sync.Mutex   = &sync.Mutex{}
	quit          chan struct{} = make(chan struct{})
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

	roundMutex := &sync.Mutex{}

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

			roundMutex.Lock()

			err := runRound(roundMutex)
			if err != nil {
				logrus.WithError(err).Error("Failed to run round")
			}

		case <-quit:
			// wait until the round is finished
			roundMutex.Lock()
			roundMutex.Unlock()

			return
		}
	}
}

func runRound(roundMutex *sync.Mutex) error {
	defer roundMutex.Unlock()

	round, err := data.Round.CreateNextRound()
	if err != nil {
		return err
	}

	checks := 0

	for _, team := range config.RunningConfig.Teams {
		for range team.Checks {
			checks++
		}
	}

	wgRound := &sync.WaitGroup{}
	wgRound.Add(checks)

	for _, team := range config.RunningConfig.Teams {
		for _, check := range team.Checks {
			go runScoreCheck(round.Number, check, team.Number, wgRound)
		}
	}

	wgRound.Wait()

	return nil
}

type checkResult struct {
	Success bool
	Message string
}

func runScoreCheck(round int, check structs.Check_s, team int8, wg *sync.WaitGroup) {
	defer wg.Done()

	entStatus, err := data.Status.Get(round, check.Name, int8(team))
	if err != nil {
		logrus.WithError(err).Errorf("Failed to get status for check: %v", check)
		return
	}

	runFunc, err := checks.GetCheckFunction(check.Release.Org, check.Release.Repo, check.Release.Tag)
	if err != nil {
		entStatus.Update().
			SetStatus(status.StatusDown).
			SetTime(time.Now()).
			SetError(err.Error()).
			Save(data.Ctx)

		logrus.WithError(err).Errorf("Failed to get check function: %v", check)
		return
	}

	checkCtx, checkCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.RunningConfig.Scoring.Interval))
	defer checkCancel()

	outerCtx, outerCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(float64(config.RunningConfig.Scoring.Interval)*1.1))
	defer outerCancel()

	returnChan := make(chan checkResult, 1)
	go func() {
		success, message := runFunc(checkCtx, check.Target, check.Command, check.ExpectedOutput, check.Credentials.Username, check.Credentials.Password)
		returnChan <- checkResult{
			Success: success,
			Message: message,
		}

		close(returnChan)
	}()

	select {
	case <-outerCtx.Done():
		databaseMutex.Lock()
		defer databaseMutex.Unlock()

		_, err := entStatus.Update().
			SetStatus(status.StatusDown).
			SetTime(time.Now()).
			SetError("timeout without return: " + outerCtx.Err().Error()).
			Save(data.Ctx)
		if err != nil {
			logrus.WithError(err).Errorf("failed to update status for check: %v", check)
		}
		return
	case checkOutput := <-returnChan:
		databaseMutex.Lock()
		defer databaseMutex.Unlock()

		_, err := entStatus.Update().
			SetStatus(func() status.Status {
				if checkOutput.Success {
					return status.StatusUp
				}
				return status.StatusDown
			}()).
			SetTime(time.Now()).
			SetError(checkOutput.Message).
			Save(data.Ctx)
		if err != nil {
			logrus.WithError(err).Errorf("failed to update status for check: %v", check)
		}
		return

	}
}
