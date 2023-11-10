package engine

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/compscore/compscore/pkg/checks/imports"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/team"
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
	Team    int
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

	checks := config.Teams.Amount * len(config.Checks)

	_, err = data.Client(
		func(client *ent.Client, ctx context.Context) (interface{}, error) {
			bulkStatusCreate := make([]*ent.StatusCreate, checks)

			entTeams, err := client.Team.Query().
				Where(
					team.NumberGTE(1),
				).
				Order(
					ent.Asc(team.FieldNumber),
				).
				All(context.Background())
			if err != nil {
				return nil, err
			}

			entChecks, err := client.Check.Query().
				Order(
					ent.Asc(check.FieldName),
				).
				All(context.Background())
			if err != nil {
				return nil, err
			}

			for i, entCheck := range entChecks {
				for j, entTeam := range entTeams {
					bulkStatusCreate[i*len(entTeams)+j] = client.Status.Create().
						SetRound(entRound).
						SetTeam(entTeam).
						SetCheck(entCheck).
						SetPoints(0).
						SetStatus(status.StatusUnknown)
				}
			}

			return client.Status.CreateBulk(bulkStatusCreate...).Save(ctx)
		},
	)
	if err != nil {
		return err
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
					i,
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
			err = targetTemplate.Execute(target, struct{ Team string }{Team: fmt.Sprintf("%02d", i)})
			if err != nil {
				_, err = data.Status.Update(
					i,
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

			entCredential, err := data.Credential.Get(i, check.Name)
			if err != nil {
				_, err = data.Status.Update(
					i,
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

			go runScoreCheck(entRound.Number, check, i, target.String(), entCredential.Password, resultsChan, wgRound)
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

func runScoreCheck(round int, check structs.Check_s, team int, target string, password string, resultsChan chan checkResult, wg *sync.WaitGroup) {
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

	var (
		command        string
		username       string
		expectedOutput string
	)

	if strings.Contains(check.Command, "{{") {
		commandTemplate, err := template.New(check.Name).Parse(check.Command)
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse command template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		output := bytes.NewBuffer([]byte{})
		err = commandTemplate.Execute(output, struct{ Team string }{Team: fmt.Sprintf("%02d", team)})
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse command template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		command = output.String()
	} else {
		command = check.Command
	}

	if strings.Contains(check.ExpectedOutput, "{{") {
		expectedOutputTemplate, err := template.New(check.Name).Parse(check.ExpectedOutput)
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse expected output template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		output := bytes.NewBuffer([]byte{})
		err = expectedOutputTemplate.Execute(output, struct{ Team string }{Team: fmt.Sprintf("%02d", team)})
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse expected output template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		expectedOutput = output.String()
	} else {
		expectedOutput = check.ExpectedOutput
	}

	if strings.Contains(check.Credentials.Username, "{{") {
		usernameTemplate, err := template.New(check.Name).Parse(check.Credentials.Username)
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse username template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		output := bytes.NewBuffer([]byte{})
		err = usernameTemplate.Execute(output, struct{ Team string }{Team: fmt.Sprintf("%02d", team)})
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse username template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		username = output.String()
	} else {
		username = check.Credentials.Username
	}

	if strings.Contains(check.Credentials.Password, "{{") {
		passwordTemplate, err := template.New(check.Name).Parse(check.Credentials.Password)
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse password template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		output := bytes.NewBuffer([]byte{})
		err = passwordTemplate.Execute(output, struct{ Team string }{Team: fmt.Sprintf("%02d", team)})
		if err != nil {
			resultsChan <- checkResult{
				Success: false,
				Message: fmt.Sprintf("failed to parse password template: %v", err),
				Team:    team,
				Check:   check,
			}
			return
		}

		password = output.String()
	} else {
		password = check.Credentials.Password
	}

	optionsCopy := make(map[string]interface{})

	for key, option := range check.Options {
		optionsCopy[key] = option

		optionStr, isStr := option.(string)
		if !isStr {
			continue
		}

		if strings.Contains(optionStr, "{{") {
			optionTemplate, err := template.New(check.Name).Parse(optionStr)
			if err != nil {
				resultsChan <- checkResult{
					Success: false,
					Message: fmt.Sprintf("failed to parse option template: %v", err),
					Team:    team,
					Check:   check,
				}
				return
			}

			output := bytes.NewBuffer([]byte{})

			err = optionTemplate.Execute(output, struct{ Team string }{Team: fmt.Sprintf("%02d", team)})
			if err != nil {
				resultsChan <- checkResult{
					Success: false,
					Message: fmt.Sprintf("failed to parse option template: %v", err),
					Team:    team,
					Check:   check,
				}
				return
			}

			optionsCopy[key] = output.String()
		}
	}

	go func() {
		success, message := runFunc(checkCtx, target, command, expectedOutput, username, password, optionsCopy)
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
