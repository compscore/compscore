package data

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/sirupsen/logrus"
)

type check_s struct{}

var Check = check_s{}

func (*check_s) create(name string, weight int) (*ent.Check, error) {
	return client.Check.
		Create().
		SetName(name).
		SetWeight(weight).
		Save(ctx)
}

func (*check_s) Create(name string, weight int) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.Create: lock")
		defer mutex.Unlock()
	}

	return Check.create(name, weight)
}

func (*check_s) exists(name string) (bool, error) {
	return client.Check.
		Query().
		Where(
			check.NameEQ(name),
		).
		Exist(ctx)
}

func (*check_s) Exists(name string) (bool, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.Exists: lock")
		defer mutex.Unlock()
	}

	return Check.exists(name)
}

func (*check_s) get(name string) (*ent.Check, error) {
	exists, err := Check.exists(name)
	if err != nil {
		return nil, err
	}

	if !exists {
		for _, check := range config.Checks {
			if name == check.Name {
				return Check.create(name, check.Weight)
			}
		}
	}

	return client.Check.
		Query().
		Where(
			check.NameEQ(name),
		).Only(ctx)
}

func (*check_s) Get(name string) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.Get: lock")
		defer mutex.Unlock()
	}

	return Check.get(name)
}

func (*check_s) getWithStatus(name string) (*ent.Check, error) {
	return client.Check.
		Query().
		WithStatus().
		Where(
			check.NameEQ(name),
		).Only(ctx)
}

func (*check_s) GetWithStatus(name string) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.GetWithStatus: lock")
		defer mutex.Unlock()
	}

	return Check.getWithStatus(name)
}

func (*check_s) getWithTeamCredential(name string, team_number int) (*ent.Check, error) {
	return client.Check.
		Query().
		WithCredential(
			func(q *ent.CredentialQuery) {
				q.WithTeam().
					Where(
						credential.HasTeamWith(
							team.NumberEQ(team_number),
						),
					)
			},
		).
		Where(
			check.NameEQ(name),
		).Only(ctx)
}

func (*check_s) GetWithTeamCredential(name string, team_number int) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.GetWithTeamCredenital: lock")
		defer mutex.Unlock()
	}

	return Check.getWithTeamCredential(name, team_number)
}

func (*check_s) getWithCredentials(name string) (*ent.Check, error) {
	return client.Check.
		Query().
		WithCredential().
		Where(
			check.NameEQ(name),
		).Only(ctx)
}

func (*check_s) GetWithCredentials(name string) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.GetWithCredentials: lock")
		defer mutex.Unlock()
	}

	return Check.getWithCredentials(name)
}

func (*check_s) getWithEdges(name string) (*ent.Check, error) {
	return client.Check.
		Query().
		WithCredential().
		WithStatus().
		Where(
			check.NameEQ(name),
		).Only(ctx)
}

func (*check_s) GetWithEdges(name string) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.GetWithEdges: lock")
		defer mutex.Unlock()
	}

	return Check.getWithEdges(name)
}

func (*check_s) getAll() ([]*ent.Check, error) {
	return client.Check.
		Query().
		Order(
			ent.Asc(check.FieldName),
		).
		All(ctx)
}

func (*check_s) GetAll() ([]*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.GetAll: lock")
		defer mutex.Unlock()
	}

	return Check.getAll()
}

func (*check_s) getAllWithStatus() ([]*ent.Check, error) {
	return client.Check.
		Query().
		WithStatus().
		Order(
			ent.Asc(check.FieldName),
		).
		All(ctx)
}

func (*check_s) GetAllWithStatus() ([]*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.GetAllWithStatus: lock")
		defer mutex.Unlock()
	}

	return Check.getAllWithStatus()
}

func (*check_s) update(check *ent.Check) (*ent.Check, error) {
	return check.Update().
		SetName(check.Name).
		Save(ctx)
}

func (*check_s) Update(check *ent.Check) (*ent.Check, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.Update: lock")
		defer mutex.Unlock()
	}

	return Check.update(check)
}

func (*check_s) delete(check *ent.Check) error {
	return client.Check.
		DeleteOne(check).
		Exec(ctx)
}

func (*check_s) Delete(check *ent.Check) error {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("check_s.Delete: lock")
		defer mutex.Unlock()
	}

	return Check.delete(check)
}
