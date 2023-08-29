package data

import (
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
)

type check_s struct{}

var Check = check_s{}

func (*check_s) Get(name string) (*ent.Check, error) {
	return Client.Check.
		Query().
		Where(
			check.NameEQ(name),
		).Only(Ctx)
}

func (*check_s) GetWithStatus(name string) (*ent.Check, error) {
	return Client.Check.
		Query().
		WithStatus().
		Where(
			check.NameEQ(name),
		).Only(Ctx)
}

func (*check_s) GetAll() ([]*ent.Check, error) {
	return Client.Check.
		Query().
		Order(
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*check_s) GetAllWithStatus() ([]*ent.Check, error) {
	return Client.Check.
		Query().
		WithStatus().
		Order(
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*check_s) Exists(name string) (bool, error) {
	return Client.Check.
		Query().
		Where(
			check.NameEQ(name),
		).
		Exist(Ctx)
}

func (*check_s) Create(name string) (*ent.Check, error) {
	exists, err := Check.Exists(name)
	if err != nil {
		return nil, err
	}

	if exists {
		return Check.Get(name)
	}

	return Client.Check.
		Create().
		SetName(name).
		Save(Ctx)
}

func (*check_s) Update(check *ent.Check) (*ent.Check, error) {
	return check.Update().
		SetName(check.Name).
		Save(Ctx)
}

func (*check_s) Delete(check *ent.Check) error {
	return Client.Check.
		DeleteOne(check).
		Exec(Ctx)
}
