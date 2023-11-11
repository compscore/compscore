// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/status"
)

// CheckCreate is the builder for creating a Check entity.
type CheckCreate struct {
	config
	mutation *CheckMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CheckCreate) SetName(s string) *CheckCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetWeight sets the "weight" field.
func (cc *CheckCreate) SetWeight(i int) *CheckCreate {
	cc.mutation.SetWeight(i)
	return cc
}

// SetID sets the "id" field.
func (cc *CheckCreate) SetID(i int) *CheckCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddCredentialIDs adds the "credential" edge to the Credential entity by IDs.
func (cc *CheckCreate) AddCredentialIDs(ids ...int) *CheckCreate {
	cc.mutation.AddCredentialIDs(ids...)
	return cc
}

// AddCredential adds the "credential" edges to the Credential entity.
func (cc *CheckCreate) AddCredential(c ...*Credential) *CheckCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCredentialIDs(ids...)
}

// AddStatuIDs adds the "status" edge to the Status entity by IDs.
func (cc *CheckCreate) AddStatuIDs(ids ...int) *CheckCreate {
	cc.mutation.AddStatuIDs(ids...)
	return cc
}

// AddStatus adds the "status" edges to the Status entity.
func (cc *CheckCreate) AddStatus(s ...*Status) *CheckCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStatuIDs(ids...)
}

// Mutation returns the CheckMutation object of the builder.
func (cc *CheckCreate) Mutation() *CheckMutation {
	return cc.mutation
}

// Save creates the Check in the database.
func (cc *CheckCreate) Save(ctx context.Context) (*Check, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CheckCreate) SaveX(ctx context.Context) *Check {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CheckCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CheckCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CheckCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Check.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := check.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Check.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Check.weight"`)}
	}
	if v, ok := cc.mutation.Weight(); ok {
		if err := check.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Check.weight": %w`, err)}
		}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := check.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Check.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CheckCreate) sqlSave(ctx context.Context) (*Check, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CheckCreate) createSpec() (*Check, *sqlgraph.CreateSpec) {
	var (
		_node = &Check{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(check.Table, sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Weight(); ok {
		_spec.SetField(check.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if nodes := cc.mutation.CredentialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.CredentialTable,
			Columns: []string{check.CredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   check.StatusTable,
			Columns: []string{check.StatusColumn},
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

// CheckCreateBulk is the builder for creating many Check entities in bulk.
type CheckCreateBulk struct {
	config
	err      error
	builders []*CheckCreate
}

// Save creates the Check entities in the database.
func (ccb *CheckCreateBulk) Save(ctx context.Context) ([]*Check, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Check, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CheckMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CheckCreateBulk) SaveX(ctx context.Context) []*Check {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CheckCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CheckCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}