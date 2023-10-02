package data

import (
	"bytes"
	"text/template"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/compscore/compscore/pkg/structs"
)

type status_s struct{}

var Status = status_s{}

func (*status_s) Get(roundNumber int, checkName string, teamNumber int8) (*ent.Status, error) {
	exist, err := Status.Exists(roundNumber, checkName, teamNumber)
	if err != nil {
		return nil, err
	}

	if !exist {
		return Status.Create(roundNumber, checkName, teamNumber, status.StatusUnknown)
	}

	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Only(Ctx)
}

func (*status_s) GetWithEdges(roundNumber int, checkName string, teamNumber int8) (*ent.Status, error) {
	exist, err := Status.Exists(roundNumber, checkName, teamNumber)
	if err != nil {
		return nil, err
	}

	if !exist {
		return Status.Create(roundNumber, checkName, teamNumber, status.StatusUnknown)
	}

	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Only(Ctx)
}

func (*status_s) GetComplex(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	exist, err := Status.Exists(entRound.Number, entCheck.Name, entTeam.Number)
	if err != nil {
		return nil, err
	}

	if !exist {
		return Status.Create(entRound.Number, entCheck.Name, entTeam.Number, status.StatusUnknown)
	}

	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.IDEQ(entRound.ID),
			),
			status.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
			status.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
		).Only(Ctx)
}

func (*status_s) GetComplexWithEdges(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	exist, err := Status.Exists(entRound.Number, entCheck.Name, entTeam.Number)
	if err != nil {
		return nil, err
	}

	if !exist {
		return Status.Create(entRound.Number, entCheck.Name, entTeam.Number, status.StatusUnknown)
	}

	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.IDEQ(entRound.ID),
			),
			status.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
			status.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
		).Only(Ctx)
}

func (*status_s) GetAll() ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllWithEdges() ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByRound(roundNumber int) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
		).
		Order(
			ent.Asc(team.FieldNumber),
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundWithEdges(roundNumber int) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
		).
		Order(
			ent.Asc(team.FieldNumber),
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheck(checkName string) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheckWithEdges(checkName string) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByTeam(teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByTeamWithEdges(teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndCheck(roundNumber int, checkName string) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndCheckWithEdges(roundNumber int, checkName string) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndTeam(roundNumber int, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndTeamWithEdges(roundNumber int, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheckAndTeam(checkName string, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheckAndTeamWithEdges(checkName string, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) Exists(roundNumber int, checkName string, teamNumber int8) (bool, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Exist(Ctx)
}

func (*status_s) Create(roundNumber int, checkName string, teamNumber int8, stat status.Status) (*ent.Status, error) {
	exist, err := Status.Exists(roundNumber, checkName, teamNumber)
	if err != nil {
		return nil, err
	}

	if exist {
		return Status.Get(roundNumber, checkName, teamNumber)
	}

	entRound, err := Round.Get(roundNumber)
	if err != nil {
		return nil, err
	}

	entCheck, err := Check.Get(checkName)
	if err != nil {
		return nil, err
	}

	entTeam, err := Team.Get(teamNumber)
	if err != nil {
		return nil, err
	}

	return Client.Status.
		Create().
		SetRound(entRound).
		SetCheck(entCheck).
		SetTeam(entTeam).
		SetStatus(stat).
		Save(Ctx)
}

func (*status_s) UpdateComplex(entStatus *ent.Status, statusEnum status.Status, message string) (*ent.Status, error) {
	return entStatus.Update().
		SetStatus(statusEnum).
		SetError(message).
		SetTime(time.Now()).
		Save(Ctx)
}

func (*status_s) Update(teamNumber int8, roundNumber int, checkName string, statusEnum status.Status, message string) (int, error) {
	return Client.Status.Update().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		SetStatus(statusEnum).
		SetError(message).
		SetTime(time.Now()).
		Save(Ctx)
}

func (*status_s) Delete(entStatus *ent.Status) error {
	return Client.Status.
		DeleteOne(entStatus).
		Exec(Ctx)
}

func (*status_s) Scoreboard() (*structs.Scoreboard, error) {
	scoreboard := structs.Scoreboard{}
	scoreboard.Scores = make([]int, config.Teams.Amount)

	entRound, err := Client.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		Offset(1).
		All(Ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Round = entRound[0].Number

	for _, configCheck := range config.Checks {
		scoreboardCheck := structs.Check{}
		scoreboardCheck.Name = configCheck.Name

		entStatus, err := Client.Status.Query().
			WithRound().
			Where(
				status.HasRoundWith(
					round.NumberEQ(scoreboard.Round),
				),
				status.HasCheckWith(
					check.NameEQ(configCheck.Name),
				),
			).
			Order(
				status.ByRoundField(
					round.FieldNumber,
					sql.OrderAsc(),
				),
			).
			All(Ctx)
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
		score, err := Team.GetScore(int8(i + 1))
		if err != nil {
			return nil, err
		}
		scoreboard.Scores[i] = score
	}

	return &scoreboard, nil
}

func (*status_s) TeamScoreboard(team_number int8, rounds int) (*structs.TeamScoreboard, error) {
	teamScoreboard := structs.TeamScoreboard{}
	teamScoreboard.Checks = make([]structs.Check, len(config.Checks))

	entRound, err := Round.GetLastRound()
	if err != nil {
		return nil, err
	}

	teamScoreboard.Round = entRound.Number

	for i, configCheck := range config.Checks {
		teamScoreboard.Checks[i].Name = configCheck.Name
		teamScoreboard.Checks[i].Status = make([]int, rounds)

		entStatuses, err := Client.Status.Query().
			Where(
				status.HasCheckWith(
					check.NameEQ(configCheck.Name),
				),
				status.HasTeamWith(
					team.NumberEQ(team_number),
				),
			).
			Order(
				status.ByRoundField(
					round.FieldNumber,
					sql.OrderDesc(),
				),
			).
			Limit(rounds).
			All(Ctx)
		if err != nil {
			return nil, err
		}

		for j, entStatus := range entStatuses {
			switch entStatus.Status {
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

func (*status_s) CheckScoreboard(check_name string, rounds int) (*structs.CheckScoreboard, error) {
	checkScoreboard := structs.CheckScoreboard{}
	checkScoreboard.Teams = make([]structs.Check, config.Teams.Amount)

	entRound, err := Round.GetLastRound()
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

		entStatuses, err := Client.Status.Query().
			Where(
				status.HasCheckWith(
					check.NameEQ(check_name),
				),
				status.HasTeamWith(
					team.NumberEQ(int8(i+1)),
				),
			).
			Order(
				status.ByRoundField(
					round.FieldNumber,
					sql.OrderDesc(),
				),
			).
			Limit(rounds).
			All(Ctx)
		if err != nil {
			return nil, err
		}

		for j, entStatus := range entStatuses {
			switch entStatus.Status {
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
