// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/ent/predicate"
	"github.com/compscore/compscore/ent/round"
	"github.com/compscore/compscore/ent/score"
	"github.com/compscore/compscore/ent/user"
	"github.com/google/uuid"
)

// ScoreUpdate is the builder for updating Score entities.
type ScoreUpdate struct {
	config
	hooks    []Hook
	mutation *ScoreMutation
}

// Where appends a list predicates to the ScoreUpdate builder.
func (su *ScoreUpdate) Where(ps ...predicate.Score) *ScoreUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetScore sets the "score" field.
func (su *ScoreUpdate) SetScore(i int) *ScoreUpdate {
	su.mutation.ResetScore()
	su.mutation.SetScore(i)
	return su
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (su *ScoreUpdate) SetNillableScore(i *int) *ScoreUpdate {
	if i != nil {
		su.SetScore(*i)
	}
	return su
}

// AddScore adds i to the "score" field.
func (su *ScoreUpdate) AddScore(i int) *ScoreUpdate {
	su.mutation.AddScore(i)
	return su
}

// SetRoundID sets the "round" edge to the Round entity by ID.
func (su *ScoreUpdate) SetRoundID(id uuid.UUID) *ScoreUpdate {
	su.mutation.SetRoundID(id)
	return su
}

// SetRound sets the "round" edge to the Round entity.
func (su *ScoreUpdate) SetRound(r *Round) *ScoreUpdate {
	return su.SetRoundID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *ScoreUpdate) SetUserID(id uuid.UUID) *ScoreUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *ScoreUpdate) SetUser(u *User) *ScoreUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the ScoreMutation object of the builder.
func (su *ScoreUpdate) Mutation() *ScoreMutation {
	return su.mutation
}

// ClearRound clears the "round" edge to the Round entity.
func (su *ScoreUpdate) ClearRound() *ScoreUpdate {
	su.mutation.ClearRound()
	return su
}

// ClearUser clears the "user" edge to the User entity.
func (su *ScoreUpdate) ClearUser() *ScoreUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScoreUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScoreUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScoreUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ScoreUpdate) check() error {
	if v, ok := su.mutation.Score(); ok {
		if err := score.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Score.score": %w`, err)}
		}
	}
	if _, ok := su.mutation.RoundID(); su.mutation.RoundCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Score.round"`)
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Score.user"`)
	}
	return nil
}

func (su *ScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(score.Table, score.Columns, sqlgraph.NewFieldSpec(score.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Score(); ok {
		_spec.SetField(score.FieldScore, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedScore(); ok {
		_spec.AddField(score.FieldScore, field.TypeInt, value)
	}
	if su.mutation.RoundCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.RoundTable,
			Columns: []string{score.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.RoundTable,
			Columns: []string{score.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.UserTable,
			Columns: []string{score.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.UserTable,
			Columns: []string{score.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{score.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ScoreUpdateOne is the builder for updating a single Score entity.
type ScoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScoreMutation
}

// SetScore sets the "score" field.
func (suo *ScoreUpdateOne) SetScore(i int) *ScoreUpdateOne {
	suo.mutation.ResetScore()
	suo.mutation.SetScore(i)
	return suo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (suo *ScoreUpdateOne) SetNillableScore(i *int) *ScoreUpdateOne {
	if i != nil {
		suo.SetScore(*i)
	}
	return suo
}

// AddScore adds i to the "score" field.
func (suo *ScoreUpdateOne) AddScore(i int) *ScoreUpdateOne {
	suo.mutation.AddScore(i)
	return suo
}

// SetRoundID sets the "round" edge to the Round entity by ID.
func (suo *ScoreUpdateOne) SetRoundID(id uuid.UUID) *ScoreUpdateOne {
	suo.mutation.SetRoundID(id)
	return suo
}

// SetRound sets the "round" edge to the Round entity.
func (suo *ScoreUpdateOne) SetRound(r *Round) *ScoreUpdateOne {
	return suo.SetRoundID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *ScoreUpdateOne) SetUserID(id uuid.UUID) *ScoreUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *ScoreUpdateOne) SetUser(u *User) *ScoreUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the ScoreMutation object of the builder.
func (suo *ScoreUpdateOne) Mutation() *ScoreMutation {
	return suo.mutation
}

// ClearRound clears the "round" edge to the Round entity.
func (suo *ScoreUpdateOne) ClearRound() *ScoreUpdateOne {
	suo.mutation.ClearRound()
	return suo
}

// ClearUser clears the "user" edge to the User entity.
func (suo *ScoreUpdateOne) ClearUser() *ScoreUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the ScoreUpdate builder.
func (suo *ScoreUpdateOne) Where(ps ...predicate.Score) *ScoreUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScoreUpdateOne) Select(field string, fields ...string) *ScoreUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Score entity.
func (suo *ScoreUpdateOne) Save(ctx context.Context) (*Score, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScoreUpdateOne) SaveX(ctx context.Context) *Score {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScoreUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScoreUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ScoreUpdateOne) check() error {
	if v, ok := suo.mutation.Score(); ok {
		if err := score.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Score.score": %w`, err)}
		}
	}
	if _, ok := suo.mutation.RoundID(); suo.mutation.RoundCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Score.round"`)
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Score.user"`)
	}
	return nil
}

func (suo *ScoreUpdateOne) sqlSave(ctx context.Context) (_node *Score, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(score.Table, score.Columns, sqlgraph.NewFieldSpec(score.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Score.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, score.FieldID)
		for _, f := range fields {
			if !score.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != score.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Score(); ok {
		_spec.SetField(score.FieldScore, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedScore(); ok {
		_spec.AddField(score.FieldScore, field.TypeInt, value)
	}
	if suo.mutation.RoundCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.RoundTable,
			Columns: []string{score.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.RoundTable,
			Columns: []string{score.RoundColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.UserTable,
			Columns: []string{score.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.UserTable,
			Columns: []string{score.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Score{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{score.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
