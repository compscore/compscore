package engine

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"text/template"
	"time"

	"github.com/compscore/compscore/pkg/checks/imports"
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
	interval := config.Scoring.Interval
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	defer runLock.Unlock()
	defer func() {
		Status = proto.StatusEnum_PAUSED
	}()

	for {
		select {
		case <-ticker.C:
			if interval != config.Scoring.Interval {
				interval = config.Scoring.Interval
				ticker = time.NewTicker(time.Duration(config.Scoring.Interval) * time.Second)
			}

			if Status != proto.StatusEnum_RUNNING {
				return
			}

			roundCount, err := data.Round.Count()
			if err != nil {
				logrus.WithError(err).Error("Failed to get round count")
				return
			}

			if roundCount == 0 {
				_, err = data.Round.Create(1)
				if err != nil {
					logrus.WithError(err).Error("Failed to create first round")
					return
				}
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

	entRound, err := data.Round.CreateNextRound()
	if err != nil {
		return err
	}

	defer func() {
		entRound, err := data.Round.Complete(entRound.Number)
		if err != nil {
			logrus.WithError(err).Errorf("Failed to complete round: %v", entRound)
		}
	}()

	checks := 0

	for i := 1; i <= config.Teams.Amount; i++ {
		for _, check := range config.Checks {
			_, err := data.Status.Create(
				entRound.Number,
				check.Name,
				int8(i),
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

	for _, check := range config.Checks {
		targetTemplate, err := template.New(check.Name).Parse(check.Target)
		if err != nil {
			for i := 1; i <= config.Teams.Amount; i++ {
				_, err = data.Status.Update(
					int8(i),
					entRound.Number,
					check.Name,
					status.StatusDown,
					fmt.Sprintf("failed to parse target template: %v", err),
				)
				if err != nil {
					logrus.WithError(err).Errorf("Failed to update status for check: %v", check)
				}
				continue
			}
		}
		for i := 1; i <= config.Teams.Amount; i++ {
			target := bytes.NewBuffer([]byte{})
			err = targetTemplate.Execute(target, struct{ Team int8 }{Team: int8(i)})
			if err != nil {
				_, err = data.Status.Update(
					int8(i),
					entRound.Number,
					check.Name,
					status.StatusDown,
					fmt.Sprintf("failed to parse target template: %v", err),
				)
				if err != nil {
					logrus.WithError(err).Errorf("Failed to update status for check: %v", check)
				}
				continue
			}

			entCredential, err := data.Credential.Get(int8(i), check.Name)
			if err != nil {
				_, err = data.Status.Update(
					int8(i),
					entRound.Number,
					check.Name,
					status.StatusDown,
					fmt.Sprintf("failed to get credential: %v", err),
				)
				if err != nil {
					logrus.WithError(err).Errorf("Failed to update status for check: %v", check)
				}
				continue
			}

			go runScoreCheck(entRound.Number, check, int8(i), target.String(), entCredential.Password, resultsChan, wgRound)
		}
	}

	go func() {
		for result := range resultsChan {
			_, err = data.Status.Update(
				result.Team,
				entRound.Number,
				result.Check.Name,
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

func runScoreCheck(round int, check structs.Check_s, team int8, target string, password string, resultsChan chan checkResult, wg *sync.WaitGroup) {
	defer wg.Done()

	runFunc, ok := imports.Imports[check.Release.Org+"/"+check.Release.Repo]
	if !ok {
		resultsChan <- checkResult{
			Success: false,
			Message: fmt.Sprintf("check not found: %s", check.Release.Org+"-"+check.Release.Repo),
			Team:    team,
			Check:   check,
		}

		logrus.Errorf("Failed to get check function: %v", check)
		return
	}

	checkCtx, checkCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(float64(config.Scoring.Interval)*0.9))
	defer checkCancel()

	outerCtx, outerCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(float64(config.Scoring.Interval)))
	defer outerCancel()

	returnChan := make(chan checkResult)

	go func() {
		success, message := runFunc(checkCtx, target, check.Command, check.ExpectedOutput, check.Credentials.Username, password, check.Options)
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
