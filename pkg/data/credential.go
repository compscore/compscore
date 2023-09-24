package data

import (
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/team"
)

type credential_s struct{}

var Credential = credential_s{}

func (*credential_s) Get(team_id int8, check_name string) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(Ctx)
}

func (*credential_s) GetComplex(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		Where(
			credential.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
			credential.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
		).Only(Ctx)
}

func (*credential_s) GetWithEdges(team_id int8, check_name string) (*ent.Credential, error) {
	return Client.Credential.
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
		).Only(Ctx)
}

func (*credential_s) GetComplexWithEdges(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		WithCheck().
		WithTeam().
		Where(
			credential.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
			credential.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
		).Only(Ctx)
}

func (*credential_s) GetWithCheck(team_id int8, check_name string) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		WithCheck().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(Ctx)
}

func (*credential_s) GetComplexWithCheck(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		WithCheck().
		Where(
			credential.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
			credential.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
		).Only(Ctx)
}

func (*credential_s) GetWithTeam(team_id int8, check_name string) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		WithTeam().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Only(Ctx)
}

func (*credential_s) GetComplexWithTeam(entTeam *ent.Team, entCheck *ent.Check) (*ent.Credential, error) {
	return Client.Credential.
		Query().
		WithTeam().
		Where(
			credential.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
			credential.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
		).Only(Ctx)
}

func (*credential_s) Create(team_id int8, check_name string, password string) (*ent.Credential, error) {
	entTeam, err := Team.Get(team_id)
	if err != nil {
		return nil, err
	}

	entCheck, err := Check.Get(check_name)
	if err != nil {
		return nil, err
	}

	return Client.Credential.
		Create().
		SetTeam(entTeam).
		SetCheck(entCheck).
		SetPassword(password).
		Save(Ctx)
}

func (*credential_s) UpdatePassword(team_id int8, check_name string, password string) (*ent.Credential, error) {
	entCredential, err := Credential.Get(team_id, check_name)
	if err != nil {
		return nil, err
	}

	return entCredential.Update().
		SetPassword(password).
		Save(Ctx)
}

func (*credential_s) Update(entCredential *ent.Credential, password string) (*ent.Credential, error) {
	return entCredential.Update().
		SetPassword(password).
		Save(Ctx)
}

func (*credential_s) Delete(entCredential *ent.Credential) error {
	return Client.Credential.
		DeleteOne(entCredential).
		Exec(Ctx)
}

func (*credential_s) DeleteComplex(entTeam *ent.Team, entCheck *ent.Check) error {
	entCredential, err := Credential.GetComplex(entTeam, entCheck)
	if err != nil {
		return err
	}

	return Client.Credential.
		DeleteOne(entCredential).
		Exec(Ctx)
}

func (*credential_s) Exists(team_id int8, check_name string) (bool, error) {
	return Client.Credential.
		Query().
		Where(
			credential.HasTeamWith(
				team.NumberEQ(team_id),
			),
			credential.HasCheckWith(
				check.NameEQ(check_name),
			),
		).Exist(Ctx)
}

func (*credential_s) ExistsComplex(entTeam *ent.Team, entCheck *ent.Check) (bool, error) {
	return Client.Credential.
		Query().
		Where(
			credential.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
			credential.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
		).Exist(Ctx)
}
