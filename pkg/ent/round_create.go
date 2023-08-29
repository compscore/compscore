// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
)

// RoundCreate is the builder for creating a Round entity.
type RoundCreate struct {
	config
	mutation *RoundMutation
	hooks    []Hook
}

// SetNumber sets the "number" field.
func (rc *RoundCreate) SetNumber(i int) *RoundCreate {
	rc.mutation.SetNumber(i)
	return rc
}

// AddStatuIDs adds the "status" edge to the Status entity by IDs.
func (rc *RoundCreate) AddStatuIDs(ids ...int) *RoundCreate {
	rc.mutation.AddStatuIDs(ids...)
	return rc
}

// AddStatus adds the "status" edges to the Status entity.
func (rc *RoundCreate) AddStatus(s ...*Status) *RoundCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddStatuIDs(ids...)
}

// Mutation returns the RoundMutation object of the builder.
func (rc *RoundCreate) Mutation() *RoundMutation {
	return rc.mutation
}

// Save creates the Round in the database.
func (rc *RoundCreate) Save(ctx context.Context) (*Round, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoundCreate) SaveX(ctx context.Context) *Round {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoundCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoundCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoundCreate) check() error {
	if _, ok := rc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Round.number"`)}
	}
	if v, ok := rc.mutation.Number(); ok {
		if err := round.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Round.number": %w`, err)}
		}
	}
	return nil
}

func (rc *RoundCreate) sqlSave(ctx context.Context) (*Round, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoundCreate) createSpec() (*Round, *sqlgraph.CreateSpec) {
	var (
		_node = &Round{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(round.Table, sqlgraph.NewFieldSpec(round.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Number(); ok {
		_spec.SetField(round.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if nodes := rc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   round.StatusTable,
			Columns: []string{round.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoundCreateBulk is the builder for creating many Round entities in bulk.
type RoundCreateBulk struct {
	config
	builders []*RoundCreate
}

// Save creates the Round entities in the database.
func (rcb *RoundCreateBulk) Save(ctx context.Context) ([]*Round, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Round, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoundMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoundCreateBulk) SaveX(ctx context.Context) []*Round {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoundCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoundCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
