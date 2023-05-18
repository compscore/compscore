package data

import (
	"fmt"

	"github.com/compscore/compscore/ent"
	"github.com/compscore/compscore/ent/check"
	"github.com/compscore/compscore/ent/checklog"
	"github.com/compscore/compscore/ent/round"
	"github.com/compscore/compscore/ent/team"
)

type checklog_s struct{}

func (c *checklog_s) Get(checkName string, teamNumber int, roundNumber int) (*ent.CheckLog, error) {
	return Client.CheckLog.Query().
		Where(
			checklog.HasCheckWith(
				check.Name(checkName),
			),
			checklog.HasTeamWith(
				team.Number(teamNumber),
			),
			checklog.HasRoundWith(
				round.Number(roundNumber),
			),
		).
		WithCheck().
		WithTeam().
		WithRound().
		Only(Ctx)
}

func (c *checklog_s) Create(checkName string, teamNumber int, roundNumber int, status bool, statusErr error) (*ent.CheckLog, error) {
	entCheck, err := Check.Get(checkName)
	if err != nil {
		return nil, err
	}

	entTeam, err := Team.Get(teamNumber)
	if err != nil {
		return nil, err
	}

	entRound, err := Round.Get(roundNumber)
	if err != nil {
		return nil, err
	}

	exist, err := Client.CheckLog.Query().
		Where(
			checklog.HasCheckWith(
				check.Name(checkName),
			),
			checklog.HasTeamWith(
				team.Number(teamNumber),
			),
			checklog.HasRoundWith(
				round.Number(roundNumber),
			),
		).Exist(Ctx)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, fmt.Errorf("CheckLog already exists; Check: \"%s\", Team: \"%d\", Round: \"%d\"", checkName, teamNumber, roundNumber)
	}

	return Client.CheckLog.Create().
		SetStatus(status).
		SetError(
			func() string {
				if statusErr != nil {
					return statusErr.Error()
				}
				return ""
			}(),
		).
		SetCheck(entCheck).
		SetTeam(entTeam).
		SetRound(entRound).
		Save(Ctx)
}

func (c *checklog_s) Update(checkName string, teamNumber int, roundNumber int, status bool, statusErr error) (int, error) {
	return Client.CheckLog.Update().
		Where(
			checklog.HasCheckWith(
				check.Name(checkName),
			),
			checklog.HasTeamWith(
				team.Number(teamNumber),
			),
			checklog.HasRoundWith(
				round.Number(roundNumber),
			),
		).
		SetStatus(status).
		SetError(
			func() string {
				if statusErr != nil {
					return statusErr.Error()
				}
				return ""
			}(),
		).
		Save(Ctx)
}

func (c *checklog_s) UpdateStatus(checkName string, teamNumber int, roundNumber int, status bool) (int, error) {
	return Client.CheckLog.Update().
		Where(
			checklog.HasCheckWith(
				check.Name(checkName),
			),
			checklog.HasTeamWith(
				team.Number(teamNumber),
			),
			checklog.HasRoundWith(
				round.Number(roundNumber),
			),
		).
		SetStatus(status).
		Save(Ctx)
}

func (c *checklog_s) UpdateError(checkName string, teamNumber int, roundNumber int, statusErr error) (int, error) {
	return Client.CheckLog.Update().
		Where(
			checklog.HasCheckWith(
				check.Name(checkName),
			),
			checklog.HasTeamWith(
				team.Number(teamNumber),
			),
			checklog.HasRoundWith(
				round.Number(roundNumber),
			),
		).
		SetError(
			func() string {
				if statusErr != nil {
					return statusErr.Error()
				}
				return ""
			}(),
		).
		Save(Ctx)
}

func (c *checklog_s) Delete(checkName string, teamNumber int, roundNumber int) (int, error) {
	return Client.CheckLog.Delete().
		Where(
			checklog.HasCheckWith(
				check.Name(checkName),
			),
			checklog.HasTeamWith(
				team.Number(teamNumber),
			),
			checklog.HasRoundWith(
				round.Number(roundNumber),
			),
		).
		Exec(Ctx)
}

func (c *checklog_s) DeleteAll() (int, error) {
	return Client.CheckLog.Delete().
		Exec(Ctx)
}
