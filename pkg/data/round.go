package data

import (
	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/round"
)

type round_s struct{}

var Round = round_s{}

func (*round_s) Get(number int) (*ent.Round, error) {
	return Client.Round.
		Query().
		Where(
			round.NumberEQ(number),
		).Only(Ctx)
}

func (*round_s) GetWithStatus(number int) (*ent.Round, error) {
	return Client.Round.
		Query().
		WithStatus().
		Where(
			round.NumberEQ(number),
		).Only(Ctx)
}

func (*round_s) GetAll(number int) ([]*ent.Round, error) {
	return Client.Round.
		Query().
		All(Ctx)
}

func (*round_s) GetAllWithStatus(number int) ([]*ent.Round, error) {
	return Client.Round.
		Query().
		WithStatus().
		All(Ctx)
}

func (*round_s) GetLastRound() (*ent.Round, error) {
	return Client.Round.
		Query().
		Order(
			round.ByNumber(
				sql.OrderDesc(),
			),
		).
		First(Ctx)
}

func (*round_s) GetLastRoundWithStatus() (*ent.Round, error) {
	return Client.Round.
		Query().
		WithStatus().
		Order(
			round.ByNumber(
				sql.OrderDesc(),
			),
		).
		First(Ctx)
}

func (*round_s) GetPreviousRound() (*ent.Round, error) {
	return Client.Round.
		Query().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		Offset(1).
		First(Ctx)
}

func (*round_s) GetPreviousRoundWithStatus() (*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		Offset(1).
		First(Ctx)
}

func (*round_s) GetPreviousXRounds(amount int) ([]*ent.Round, error) {
	return Client.Round.
		Query().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		Limit(amount).
		All(Ctx)
}

func (*round_s) GetPreviousXRoundsWithStatus(amount int) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		Limit(amount).
		All(Ctx)
}

func (*round_s) Exists(number int) (bool, error) {
	return Client.Round.
		Query().
		Where(
			round.NumberEQ(number),
		).Exist(Ctx)
}

func (*round_s) Create(number int) (*ent.Round, error) {
	exists, err := Round.Exists(number)
	if err != nil {
		return nil, err
	}

	if exists {
		return Round.Get(number)
	}

	return Client.Round.
		Create().
		SetNumber(number).
		Save(Ctx)
}

func (*round_s) Update(round *ent.Round, number int) (*ent.Round, error) {
	return round.Update().
		SetNumber(number).
		Save(Ctx)
}

func (*round_s) Delete(round *ent.Round) error {
	return Client.Round.
		DeleteOne(round).
		Exec(Ctx)
}
