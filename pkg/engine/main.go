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

	runLock    *structs.Lock = &structs.Lock{}
	quit       chan struct{} = make(chan struct{})
	exited     chan struct{} = make(chan struct{})
	roundMutex *sync.Mutex   = &sync.Mutex{}
)

type checkResult struct {
	Success bool
	Message string
	Team    int8
	Check   structs.Check_s
}

func Pause() error {
	if !runLock.IsLocked() {
		return fmt.Errorf("engine is not running")
	}

	Status = proto.StatusEnum_PAUSED
	quit <- struct{}{}
	<-exited
	roundMutex.Unlock()

	logrus.Info("Engine paused")

	return nil
}

func Stop() error {
	if !runLock.IsLocked() {
		return nil
	}

	Status = proto.StatusEnum_UNKNOWN
	quit <- struct{}{}
	<-exited
	roundMutex.Unlock()

	logrus.Info("Engine stopped")

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
			if interval != config.RunningConfig.Scoring.Interval {
				interval = config.RunningConfig.Scoring.Interval
				ticker = time.NewTicker(time.Duration(config.RunningConfig.Scoring.Interval) * time.Second)
			}

			if Status != proto.StatusEnum_RUNNING {
				return
			}

			entRound, err := data.Round.GetLastRound()
			if err != nil {
				logrus.WithError(err).Error("Failed to get last round")
				return
			}
			logrus.Infof("Running Round %d", entRound.Number)

			roundMutex.Lock()
			err = runRound(roundMutex)
			if err != nil {
				logrus.WithError(err).Error("Failed to run round")
			}

		case <-quit:
			// wait until the round is finished
			roundMutex.Lock()
			exited <- struct{}{}

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
		for _, check := range team.Checks {
			_, err := data.Status.Create(
				round.Number,
				check.Name,
				team.Number,
				status.StatusUnknown,
			)
			if err != nil {
				return err
			}

			checks++
		}
	}

	wgRound := &sync.WaitGroup{}
	wgRound.Add(checks)

	resultsChan := make(chan checkResult)
	defer close(resultsChan)

	for _, team := range config.RunningConfig.Teams {
		for _, check := range team.Checks {
			go runScoreCheck(round.Number, check, team.Number, resultsChan, wgRound)
		}
	}

	go func() {
		for result := range resultsChan {
			entStatus, err := data.Status.Get(round.Number, result.Check.Name, result.Team)
			if err != nil {
				logrus.WithError(err).Errorf("Failed to get status for check: %v", result.Check)
				continue
			}

			_, err = data.Status.Update(
				entStatus,
				func() status.Status {
					if result.Success {
						return status.StatusUp
					}
					return status.StatusDown
				}(),
				result.Message,
			)
			if err != nil {
				logrus.WithError(err).Errorf("Failed to update status for check: %v", result.Check)
				continue
			}
		}
	}()

	wgRound.Wait()

	return nil
}

func runScoreCheck(round int, check structs.Check_s, team int8, resultsChan chan checkResult, wg *sync.WaitGroup) {
	defer wg.Done()

	runFunc, err := checks.GetCheckFunction(check.Release.Org, check.Release.Repo, check.Release.Tag)
	if err != nil {
		resultsChan <- checkResult{
			Success: false,
			Message: err.Error(),
			Team:    team,
			Check:   check,
		}

		logrus.WithError(err).Errorf("Failed to get check function: %v", check)
		return
	}

	checkCtx, checkCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(float64(config.RunningConfig.Scoring.Interval)*0.9))
	defer checkCancel()

	outerCtx, outerCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(float64(config.RunningConfig.Scoring.Interval)*1))
	defer outerCancel()

	returnChan := make(chan checkResult)

	go func() {
		success, message := runFunc(checkCtx, check.Target, check.Command, check.ExpectedOutput, check.Credentials.Username, check.Credentials.Password)
		err := recover()
		if err != nil {
			logrus.WithError(err.(error)).Errorf("Failed to run check: %v, due to panic: %v", check, err)
			returnChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("check panicked: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		returnChan <- checkResult{
			Success: success,
			Message: message,
			Team:    team,
			Check:   check,
		}

	}()

	select {
	case <-outerCtx.Done():
		resultsChan <- checkResult{
			Success: false,
			Team:    team,
			Check:   check,
			Message: fmt.Sprintf("check timed out without return: %s", outerCtx.Err().Error()),
		}
		return
	case results := <-returnChan:
		resultsChan <- results
		return

	}
}
