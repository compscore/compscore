package data

import (
	"github.com/compscore/compscore/ent"
	"github.com/compscore/compscore/ent/round"
)

type round_s struct{}

func (r *round_s) Get(number int) (*ent.Round, error) {
	return Client.Round.Query().
		Where(
			round.Number(number),
		).Only(Ctx)
}

func (r *round_s) GetWithCheckLogs(number int) (*ent.Round, error) {
	return Client.Round.Query().
		Where(
			round.Number(number),
		).
		WithChecklogs().
		Only(Ctx)
}

func (r *round_s) GetLast() (*ent.Round, error) {
	return Client.Round.Query().
		Order(ent.Desc(round.FieldNumber)).
		First(Ctx)
}

func (r *round_s) GetRounds() (int, error) {
	return Client.Round.Query().
		Count(Ctx)
}

func (r *round_s) Create() (*ent.Round, error) {
	lastRoundNumber, err := Round.GetRounds()
	if err != nil {
		return nil, err
	}

	return Client.Round.Create().
		SetNumber(lastRoundNumber + 1).
		Save(Ctx)
}

func (r *round_s) Delete(number int) (int, error) {
	return Client.Round.Delete().
		Where(
			round.Number(number),
		).Exec(Ctx)
}

func (r *round_s) DeleteAll() (int, error) {
	return Client.Round.Delete().
		Exec(Ctx)
}
