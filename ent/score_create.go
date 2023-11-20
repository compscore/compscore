// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/ent/round"
	"github.com/compscore/compscore/ent/score"
	"github.com/compscore/compscore/ent/user"
	"github.com/google/uuid"
)

// ScoreCreate is the builder for creating a Score entity.
type ScoreCreate struct {
	config
	mutation *ScoreMutation
	hooks    []Hook
}

// SetScore sets the "score" field.
func (sc *ScoreCreate) SetScore(i int) *ScoreCreate {
	sc.mutation.SetScore(i)
	return sc
}

// SetID sets the "id" field.
func (sc *ScoreCreate) SetID(u uuid.UUID) *ScoreCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *ScoreCreate) SetNillableID(u *uuid.UUID) *ScoreCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetRoundID sets the "round" edge to the Round entity by ID.
func (sc *ScoreCreate) SetRoundID(id uuid.UUID) *ScoreCreate {
	sc.mutation.SetRoundID(id)
	return sc
}

// SetRound sets the "round" edge to the Round entity.
func (sc *ScoreCreate) SetRound(r *Round) *ScoreCreate {
	return sc.SetRoundID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *ScoreCreate) SetUserID(id uuid.UUID) *ScoreCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *ScoreCreate) SetUser(u *User) *ScoreCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the ScoreMutation object of the builder.
func (sc *ScoreCreate) Mutation() *ScoreMutation {
	return sc.mutation
}

// Save creates the Score in the database.
func (sc *ScoreCreate) Save(ctx context.Context) (*Score, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScoreCreate) SaveX(ctx context.Context) *Score {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScoreCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScoreCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScoreCreate) defaults() {
	if _, ok := sc.mutation.ID(); !ok {
		v := score.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScoreCreate) check() error {
	if _, ok := sc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Score.score"`)}
	}
	if v, ok := sc.mutation.Score(); ok {
		if err := score.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Score.score": %w`, err)}
		}
	}
	if _, ok := sc.mutation.RoundID(); !ok {
		return &ValidationError{Name: "round", err: errors.New(`ent: missing required edge "Score.round"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Score.user"`)}
	}
	return nil
}

func (sc *ScoreCreate) sqlSave(ctx context.Context) (*Score, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScoreCreate) createSpec() (*Score, *sqlgraph.CreateSpec) {
	var (
		_node = &Score{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(score.Table, sqlgraph.NewFieldSpec(score.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Score(); ok {
		_spec.SetField(score.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if nodes := sc.mutation.RoundIDs(); len(nodes) > 0 {
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
		_node.round_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScoreCreateBulk is the builder for creating many Score entities in bulk.
type ScoreCreateBulk struct {
	config
	err      error
	builders []*ScoreCreate
}

// Save creates the Score entities in the database.
func (scb *ScoreCreateBulk) Save(ctx context.Context) ([]*Score, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Score, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScoreMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScoreCreateBulk) SaveX(ctx context.Context) []*Score {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScoreCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScoreCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
