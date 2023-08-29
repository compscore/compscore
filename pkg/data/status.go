package data

import (
	"time"

	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/team"
)

type status_s struct{}

var Status = status_s{}

func (*status_s) Get(roundNumber int, checkName string, teamNumber int8) (*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Only(Ctx)
}

func (*status_s) GetWithEdges(roundNumber int, checkName string, teamNumber int8) (*ent.Status, error) {
	return Client.Status.
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
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Only(Ctx)
}

func (*status_s) GetComplex(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.IDEQ(entRound.ID),
			),
			status.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
			status.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
		).Only(Ctx)
}

func (*status_s) GetComplexWithEdges(entRound *ent.Round, entCheck *ent.Check, entTeam *ent.Team) (*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Where(
			status.HasRoundWith(
				round.IDEQ(entRound.ID),
			),
			status.HasCheckWith(
				check.IDEQ(entCheck.ID),
			),
			status.HasTeamWith(
				team.IDEQ(entTeam.ID),
			),
		).Only(Ctx)
}

func (*status_s) GetAll() ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllWithEdges() ([]*ent.Status, error) {
	return Client.Status.
		Query().
		WithRound().
		WithCheck().
		WithTeam().
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByRound(roundNumber int) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
		).
		Order(
			ent.Asc(team.FieldNumber),
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundWithEdges(roundNumber int) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(team.FieldNumber),
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheck(checkName string) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheckWithEdges(checkName string) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByTeam(teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
		Query().
		Where(
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).
		Order(
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByTeamWithEdges(teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(round.FieldNumber),
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndCheck(roundNumber int, checkName string) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndCheckWithEdges(roundNumber int, checkName string) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndTeam(roundNumber int, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByRoundAndTeamWithEdges(roundNumber int, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(check.FieldName),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheckAndTeam(checkName string, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(round.FieldNumber),
			ent.Asc(team.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) GetAllByCheckAndTeamWithEdges(checkName string, teamNumber int8) ([]*ent.Status, error) {
	return Client.Status.
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
			ent.Asc(round.FieldNumber),
		).
		All(Ctx)
}

func (*status_s) Exists(roundNumber int, checkName string, teamNumber int8) (bool, error) {
	return Client.Status.
		Query().
		Where(
			status.HasRoundWith(
				round.NumberEQ(roundNumber),
			),
			status.HasCheckWith(
				check.NameEQ(checkName),
			),
			status.HasTeamWith(
				team.NumberEQ(teamNumber),
			),
		).Exist(Ctx)
}

func (*status_s) Create(roundNumber int, checkName string, teamNumber int8, status string) (*ent.Status, error) {
	exist, err := Status.Exists(roundNumber, checkName, teamNumber)
	if err != nil {
		return nil, err
	}

	if exist {
		return Status.Get(roundNumber, checkName, teamNumber)
	}

	entRound, err := Round.Get(roundNumber)
	if err != nil {
		return nil, err
	}

	entCheck, err := Check.Get(checkName)
	if err != nil {
		return nil, err
	}

	entTeam, err := Team.Get(teamNumber)
	if err != nil {
		return nil, err
	}

	return Client.Status.
		Create().
		SetRound(entRound).
		SetCheck(entCheck).
		SetTeam(entTeam).
		Save(Ctx)
}

func (*status_s) Update(entStatus *ent.Status, statusEnum status.Status, error string) (*ent.Status, error) {
	return entStatus.Update().
		SetStatus(statusEnum).
		SetTime(time.Now()).
		Save(Ctx)
}

func (*status_s) Delete(entStatus *ent.Status) error {
	return Client.Status.
		DeleteOne(entStatus).
		Exec(Ctx)
}
