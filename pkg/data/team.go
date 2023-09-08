package data

import (
	"fmt"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/team"
)

type team_s struct{}

var Team = team_s{}

func (*team_s) Get(number int8) (*ent.Team, error) {
	exists, err := Team.Exists(number)
	if err != nil {
		return nil, err
	}

	if !exists {
		var teamName string

		for _, team := range config.RunningConfig.Teams {
			if team.Number == number {
				teamName = team.Name
				break
			}
		}
		if teamName == "" {
			return nil, fmt.Errorf("unable to find name for team: %d", number)
		}

		return Team.Create(number, teamName)
	}

	return Client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).Only(Ctx)
}

func (*team_s) GetByName(name string) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).Only(Ctx)
}

func (*team_s) GetByNumberWithStatus(number int8) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		WithStatus().
		Only(Ctx)
}

func (*team_s) GetByNameWithStatus(name string) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		WithStatus().
		Only(Ctx)
}

func (*team_s) GetAll() ([]*ent.Team, error) {
	return Client.Team.
		Query().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*team_s) GetAllWithStatus() ([]*ent.Team, error) {
	return Client.Team.
		Query().
		WithStatus().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*team_s) Exists(number int8) (bool, error) {
	return Client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).Exist(Ctx)
}

func (*team_s) Create(number int8, name string) (*ent.Team, error) {
	exist, err := Team.Exists(number)
	if err != nil {
		return nil, err
	}

	if exist {
		return Team.Get(number)
	}

	return Client.Team.
		Create().
		SetNumber(number).
		SetName(name).
		Save(Ctx)
}

func (*team_s) Update(team *ent.Team, number int8, name string) (*ent.Team, error) {
	return team.Update().
		SetNumber(number).
		SetName(name).
		Save(Ctx)
}

func (*team_s) Delete(team *ent.Team) error {
	return Client.Team.
		DeleteOne(team).
		Exec(Ctx)
}
