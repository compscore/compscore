package data

import (
	"fmt"
	"strings"

	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/team"
	"golang.org/x/crypto/bcrypt"
)

type team_s struct{}

var Team = team_s{}

func (*team_s) Get(number int8) (*ent.Team, error) {
	exists, err := Team.Exists(number)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("team %d does not exist", number)
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

func (*team_s) GetByNumberWithCredentials(number int8) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		WithCredential().
		Only(Ctx)
}

func (*team_s) GetByNameWithCredentials(name string) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		WithCredential().
		Only(Ctx)
}

func (*team_s) GetByNumberWithEdges(number int8) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		WithCredential().
		WithStatus().
		Only(Ctx)
}

func (*team_s) GetByNameWithEdges(name string) (*ent.Team, error) {
	return Client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		WithCredential().
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

func (*team_s) GetAllWithCredentials() ([]*ent.Team, error) {
	return Client.Team.
		Query().
		WithCredential().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*team_s) GetAllWithEdges() ([]*ent.Team, error) {
	return Client.Team.
		Query().
		WithStatus().
		WithCredential().
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

func (*team_s) Create(number int8, name string, password string) (*ent.Team, error) {
	exist, err := Team.Exists(number)
	if err != nil {
		return nil, err
	}

	if exist {
		return Team.Get(number)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return Client.Team.
		Create().
		SetNumber(number).
		SetName(name).
		SetPassword(string(hashedPassword)).
		Save(Ctx)
}

func (*team_s) Update(team *ent.Team, number int8, name string) (*ent.Team, error) {
	return team.Update().
		SetNumber(number).
		SetName(name).
		Save(Ctx)
}

func (*team_s) UpdatePassword(team *ent.Team, password string) (*ent.Team, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return team.Update().
		SetPassword(string(hashedPassword)).
		Save(Ctx)
}

func (*team_s) CheckPassword(team int8, password string) (bool, error) {
	teamEnt, err := Team.Get(team)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(teamEnt.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (*team_s) CheckPasswordByName(team_name string, password string) (bool, error) {
	team_name = strings.Replace(team_name, "_", " ", -1)
	teamEnt, err := Team.GetByName(team_name)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(teamEnt.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (*team_s) Delete(team *ent.Team) error {
	return Client.Team.
		DeleteOne(team).
		Exec(Ctx)
}
