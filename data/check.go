package data

import (
	"github.com/compscore/compscore/ent"
	"github.com/compscore/compscore/ent/check"
)

type check_s struct{}

func (c *check_s) Get(name string) (*ent.Check, error) {
	return Client.Check.Query().
		Where(
			check.Name(name),
		).
		Only(Ctx)
}

func (c *check_s) GetWithCheckLogs(name string) (*ent.Check, error) {
	return Client.Check.Query().
		Where(
			check.Name(name),
		).
		WithChecklogs().
		Only(Ctx)
}

func (c *check_s) Create(name string, description string, function string, host string) (*ent.Check, error) {
	exist, err := Client.Check.Query().
		Where(
			check.Name(name),
		).Exist(Ctx)
	if err != nil {
		return nil, err
	}

	if exist {
		return c.Get(name)
	}

	return Client.Check.Create().
		SetName(name).
		SetDescription(description).
		SetFunction(function).
		SetHost(host).
		Save(Ctx)
}

func (c *check_s) Update(name string, description string, function string, host string) (int, error) {
	return Client.Check.Update().
		Where(
			check.Name(name),
		).
		SetDescription(description).
		SetFunction(function).
		SetHost(host).
		Save(Ctx)
}

func (c *check_s) Delete(name string) (int, error) {
	return Client.Check.Delete().
		Where(
			check.Name(name),
		).
		Exec(Ctx)
}

func (c *check_s) UpdateDescription(name string, description string) (int, error) {
	return Client.Check.Update().
		Where(
			check.Name(name),
		).
		SetDescription(description).
		Save(Ctx)
}

func (c *check_s) UpdateFunction(name string, function string) (int, error) {
	return Client.Check.Update().
		Where(
			check.Name(name),
		).
		SetFunction(function).
		Save(Ctx)
}

func (c *check_s) UpdateHost(name string, host string) (int, error) {
	return Client.Check.Update().
		Where(
			check.Name(name),
		).
		SetHost(host).
		Save(Ctx)
}

func (c *check_s) UpdateName(oldName string, newName string) (int, error) {
	return Client.Check.Update().
		Where(
			check.Name(oldName),
		).
		SetName(newName).
		Save(Ctx)
}

func (c *check_s) DeleteAll() (int, error) {
	return Client.Check.Delete().
		Exec(Ctx)
}
