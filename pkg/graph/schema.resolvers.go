package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"time"

	"github.com/compscore/compscore/pkg/ent"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/score"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/user"
	"github.com/google/uuid"
)

// ID is the resolver for the id field.
func (r *checkResolver) ID(ctx context.Context, obj *ent.Check) (string, error) {
	return obj.ID.String(), nil
}

// Credentials is the resolver for the credentials field.
func (r *checkResolver) Credentials(ctx context.Context, obj *ent.Check) ([]*ent.Credential, error) {
	return obj.QueryCredential().All(ctx)
}

// Statuses is the resolver for the statuses field.
func (r *checkResolver) Statuses(ctx context.Context, obj *ent.Check) ([]*ent.Status, error) {
	return obj.QueryStatuses().All(ctx)
}

// ID is the resolver for the id field.
func (r *credentialResolver) ID(ctx context.Context, obj *ent.Credential) (string, error) {
	return obj.ID.String(), nil
}

// User is the resolver for the user field.
func (r *credentialResolver) User(ctx context.Context, obj *ent.Credential) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// CreateCheck is the resolver for the createCheck field.
func (r *mutationResolver) CreateCheck(ctx context.Context, name string, weight int) (*ent.Check, error) {
	return r.Ent.Check.Create().
		SetName(name).
		SetWeight(weight).
		Save(ctx)
}

// UpdateCheck is the resolver for the updateCheck field.
func (r *mutationResolver) UpdateCheck(ctx context.Context, id string, name string, weight int) (*ent.Check, error) {
	// TODO: make fields optional
	// TODO: expand options

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Check.UpdateOneID(uuid).
		SetName(name).
		SetWeight(weight).
		Save(ctx)
}

// DeleteCheck is the resolver for the deleteCheck field.
func (r *mutationResolver) DeleteCheck(ctx context.Context, id string) (string, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	err = r.Ent.Check.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("successfully deleted check: id=%s", id), nil
}

// UpdateCredential is the resolver for the updateCredential field.
func (r *mutationResolver) UpdateCredential(ctx context.Context, id string, password string) (*ent.Credential, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Credential.UpdateOneID(uuid).
		SetPassword(password).
		Save(ctx)
}

// UpdateRound is the resolver for the updateRound field.
func (r *mutationResolver) UpdateRound(ctx context.Context, id string, number int, completed bool) (*ent.Round, error) {
	// TODO: make fields optional

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Round.UpdateOneID(uuid).
		SetNumber(number).
		SetCompleted(completed).
		Save(ctx)
}

// DeleteRound is the resolver for the deleteRound field.
func (r *mutationResolver) DeleteRound(ctx context.Context, id string) (string, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	err = r.Ent.Round.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("successfully deleted round: id=%s", id), nil
}

// UpdateScore is the resolver for the updateScore field.
func (r *mutationResolver) UpdateScore(ctx context.Context, id string) (*ent.Score, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	entScore, err := r.Ent.Score.Query().
		WithRound().
		WithUser().
		Where(
			score.IDEQ(uuid),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	newScore, err := r.Ent.Status.Query().
		Where(
			status.HasRoundWith(
				round.NumberLTE(
					entScore.Edges.Round.Number,
				),
			),
			status.HasUserWith(
				user.TeamNumberEQ(
					entScore.Edges.User.TeamNumber,
				),
			),
			status.StatusEQ(
				status.StatusSuccess,
			),
		).
		Aggregate(
			ent.Sum(
				status.FieldPoints,
			),
		).Int(ctx)
	if err != nil {
		return nil, err
	}

	return r.Ent.Score.UpdateOneID(uuid).
		SetScore(newScore).
		Save(ctx)
}

// DeleteScore is the resolver for the deleteScore field.
func (r *mutationResolver) DeleteScore(ctx context.Context, id string) (string, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	err = r.Ent.Score.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("successfully deleted score: id=%s", id), nil
}

// UpdateStatus is the resolver for the updateStatus field.
func (r *mutationResolver) UpdateStatus(ctx context.Context, id string, status status.Status, message string, timestamp string, points int) (*ent.Status, error) {
	// TODO: make fields optional

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	time, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return nil, err
	}

	return r.Ent.Status.UpdateOneID(uuid).
		SetStatus(status).
		SetMessage(message).
		SetTimestamp(time).
		SetPoints(points).
		Save(ctx)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, teamNumber int, role user.Role) (*ent.User, error) {
	// TODO: add password
	// TODO: hash password

	return r.Ent.User.Create().
		SetName(name).
		SetTeamNumber(teamNumber).
		SetRole(role).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name string, teamNumber int, role user.Role) (*ent.User, error) {
	// TODO: add password
	// TODO: hash password
	// TODO: make fields optional

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.User.UpdateOneID(uuid).
		SetName(name).
		SetTeamNumber(teamNumber).
		SetRole(role).
		Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	err = r.Ent.User.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("successfully deleted user: id=%s", id), nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, name string, password string) (string, error) {
	entUser, err := r.Ent.User.Query().
		Where(
			user.NameEQ(name),
		).
		Only(ctx)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	// TODO: hash password

	if entUser.Password != password {
		return "", fmt.Errorf("invalid username or password")
	}

	// TODO: generate JWT token

	return entUser.ID.String(), nil
}

// Checks is the resolver for the checks field.
func (r *queryResolver) Checks(ctx context.Context) ([]*ent.Check, error) {
	return r.Ent.Check.Query().All(ctx)
}

// Check is the resolver for the check field.
func (r *queryResolver) Check(ctx context.Context, id string) (*ent.Check, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Check.Get(ctx, uuid)
}

// Credentials is the resolver for the credentials field.
func (r *queryResolver) Credentials(ctx context.Context) ([]*ent.Credential, error) {
	return r.Ent.Credential.Query().All(ctx)
}

// Credential is the resolver for the credential field.
func (r *queryResolver) Credential(ctx context.Context, id string) (*ent.Credential, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Credential.Get(ctx, uuid)
}

// Rounds is the resolver for the rounds field.
func (r *queryResolver) Rounds(ctx context.Context) ([]*ent.Round, error) {
	return r.Ent.Round.Query().All(ctx)
}

// Round is the resolver for the round field.
func (r *queryResolver) Round(ctx context.Context, id string) (*ent.Round, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Round.Get(ctx, uuid)
}

// Scores is the resolver for the scores field.
func (r *queryResolver) Scores(ctx context.Context) ([]*ent.Score, error) {
	return r.Ent.Score.Query().All(ctx)
}

// Score is the resolver for the score field.
func (r *queryResolver) Score(ctx context.Context, id string) (*ent.Score, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Score.Get(ctx, uuid)
}

// Statuses is the resolver for the statuses field.
func (r *queryResolver) Statuses(ctx context.Context) ([]*ent.Status, error) {
	return r.Ent.Status.Query().All(ctx)
}

// Status is the resolver for the status field.
func (r *queryResolver) Status(ctx context.Context, id string) (*ent.Status, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.Status.Get(ctx, uuid)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.Ent.User.Query().All(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.Ent.User.Get(ctx, uuid)
}

// ID is the resolver for the id field.
func (r *roundResolver) ID(ctx context.Context, obj *ent.Round) (string, error) {
	return obj.ID.String(), nil
}

// Statuses is the resolver for the statuses field.
func (r *roundResolver) Statuses(ctx context.Context, obj *ent.Round) ([]*ent.Status, error) {
	return obj.QueryStatuses().All(ctx)
}

// Scores is the resolver for the scores field.
func (r *roundResolver) Scores(ctx context.Context, obj *ent.Round) ([]*ent.Score, error) {
	return obj.QueryScores().All(ctx)
}

// ID is the resolver for the id field.
func (r *scoreResolver) ID(ctx context.Context, obj *ent.Score) (string, error) {
	return obj.ID.String(), nil
}

// Round is the resolver for the round field.
func (r *scoreResolver) Round(ctx context.Context, obj *ent.Score) (*ent.Round, error) {
	return obj.QueryRound().Only(ctx)
}

// User is the resolver for the user field.
func (r *scoreResolver) User(ctx context.Context, obj *ent.Score) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// ID is the resolver for the id field.
func (r *statusResolver) ID(ctx context.Context, obj *ent.Status) (string, error) {
	return obj.ID.String(), nil
}

// Timestamp is the resolver for the timestamp field.
func (r *statusResolver) Timestamp(ctx context.Context, obj *ent.Status) (string, error) {
	return obj.Timestamp.String(), nil
}

// Round is the resolver for the round field.
func (r *statusResolver) Round(ctx context.Context, obj *ent.Status) (*ent.Round, error) {
	return obj.QueryRound().Only(ctx)
}

// Check is the resolver for the check field.
func (r *statusResolver) Check(ctx context.Context, obj *ent.Status) (*ent.Check, error) {
	return obj.QueryCheck().Only(ctx)
}

// User is the resolver for the user field.
func (r *statusResolver) User(ctx context.Context, obj *ent.Status) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Credentials is the resolver for the credentials field.
func (r *userResolver) Credentials(ctx context.Context, obj *ent.User) ([]*ent.Credential, error) {
	return obj.QueryCredentials().All(ctx)
}

// Statuses is the resolver for the statuses field.
func (r *userResolver) Statuses(ctx context.Context, obj *ent.User) ([]*ent.Status, error) {
	return obj.QueryStatuses().All(ctx)
}

// Scores is the resolver for the scores field.
func (r *userResolver) Scores(ctx context.Context, obj *ent.User) ([]*ent.Score, error) {
	return obj.QueryScores().All(ctx)
}

// Check returns CheckResolver implementation.
func (r *Resolver) Check() CheckResolver { return &checkResolver{r} }

// Credential returns CredentialResolver implementation.
func (r *Resolver) Credential() CredentialResolver { return &credentialResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Round returns RoundResolver implementation.
func (r *Resolver) Round() RoundResolver { return &roundResolver{r} }

// Score returns ScoreResolver implementation.
func (r *Resolver) Score() ScoreResolver { return &scoreResolver{r} }

// Status returns StatusResolver implementation.
func (r *Resolver) Status() StatusResolver { return &statusResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type checkResolver struct{ *Resolver }
type credentialResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roundResolver struct{ *Resolver }
type scoreResolver struct{ *Resolver }
type statusResolver struct{ *Resolver }
type userResolver struct{ *Resolver }