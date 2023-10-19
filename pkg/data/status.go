package data

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/team"
	"github.com/sirupsen/logrus"
)

type status_s struct{}

var Status = status_s{}

func (*status_s) exists(roundNumber int, checkNumber string, teamNumber int) (bool, error) {
	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Exist(ctx)
}

func (*status_s) Exists(roundNumber int, checkNumber string, teamNumber int) (bool, error) {
	mutex.Lock()
	logrus.Trace("status_s.Exists: lock")
	defer mutex.Unlock()

	return Status.exists(roundNumber, checkNumber, teamNumber)
}

func (*status_s) create(roundNumber int, checkNumber string, teamNumber int, status status.Status) (*ent.Status, error) {
	entRound, err := Round.get(roundNumber)
	if err != nil {
		return nil, err
	}

	entCheck, err := Check.get(checkNumber)
	if err != nil {
		return nil, err
	}

	entTeam, err := Team.get(teamNumber)
	if err != nil {
		return nil, err
	}

	return client.Status.
		Create().
		SetRound(
			entRound,
		).
		SetCheck(
			entCheck,
		).
		SetTeam(
			entTeam,
		).
		SetStatus(
			status,
		).
		Save(ctx)
}

func (*status_s) Create(roundNumber int, checkNumber string, teamNumber int, status status.Status) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.Create: lock")
	defer mutex.Unlock()

	return Status.create(roundNumber, checkNumber, teamNumber, status)
}

func (*status_s) createComplex(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	return client.Status.
		Create().
		SetRound(
			entRound,
		).
		SetCheck(
			entCheck,
		).
		SetTeam(
			entTeam,
		).
		Save(ctx)
}

func (*status_s) CreateComplex(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.CreateComplex: lock")
	defer mutex.Unlock()

	return Status.createComplex(entRound, entCheck, entTeam)
}

func (*status_s) get(roundNumber int, checkNumber string, teamNumber int) (*ent.Status, error) {
	exist, err := Status.exists(roundNumber, checkNumber, teamNumber)
	if err != nil {
		return nil, err
	}

	if !exist {
		return Status.create(roundNumber, checkNumber, teamNumber, status.StatusUnknown)
	}

	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Only(ctx)
}

func (*status_s) Get(roundNumber int, checkNumber string, teamNumber int) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.Get: lock")
	defer mutex.Unlock()

	return Status.get(roundNumber, checkNumber, teamNumber)
}

func (*status_s) getWithEdges(roundNumber int, checkNumber string, teamNumber int) (*ent.Status, error) {
	exist, err := Status.exists(roundNumber, checkNumber, teamNumber)
	if err != nil {
		return nil, err
	}

	if !exist {
		return Status.create(roundNumber, checkNumber, teamNumber, status.StatusUnknown)
	}

	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		WithRound().
		WithCheck().
		WithTeam().
		Only(ctx)
}

func (*status_s) GetWithEdges(roundNumber int, checkNumber string, teamNumber int) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetWithEdges: lock")
	defer mutex.Unlock()

	return Status.getWithEdges(roundNumber, checkNumber, teamNumber)
}

func (*status_s) getComplex(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	exists, err := Status.exists(entRound.Number, entCheck.Name, entTeam.Number)
	if err != nil {
		return nil, err
	}

	if !exists {
		return Status.createComplex(entRound, entCheck, entTeam)
	}

	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(entRound.Number),
			),
			status.HasCheckWith(
				check.NameEQ(entCheck.Name),
			),
			status.HasTeamWith(
				team.NumberEQ(entTeam.Number),
			),
		).
		Only(ctx)
}

func (*status_s) GetComplex(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetComplex: lock")
	defer mutex.Unlock()

	return Status.getComplex(entRound, entCheck, entTeam)
}

func (*status_s) getComplexWithEdges(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	exists, err := Status.exists(entRound.Number, entCheck.Name, entTeam.Number)
	if err != nil {
		return nil, err
	}

	if !exists {
		return Status.createComplex(entRound, entCheck, entTeam)
	}

	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(entRound.Number),
			),
			status.HasCheckWith(
				check.NameEQ(entCheck.Name),
			),
			status.HasTeamWith(
				team.NumberEQ(entTeam.Number),
			),
		).
		WithRound().
		WithCheck().
		WithTeam().
		Only(ctx)
}

func (*status_s) GetComplexWithEdges(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetComplexWithEdges: lock")
	defer mutex.Unlock()

	return Status.getComplexWithEdges(entRound, entCheck, entTeam)
}

func (*status_s) getAll() ([]*ent.Status, error) {
	return client.Status.
		Query().
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByTeamField(
				team.FieldNumber,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAll() ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAll: lock")
	defer mutex.Unlock()

	return Status.getAll()
}

func (*status_s) getAllWithEdges() ([]*ent.Status, error) {
	return client.Status.
		Query().
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByTeamField(
				team.FieldNumber,
				sql.OrderAsc(),
			),
		).
		WithRound().
		WithCheck().
		WithTeam().
		All(ctx)
}

func (*status_s) GetAllWithEdges() ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllWithEdges()
}

func (*status_s) getAllByRound(roundNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
		).
		Order(
			status.ByTeamField(
				team.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByCheckField(
				check.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByRound(roundNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByRound: lock")
	defer mutex.Unlock()

	return Status.getAllByRound(roundNumber)
}

func (*status_s) getAllByRoundWithEdges(roundNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
		).
		Order(
			status.ByTeamField(
				team.FieldName,
				sql.OrderAsc(),
			),
			status.ByCheckField(
				check.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByRoundWithEdges(roundNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByRoundWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllByRoundWithEdges(roundNumber)
}

func (*status_s) getAllByCheck(checkName string) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByTeamField(
				team.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByCheck(checkName string) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheck: lock")
	defer mutex.Unlock()

	return Status.getAllByCheck(checkName)
}

func (*status_s) getAllByCheckWithEdges(checkName string) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByTeamField(
				team.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByCheckWithEdges(checkName string) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheckWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllByCheckWithEdges(checkName)
}

func (*status_s) getAllByTeam(teamNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByTeamField(
				team.FieldNumber,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByTeam(teamNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByTeam: lock")
	defer mutex.Unlock()

	return Status.getAllByTeam(teamNumber)
}

func (*status_s) getAllByTeamWithEdges(teamNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
			status.ByCheckField(
				check.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByTeamWithEdges(teamNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByTeamWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllByTeamWithEdges(teamNumber)
}

func (*status_s) getAllByRoundAndCheck(roundNumber int, checkName string) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			status.ByTeamField(
				team.FieldNumber,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByRoundAndCheck(roundNumber int, checkName string) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByRoundAndCheck: lock")
	defer mutex.Unlock()

	return Status.getAllByRoundAndCheck(roundNumber, checkName)
}

func (*status_s) getAllByRoundAndCheckWithEdges(roundNumber int, checkName string) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			status.ByTeamField(
				team.FieldNumber,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByRoundAndCheckWithEdges(roundNumber int, checkName string) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByRoundAndCheckWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllByRoundAndCheckWithEdges(roundNumber, checkName)
}

func (*status_s) getAllByRoundAndTeam(roundNumber int, teamNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByCheckField(
				check.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByRoundAndTeam(roundNumber int, teamNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByRoundAndTeam: lock")
	defer mutex.Unlock()

	return Status.getAllByRoundAndTeam(roundNumber, teamNumber)
}

func (*status_s) getAllByRoundAndTeamWithEdges(roundNumber int, teamNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByCheckField(
				check.FieldName,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByRoundAndTeamWithEdges(roundNumber int, teamNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByRoundAndTeamWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllByRoundAndTeamWithEdges(roundNumber, teamNumber)
}

func (*status_s) getAllByCheckAndTeam(checkName string, teamNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
		).
		All(ctx)
}

func (*status_s) GetAllByCheckAndTeam(checkName string, teamNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheckAndTeam: lock")
	defer mutex.Unlock()

	return Status.getAllByCheckAndTeam(checkName, teamNumber)
}

func (*status_s) getAllByCheckAndTeamWithLimit(checkName string, teamNumber int, limit int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderAsc(),
			),
		).
		Limit(limit).
		All(ctx)
}

func (*status_s) GetAllByCheckAndTeamWithLimit(checkName string, teamNumber int, limit int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheckAndTeamWithLimit: lock")
	defer mutex.Unlock()

	return Status.getAllByCheckAndTeamWithLimit(checkName, teamNumber, limit)
}

func (*status_s) getAllByCheckAndTeamWithEdges(checkName string, teamNumber int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderDesc(),
			)).
		All(ctx)
}

func (*status_s) GetAllByCheckAndTeamWithEdges(checkName string, teamNumber int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheckAndTeamWithEdges: lock")
	defer mutex.Unlock()

	return Status.getAllByCheckAndTeamWithEdges(checkName, teamNumber)
}

func (*status_s) getAllByCheckAndTeamWithEdgesWithLimit(checkName string, teamNumber int, limit int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderDesc(),
			)).
		Limit(limit).
		All(ctx)
}

func (*status_s) GetAllByCheckAndTeamWithEdgesWithLimit(checkName string, teamNumber int, limit int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheckAndTeamWithEdgesWithLimit: lock")
	defer mutex.Unlock()

	return Status.getAllByCheckAndTeamWithEdgesWithLimit(checkName, teamNumber, limit)
}

func (*status_s) getAllByCheckAndTeamWithEdgesFromRoundWithLimit(checkName string, teamNumber int, roundNumber int, limit int) ([]*ent.Status, error) {
	return client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
			status.HasRoundWith(
				round.NumberLTE(roundNumber),
			),
		).
		Order(
			status.ByRoundField(
				round.FieldNumber,
				sql.OrderDesc(),
			)).
		Limit(limit).
		All(ctx)
}

func (*status_s) GetAllByCheckAndTeamWithEdgesFromRoundWithLimit(checkName string, teamNumber int, roundNumber int, limit int) ([]*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.GetAllByCheckAndTeamWithEdgesWithLimit: lock")
	defer mutex.Unlock()

	return Status.getAllByCheckAndTeamWithEdgesFromRoundWithLimit(checkName, teamNumber, roundNumber, limit)
}

func (*status_s) update(teamNumber int, roundNumber int, checkName string, statusEnum status.Status, message string) (int, error) {
	return client.Status.Update().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		SetStatus(statusEnum).
		SetError(message).
		SetTime(time.Now()).
		Save(ctx)
}

func (*status_s) Update(teamNumber int, roundNumber int, checkName string, statusEnum status.Status, message string) (int, error) {
	mutex.Lock()
	logrus.Trace("status_s.Update: lock")
	defer mutex.Unlock()

	return Status.update(teamNumber, roundNumber, checkName, statusEnum, message)
}

func (*status_s) updateComplex(entStatus *ent.Status, statusEnum status.Status, message string) (*ent.Status, error) {
	return entStatus.Update().
		SetStatus(statusEnum).
		SetError(message).
		SetTime(time.Now()).
		Save(ctx)
}

func (*status_s) UpdateComplex(entStatus *ent.Status, statusEnum status.Status, message string) (*ent.Status, error) {
	mutex.Lock()
	logrus.Trace("status_s.UpdateComplex: lock")
	defer mutex.Unlock()

	return Status.updateComplex(entStatus, statusEnum, message)
}

func (*status_s) delete(teamNumber int, roundNumber int, checkName string) (int, error) {
	return client.Status.
		Delete().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Exec(ctx)
}

func (*status_s) Delete(teamNumber int, roundNumber int, checkName string) (int, error) {
	mutex.Lock()
	logrus.Trace("status_s.Delete: lock")
	defer mutex.Unlock()

	return Status.delete(teamNumber, roundNumber, checkName)
}

func (*status_s) deleteComplex(entStatus *ent.Status) error {
	return client.Status.
		DeleteOne(entStatus).
		Exec(ctx)
}

func (*status_s) DeleteComplex(entStatus *ent.Status) error {
	mutex.Lock()
	logrus.Trace("status_s.DeleteComplex: lock")
	defer mutex.Unlock()

	return Status.deleteComplex(entStatus)
}

func (*status_s) deleteAll() (int, error) {
	return client.Status.
		Delete().
		Exec(ctx)
}

func (*status_s) DeleteAll() (int, error) {
	mutex.Lock()
	logrus.Trace("status_s.DeleteAll: lock")
	defer mutex.Unlock()

	return Status.deleteAll()
}
