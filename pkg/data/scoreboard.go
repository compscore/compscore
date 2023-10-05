package data

import (
	"bytes"
	"text/template"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/sirupsen/logrus"
)

type scoreboard_s struct{}

var Scoreboard scoreboard_s

func (*scoreboard_s) round(round_number int) (*structs.Scoreboard, error) {
	scoreboard := structs.Scoreboard{}
	scoreboard.Scores = make([]int, config.Teams.Amount)

	scoreboard.Round = round_number
	for _, configCheck := range config.Checks {
		scoreboardCheck := structs.Check{}
		scoreboardCheck.Name = configCheck.Name

		entStatus, err := Status.getAllByRoundAndCheckWithEdges(round_number, configCheck.Name)
		if err != nil {
			return nil, err
		}

		statuses := make([]int, config.Teams.Amount)
		for i, entStat := range entStatus {
			switch entStat.Status {
			case status.StatusDown:
				statuses[i] = 0
			case status.StatusUp:
				statuses[i] = 1
			case status.StatusUnknown:
				statuses[i] = 2
			}
		}
		scoreboardCheck.Status = statuses
		scoreboard.Checks = append(scoreboard.Checks, scoreboardCheck)
	}

	for i := 0; i < config.Teams.Amount; i++ {
		score, err := Team.getScoreBeforeRound(int8(i+1), round_number)
		if err != nil {
			return nil, err
		}
		scoreboard.Scores[i] = score
	}

	return &scoreboard, nil
}

func (*scoreboard_s) Round(round_number int) (*structs.Scoreboard, error) {
	mutex.Lock()
	logrus.Trace("scoreboard_s.Round: lock")
	defer mutex.Unlock()

	return Scoreboard.round(round_number)
}

func (*scoreboard_s) main() (*structs.Scoreboard, error) {
	entRound, err := Round.getLastCompleteRound()
	if err != nil {
		return nil, err
	}
	return Scoreboard.round(entRound.Number)
}

func (*scoreboard_s) Main() (*structs.Scoreboard, error) {
	mutex.Lock()
	logrus.Trace("scoreboard_s.Main: lock")
	defer mutex.Unlock()

	return Scoreboard.main()
}

func (*scoreboard_s) team(team_number int8, rounds int) (*structs.TeamScoreboard, error) {
	teamScoreboard := structs.TeamScoreboard{}
	teamScoreboard.Checks = make([]structs.Check, len(config.Checks))

	entRound, err := Round.getLastRound()
	if err != nil {
		return nil, err
	}

	teamScoreboard.Round = entRound.Number

	for i, configCheck := range config.Checks {
		teamScoreboard.Checks[i].Name = configCheck.Name
		teamScoreboard.Checks[i].Status = make([]int, rounds)

		entStatus, err := Status.getAllByCheckAndTeamWithLimit(configCheck.Name, team_number, rounds)
		if err != nil {
			return nil, err
		}

		for j, entStat := range entStatus {
			switch entStat.Status {
			case status.StatusDown:
				teamScoreboard.Checks[i].Status[j] = 0
			case status.StatusUp:
				teamScoreboard.Checks[i].Status[j] = 1
			case status.StatusUnknown:
				teamScoreboard.Checks[i].Status[j] = 2
			}
		}
	}

	return &teamScoreboard, nil
}

func (*scoreboard_s) Team(team_number int8, rounds int) (*structs.TeamScoreboard, error) {
	mutex.Lock()
	logrus.Trace("scoreboard_s.Team: lock")
	defer mutex.Unlock()

	return Scoreboard.team(team_number, rounds)
}

func (*scoreboard_s) check(check_name string, rounds int) (*structs.CheckScoreboard, error) {
	checkScoreboard := structs.CheckScoreboard{}
	checkScoreboard.Teams = make([]structs.Check, config.Teams.Amount)

	entRound, err := Round.getLastRound()
	if err != nil {
		return nil, err
	}

	checkScoreboard.Round = entRound.Number

	teamNameTemplate, err := template.New("Name Template").Parse(config.Teams.NameFormat)
	if err != nil {
		return nil, err
	}

	for i := 0; i < config.Teams.Amount; i++ {
		output := bytes.NewBuffer([]byte{})
		teamNameTemplate.Execute(output, struct{ Team int }{Team: i + 1})

		checkScoreboard.Teams[i].Name = output.String()
		checkScoreboard.Teams[i].Status = make([]int, rounds)

		entStatus, err := Status.getAllByCheckAndTeamWithLimit(check_name, int8(i+1), rounds)
		if err != nil {
			return nil, err
		}

		for j, entStat := range entStatus {
			switch entStat.Status {
			case status.StatusDown:
				checkScoreboard.Teams[i].Status[j] = 0
			case status.StatusUp:
				checkScoreboard.Teams[i].Status[j] = 1
			case status.StatusUnknown:
				checkScoreboard.Teams[i].Status[j] = 2
			}
		}
	}

	return &checkScoreboard, nil
}

func (*scoreboard_s) Check(check_name string, rounds int) (*structs.CheckScoreboard, error) {
	mutex.Lock()
	logrus.Trace("scoreboard_s.Check: lock")
	defer mutex.Unlock()

	return Scoreboard.check(check_name, rounds)
}

func (*scoreboard_s) history(check_name string, team_number int8, rounds int) (*[]structs.Status, error) {
	entStatus, err := Status.getAllByCheckAndTeamWithEdgesWithLimit(check_name, team_number, rounds)
	if err != nil {
		return nil, err
	}

	statuses := make([]structs.Status, len(entStatus))
	for i, entStat := range entStatus {
		statuses[i].Round = entStat.Edges.Round.Number
		statuses[i].Error = entStat.Error
		statuses[i].Time = entStat.Time.Format("2006-01-02 15:04:05")

		switch entStat.Status {
		case status.StatusDown:
			statuses[i].Status = 0
		case status.StatusUp:
			statuses[i].Status = 1
		case status.StatusUnknown:
			statuses[i].Status = 2
		}
	}

	return &statuses, nil
}

func (*scoreboard_s) History(check_name string, team_number int8, rounds int) (*[]structs.Status, error) {
	mutex.Lock()
	logrus.Trace("scoreboard_s.History: lock")
	defer mutex.Unlock()

	return Scoreboard.history(check_name, team_number, rounds)
}
