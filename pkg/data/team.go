package data

import (
	"fmt"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type team_s struct{}

var Team = team_s{}

func (*team_s) exists(number int) (bool, error) {
	return client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		Exist(ctx)
}

func (*team_s) Exists(number int) (bool, error) {
	mutex.Lock()
	logrus.Trace("team_s.Exists: lock")
	defer mutex.Unlock()

	return Team.exists(number)
}

func existsByName(name string) (bool, error) {
	return client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		Exist(ctx)
}

func (*team_s) ExistsByName(name string) (bool, error) {
	mutex.Lock()
	logrus.Trace("team_s.ExistsByName: lock")
	defer mutex.Unlock()

	return existsByName(name)
}

func (*team_s) get(number int) (*ent.Team, error) {
	exists, err := Team.exists(number)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("team %d does not exist", number)
	}

	return client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).Only(ctx)
}

func (*team_s) Get(number int) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.Get: lock")
	defer mutex.Unlock()

	return Team.get(number)
}

func (*team_s) getByName(name string) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).Only(ctx)
}

func (*team_s) GetByName(name string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByName: lock")
	defer mutex.Unlock()

	return Team.getByName(name)
}

func (*team_s) getByNumberWithStatus(number int) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		WithStatus().
		Only(ctx)
}

func (*team_s) GetByNumberWithStatus(number int) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByNumberWithStatus: lock")
	defer mutex.Unlock()

	return Team.getByNumberWithStatus(number)
}

func (*team_s) getByNameWithStatus(name string) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		WithStatus().
		Only(ctx)
}

func (*team_s) GetByNameWithStatus(name string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByNameWithStatus: lock")
	defer mutex.Unlock()

	return Team.getByNameWithStatus(name)
}

func (*team_s) getByNumberWithCredentials(number int) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		WithCredential().
		Only(ctx)
}

func (*team_s) GetByNumberWithCredentials(number int) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByNumberWithCredentials: lock")
	defer mutex.Unlock()

	return Team.getByNumberWithCredentials(number)
}

func (*team_s) getByNameWithCredentials(name string) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		WithCredential().
		Only(ctx)
}

func (*team_s) GetByNameWithCredentials(name string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByNameWithCredentials: lock")
	defer mutex.Unlock()

	return Team.getByNameWithCredentials(name)
}

func (*team_s) getByNumberWithEdges(number int) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NumberEQ(number),
		).
		WithCredential().
		WithStatus().
		Only(ctx)
}

func (*team_s) GetByNumberWithEdges(number int) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByNumberWithEdges: lock")
	defer mutex.Unlock()

	return Team.getByNumberWithEdges(number)
}

func (*team_s) getByNameWithEdges(name string) (*ent.Team, error) {
	return client.Team.
		Query().
		Where(
			team.NameEQ(name),
		).
		WithCredential().
		WithStatus().
		Only(ctx)
}

func (*team_s) GetByNameWithEdges(name string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetByNameWithEdges: lock")
	defer mutex.Unlock()

	return Team.getByNameWithEdges(name)
}

func (*team_s) getAll() ([]*ent.Team, error) {
	return client.Team.
		Query().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(ctx)
}

func (*team_s) GetAll() ([]*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetAll: lock")
	defer mutex.Unlock()

	return Team.getAll()
}

func (*team_s) getAllWithStatus() ([]*ent.Team, error) {
	return client.Team.
		Query().
		WithStatus().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(ctx)
}

func (*team_s) GetAllWithStatus() ([]*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetAllWithStatus: lock")
	defer mutex.Unlock()

	return Team.getAllWithStatus()
}

func (*team_s) getAllWithCredentials() ([]*ent.Team, error) {
	return client.Team.
		Query().
		WithCredential().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(ctx)
}

func (*team_s) GetAllWithCredentials() ([]*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetAllWithCredentials: lock")
	defer mutex.Unlock()

	return Team.getAllWithCredentials()
}

func (*team_s) getAllWithEdges() ([]*ent.Team, error) {
	return client.Team.
		Query().
		WithStatus().
		WithCredential().
		Order(
			ent.Asc(team.FieldNumber),
		).
		All(ctx)
}

func (*team_s) GetAllWithEdges() ([]*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetAllWithEdges: lock")
	defer mutex.Unlock()

	return Team.getAllWithEdges()
}

func (*team_s) create(number int, name string, password string) (*ent.Team, error) {
	exists, err := Team.exists(number)
	if err != nil {
		return nil, err
	}

	if exists {
		return Team.get(number)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return client.Team.
		Create().
		SetNumber(number).
		SetName(name).
		SetPassword(string(hashedPassword)).
		Save(ctx)
}

func (*team_s) Create(number int, name string, password string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.Create: lock")
	defer mutex.Unlock()

	return Team.create(number, name, password)
}

func (*team_s) createAdminUser(username string, password string) (*ent.Team, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return client.Team.
		Create().
		SetName(username).
		SetPassword(string(hashedPassword)).
		SetRole(team.RoleAdmin).
		Save(ctx)
}

func (*team_s) CreateAdminUser(username string, password string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.CreateAdminUser: lock")
	defer mutex.Unlock()

	return Team.createAdminUser(username, password)
}

func (*team_s) getRole(name string) (string, error) {
	teamEnt, err := Team.getByName(name)
	if err != nil {
		return "", err
	}

	return string(teamEnt.Role), nil
}

func (*team_s) GetRole(name string) (string, error) {
	mutex.Lock()
	logrus.Trace("team_s.GetTeamRole: lock")
	defer mutex.Unlock()

	return Team.getRole(name)
}

func (*team_s) update(team *ent.Team, number int, name string) (*ent.Team, error) {
	return team.Update().
		SetNumber(number).
		SetName(name).
		Save(ctx)
}

func (*team_s) Update(team *ent.Team, number int, name string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.Update: lock")
	defer mutex.Unlock()

	return Team.update(team, number, name)
}

func (*team_s) updatePassword(team *ent.Team, password string) (*ent.Team, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return team.Update().
		SetPassword(string(hashedPassword)).
		Save(ctx)
}

func (*team_s) UpdatePassword(team *ent.Team, password string) (*ent.Team, error) {
	mutex.Lock()
	logrus.Trace("team_s.UpdatePassword: lock")
	defer mutex.Unlock()

	return Team.updatePassword(team, password)
}

func (*team_s) checkPassword(team int, password string) (bool, error) {
	teamEnt, err := Team.get(team)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(teamEnt.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (*team_s) CheckPassword(team int, password string) (bool, error) {
	mutex.Lock()
	logrus.Trace("team_s.CheckPassword: lock")
	defer mutex.Unlock()

	return Team.checkPassword(team, password)
}

func (*team_s) checkPasswordByName(team_name string, password string) (bool, error) {
	teamEnt, err := Team.getByName(team_name)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(teamEnt.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (*team_s) CheckPasswordByName(team_name string, password string) (bool, error) {
	mutex.Lock()
	logrus.Trace("team_s.CheckPasswordByName: lock")
	defer mutex.Unlock()

	return Team.checkPasswordByName(team_name, password)
}

func (*team_s) delete(team *ent.Team) error {
	return client.Team.
		DeleteOne(team).
		Exec(ctx)
}

func (*team_s) Delete(team *ent.Team) error {
	mutex.Lock()
	logrus.Trace("team_s.Delete: lock")
	defer mutex.Unlock()

	return Team.delete(team)
}

func getScore(teamNumber int) (score int, err error) {
	for _, configCheck := range config.Checks {
		count, err := client.Status.
			Query().
			Where(
				status.HasCheckWith(
					check.NameEQ(configCheck.Name),
				),
				status.HasTeamWith(
					team.NumberEQ(teamNumber),
				),
				status.StatusEQ(status.StatusUp),
			).
			Count(ctx)
		if err != nil {
			return 0, err
		}

		score += count * configCheck.Weight
	}

	return score, nil
}

func (*team_s) GetScore(teamNumber int) (score int, err error) {
	mutex.Lock()
	logrus.Trace("team_s.GetScore: lock")
	defer mutex.Unlock()

	return getScore(teamNumber)
}

func (*team_s) getScoreBeforeRound(team_number int, round_number int) (score int, err error) {
	for _, configCheck := range config.Checks {
		count, err := client.Status.
			Query().
			Where(
				status.HasCheckWith(
					check.NameEQ(configCheck.Name),
				),
				status.HasTeamWith(
					team.NumberEQ(team_number),
				),
				status.StatusEQ(status.StatusUp),
				status.HasRoundWith(
					round.NumberLT(round_number),
				),
			).
			Count(ctx)
		if err != nil {
			return 0, err
		}

		score += count * configCheck.Weight
	}

	return score, nil
}

func (*team_s) GetScoreBeforeRound(team_number int, round_number int) (score int, err error) {
	mutex.Lock()
	logrus.Trace("team_s.GetScoreBeforeRound: lock")
	defer mutex.Unlock()

	return Team.getScoreBeforeRound(team_number, round_number)
}
