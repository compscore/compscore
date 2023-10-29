package data

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/sirupsen/logrus"
)

type credential_s struct{}

var Credential = credential_s{}

func (*credential_s) get(team_id int, check_name string) (*ent.Credential, error) {
	return client.Credential.
		Query().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(ctx)
}

func (*credential_s) Get(team_id int, check_name string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.Get: lock")
		defer mutex.Unlock()
	}

	return Credential.get(team_id, check_name)
}

func (*credential_s) GetComplex(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetComplex: lock")
		defer mutex.Unlock()
	}

	return Credential.get(entTeam.Number, entCheck.Name)
}

func (*credential_s) getWithEdges(team_id int, check_name string) (*ent.Credential, error) {
	return client.Credential.
		Query().
		WithTeam().
		WithCheck().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(ctx)
}

func (*credential_s) GetWithEdges(team_id int, check_name string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetWithEdges: lock")
		defer mutex.Unlock()
	}

	return Credential.getWithEdges(team_id, check_name)
}

func (*credential_s) GetComplexWithEdges(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetComplexWithEdges: lock")
		defer mutex.Unlock()
	}

	return Credential.getWithEdges(entTeam.Number, entCheck.Name)
}

func (*credential_s) getWithCheck(team_id int, check_name string) (*ent.Credential, error) {
	return client.Credential.
		Query().
		WithCheck().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(ctx)
}

func (*credential_s) GetWithCheck(team_id int, check_name string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetWithCheck: lock")
		defer mutex.Unlock()
	}

	return Credential.getWithCheck(team_id, check_name)
}

func (*credential_s) GetComplexWithCheck(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetComplexWithCheck: lock")
		defer mutex.Unlock()
	}

	return Credential.getWithCheck(entTeam.Number, entCheck.Name)
}

func (*credential_s) getWithTeam(team_id int, check_name string) (*ent.Credential, error) {
	return client.Credential.
		Query().
		WithTeam().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(ctx)
}

func (*credential_s) GetWithTeam(team_id int, check_name string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetWithTeam: lock")
		defer mutex.Unlock()
	}

	return Credential.getWithTeam(team_id, check_name)
}

func (*credential_s) GetComplexWithTeam(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.GetComplexWithTeam: lock")
		defer mutex.Unlock()
	}

	return Credential.getWithTeam(entTeam.Number, entCheck.Name)
}

func (*credential_s) exists(team_id int, check_name string) (bool, error) {
	return client.Credential.
		Query().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).
		Exist(ctx)
}

func (*credential_s) Exists(team_id int, check_name string) (bool, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.Exists: lock")
		defer mutex.Unlock()
	}

	return Credential.exists(team_id, check_name)
}

func (*credential_s) ExistsComplex(entTeam *ent.Team, entCheck *ent.Check) (bool, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.ExistsComplex: lock")
		defer mutex.Unlock()
	}

	return Credential.exists(entTeam.Number, entCheck.Name)
}

func (*credential_s) create(team_id int, check_name string, password string) (*ent.Credential, error) {
	entTeam, err := Team.get(team_id)
	if err != nil {
		return nil, err
	}

	exists, err := Credential.exists(team_id, check_name)
	if err != nil {
		return nil, err
	}

	if exists {
		return Credential.get(team_id, check_name)
	}

	entCheck, err := Check.get(check_name)
	if err != nil {
		return nil, err
	}

	return client.Credential.
		Create().
		SetTeam(entTeam).
		SetCheck(entCheck).
		SetPassword(password).
		Save(ctx)
}

func (*credential_s) Create(team_id int, check_name string, password string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.Create: lock")
		defer mutex.Unlock()
	}

	return Credential.create(team_id, check_name, password)
}

func (*credential_s) CreateComplex(entTeam *ent.Team, entCheck *ent.Check, password string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.CreateComplex: lock")
		defer mutex.Unlock()
	}

	return Credential.create(entTeam.Number, entCheck.Name, password)
}

func (*credential_s) update(entCredential *ent.Credential, password string) (*ent.Credential, error) {
	return entCredential.Update().
		SetPassword(password).
		Save(ctx)
}

func (*credential_s) UpdatePassword(team_id int, check_name string, password string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.UpdatePassword: lock")
		defer mutex.Unlock()
	}

	entCredential, err := Credential.get(team_id, check_name)
	if err != nil {
		return nil, err
	}

	return Credential.update(entCredential, password)
}

func (*credential_s) Update(entCredential *ent.Credential, password string) (*ent.Credential, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.Update: lock")
		defer mutex.Unlock()
	}

	return Credential.update(entCredential, password)
}

func (*credential_s) delete(entCredential *ent.Credential) error {
	return client.Credential.
		DeleteOne(entCredential).
		Exec(ctx)
}

func (*credential_s) Delete(team_id int, check_name string) error {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.Delete: lock")
		defer mutex.Unlock()
	}

	entCredential, err := Credential.get(team_id, check_name)
	if err != nil {
		return err
	}

	return Credential.delete(entCredential)
}

func (*credential_s) DeleteComplex(entTeam *ent.Team, entCheck *ent.Check) error {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("credential_s.DeleteComplex: lock")
		defer mutex.Unlock()
	}

	entCredential, err := Credential.get(entTeam.Number, entCheck.Name)
	if err != nil {
		return err
	}

	return Credential.delete(entCredential)
}
