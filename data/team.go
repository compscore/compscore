package data

import (
	"github.com/compscore/compscore/ent"
	"github.com/compscore/compscore/ent/team"
)

type team_s struct{}

func (t *team_s) Get(number int) (*ent.Team, error) {
	return Client.Team.Query().
		Where(
			team.Number(number),
		).Only(Ctx)
}

func (t *team_s) GetWithCheckLogs(number int) (*ent.Team, error) {
	return Client.Team.Query().
		Where(
			team.Number(number),
		).
		WithChecklogs().
		Only(Ctx)
}

func (t *team_s) Create(teamNumber int, teamName string) (*ent.Team, error) {
	return Client.Team.Create().
		SetNumber(teamNumber).
		SetName(teamName).
		Save(Ctx)
}

func (t *team_s) UpdateName(teamNumber int, teamName string) (int, error) {
	return Client.Team.Update().
		Where(
			team.Number(teamNumber),
		).
		SetName(teamName).
		Save(Ctx)
}

func (t *team_s) Delete(teamNumber int) (int, error) {
	return Client.Team.Delete().
		Where(
			team.Number(teamNumber),
		).Exec(Ctx)
}

func (t *team_s) DeleteAll() (int, error) {
	return Client.Team.Delete().
		Exec(Ctx)
}
