package data

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/sirupsen/logrus"
)

type round_s struct{}

var Round = round_s{}

func (*round_s) get(number int) (*ent.Round, error) {
	return client.Round.
		Query().
		Where(
			round.NumberEQ(number),
		).Only(ctx)
}

func (*round_s) Get(number int) (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Get: lock")
		defer mutex.Unlock()
	}

	return Round.get(number)
}

func (*round_s) getWithStatus(number int) (*ent.Round, error) {
	return client.Round.
		Query().
		WithStatus().
		Where(
			round.NumberEQ(number),
		).Only(ctx)
}

func (*round_s) GetWithStatus(number int) (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetWithStatus: lock")
		defer mutex.Unlock()
	}

	return Round.getWithStatus(number)
}

func (*round_s) getAll() ([]*ent.Round, error) {
	return client.Round.
		Query().
		All(ctx)
}

func (*round_s) GetAll() ([]*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetAll: lock")
		defer mutex.Unlock()
	}

	return Round.getAll()
}

func (*round_s) getAllWithStatus() ([]*ent.Round, error) {
	return client.Round.
		Query().
		WithStatus().
		All(ctx)
}

func (*round_s) GetAllWithStatus() ([]*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetAllWithStatus: lock")
		defer mutex.Unlock()
	}

	return Round.getAllWithStatus()
}

func (*round_s) getLastRound() (*ent.Round, error) {
	return client.Round.
		Query().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		First(ctx)
}

func (*round_s) GetLastRound() (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetLastRound: lock")
		defer mutex.Unlock()
	}

	return Round.getLastRound()
}

func (*round_s) getLastRoundWithStatus() (*ent.Round, error) {
	return client.Round.
		Query().
		WithStatus().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		First(ctx)
}

func (*round_s) GetLastRoundWithStatus() (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetLastRoundWithStatus: lock")
		defer mutex.Unlock()
	}

	return Round.getLastRoundWithStatus()
}

func (*round_s) complete(entRound *ent.Round) (*ent.Round, error) {
	return entRound.Update().
		SetComplete(true).
		Save(ctx)
}

func (*round_s) Complete(number int) (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Complete: lock")
		defer mutex.Unlock()
	}

	entRound, err := Round.get(number)
	if err != nil {
		return nil, err
	}

	return Round.complete(entRound)
}

func (*round_s) CompleteComplex(entRound *ent.Round) (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.CompleteComplex: lock")
		defer mutex.Unlock()
	}

	return Round.complete(entRound)
}

func (*round_s) getLastCompleteRound() (*ent.Round, error) {
	return client.Round.
		Query().
		Where(
			round.CompleteEQ(true),
		).
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		First(ctx)
}

func (*round_s) GetLastCompleteRound() (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetLastCompleteRound: lock")
		defer mutex.Unlock()
	}

	return Round.getLastCompleteRound()
}

func (*round_s) getLastCompleteRoundWithStatus() (*ent.Round, error) {
	return client.Round.
		Query().
		WithStatus().
		Where(
			round.CompleteEQ(true),
		).
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		First(ctx)
}

func (*round_s) GetLastCompleteRoundWithStatus() (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetLastCompleteRoundWithStatus: lock")
		defer mutex.Unlock()
	}

	return Round.getLastCompleteRoundWithStatus()
}

func (*round_s) getPreviousXRounds(amount int) ([]*ent.Round, error) {
	return client.Round.
		Query().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		Limit(amount).
		All(ctx)
}

func (*round_s) GetPreviousXRounds(amount int) ([]*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetPreviousXRounds: lock")
		defer mutex.Unlock()
	}

	return Round.getPreviousXRounds(amount)
}

func (*round_s) getPreviousXRoundsWithStatus(amount int) ([]*ent.Round, error) {
	return client.Round.
		Query().
		WithStatus().
		Order(
			ent.Desc(
				round.FieldNumber,
			),
		).
		Limit(amount).
		All(ctx)
}

func (*round_s) GetPreviousXRoundsWithStatus(amount int) ([]*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.GetPreviousXRoundsWithStatus: lock")
		defer mutex.Unlock()
	}

	return Round.getPreviousXRoundsWithStatus(amount)
}

func (*round_s) exists(number int) (bool, error) {
	return client.Round.
		Query().
		Where(
			round.NumberEQ(number),
		).Exist(ctx)
}

func (*round_s) Exists(number int) (bool, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Exists: lock")
		defer mutex.Unlock()
	}

	return Round.exists(number)
}

func (*round_s) create(number int) (*ent.Round, error) {
	exists, err := Round.exists(number)
	if err != nil {
		return nil, err
	}

	if exists {
		return Round.get(number)
	}

	return client.Round.
		Create().
		SetNumber(number).
		Save(ctx)
}

func (*round_s) Create(number int) (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Create: lock")
		defer mutex.Unlock()
	}

	return Round.create(number)
}

func (*round_s) count() (int, error) {
	return client.Round.
		Query().
		Count(ctx)
}

func (*round_s) Count() (int, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Count: lock")
		defer mutex.Unlock()
	}

	return Round.count()
}

func (*round_s) createNextRound() (*ent.Round, error) {
	roundCount, err := Round.count()
	if err != nil {
		return nil, err
	}

	if roundCount == 0 {
		return Round.create(1)
	}

	lastRound, err := Round.getLastRound()
	if err != nil {
		return nil, err
	}

	return Round.create(lastRound.Number + 1)
}

func (*round_s) CreateNextRound() (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.CreateNextRound: lock")
		defer mutex.Unlock()
	}

	return Round.createNextRound()
}

func (*round_s) delete(round *ent.Round) error {
	return client.Round.
		DeleteOne(round).
		Exec(ctx)
}

func (*round_s) Delete(round *ent.Round) error {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Delete: lock")
		defer mutex.Unlock()
	}

	return Round.delete(round)
}

func (*round_s) update(round *ent.Round, number int, complete bool) (*ent.Round, error) {
	return round.Update().
		SetNumber(number).
		SetComplete(complete).
		Save(ctx)
}

func (*round_s) Update(round *ent.Round, number int, complete bool) (*ent.Round, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.Update: lock")
		defer mutex.Unlock()
	}

	return Round.update(round, number, complete)
}

func (*round_s) deleteAll() (int, error) {
	return client.Round.
		Delete().
		Exec(ctx)
}

func (*round_s) DeleteAll() (int, error) {
	if !config.Production {
		mutex.Lock()
		logrus.Trace("round_s.DeleteAll: lock")
		defer mutex.Unlock()
	}

	return Round.deleteAll()
}
