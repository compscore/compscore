// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/predicate"
	"github.com/compscore/compscore/pkg/ent/status"
)

// CheckUpdate is the builder for updating Check entities.
type CheckUpdate struct {
	config
	hooks    []Hook
	mutation *CheckMutation
}

// Where appends a list predicates to the CheckUpdate builder.
func (cu *CheckUpdate) Where(ps ...predicate.Check) *CheckUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CheckUpdate) SetName(s string) *CheckUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetWeight sets the "weight" field.
func (cu *CheckUpdate) SetWeight(i int) *CheckUpdate {
	cu.mutation.ResetWeight()
	cu.mutation.SetWeight(i)
	return cu
}

// AddWeight adds i to the "weight" field.
func (cu *CheckUpdate) AddWeight(i int) *CheckUpdate {
	cu.mutation.AddWeight(i)
	return cu
}

// AddStatuIDs adds the "status" edge to the Status entity by IDs.
func (cu *CheckUpdate) AddStatuIDs(ids ...int) *CheckUpdate {
	cu.mutation.AddStatuIDs(ids...)
	return cu
}

// AddStatus adds the "status" edges to the Status entity.
func (cu *CheckUpdate) AddStatus(s ...*Status) *CheckUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddStatuIDs(ids...)
}

// AddCredentialIDs adds the "credential" edge to the Credential entity by IDs.
func (cu *CheckUpdate) AddCredentialIDs(ids ...int) *CheckUpdate {
	cu.mutation.AddCredentialIDs(ids...)
	return cu
}

// AddCredential adds the "credential" edges to the Credential entity.
func (cu *CheckUpdate) AddCredential(c ...*Credential) *CheckUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCredentialIDs(ids...)
}

// Mutation returns the CheckMutation object of the builder.
func (cu *CheckUpdate) Mutation() *CheckMutation {
	return cu.mutation
}

// ClearStatus clears all "status" edges to the Status entity.
func (cu *CheckUpdate) ClearStatus() *CheckUpdate {
	cu.mutation.ClearStatus()
	return cu
}

// RemoveStatuIDs removes the "status" edge to Status entities by IDs.
func (cu *CheckUpdate) RemoveStatuIDs(ids ...int) *CheckUpdate {
	cu.mutation.RemoveStatuIDs(ids...)
	return cu
}

// RemoveStatus removes "status" edges to Status entities.
func (cu *CheckUpdate) RemoveStatus(s ...*Status) *CheckUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveStatuIDs(ids...)
}

// ClearCredential clears all "credential" edges to the Credential entity.
func (cu *CheckUpdate) ClearCredential() *CheckUpdate {
	cu.mutation.ClearCredential()
	return cu
}

// RemoveCredentialIDs removes the "credential" edge to Credential entities by IDs.
func (cu *CheckUpdate) RemoveCredentialIDs(ids ...int) *CheckUpdate {
	cu.mutation.RemoveCredentialIDs(ids...)
	return cu
}

// RemoveCredential removes "credential" edges to Credential entities.
func (cu *CheckUpdate) RemoveCredential(c ...*Credential) *CheckUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCredentialIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CheckUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CheckUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CheckUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CheckUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CheckUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := check.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Check.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Weight(); ok {
		if err := check.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Check.weight": %w`, err)}
		}
	}
	return nil
}

func (cu *CheckUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Weight(); ok {
		_spec.SetField(check.FieldWeight, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedWeight(); ok {
		_spec.AddField(check.FieldWeight, field.TypeInt, value)
	}
	if cu.mutation.StatusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStatusIDs(); len(nodes) > 0 && !cu.mutation.StatusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StatusIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CredentialCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCredentialIDs(); len(nodes) > 0 && !cu.mutation.CredentialCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CredentialIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CheckUpdateOne is the builder for updating a single Check entity.
type CheckUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckMutation
}

// SetName sets the "name" field.
func (cuo *CheckUpdateOne) SetName(s string) *CheckUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetWeight sets the "weight" field.
func (cuo *CheckUpdateOne) SetWeight(i int) *CheckUpdateOne {
	cuo.mutation.ResetWeight()
	cuo.mutation.SetWeight(i)
	return cuo
}

// AddWeight adds i to the "weight" field.
func (cuo *CheckUpdateOne) AddWeight(i int) *CheckUpdateOne {
	cuo.mutation.AddWeight(i)
	return cuo
}

// AddStatuIDs adds the "status" edge to the Status entity by IDs.
func (cuo *CheckUpdateOne) AddStatuIDs(ids ...int) *CheckUpdateOne {
	cuo.mutation.AddStatuIDs(ids...)
	return cuo
}

// AddStatus adds the "status" edges to the Status entity.
func (cuo *CheckUpdateOne) AddStatus(s ...*Status) *CheckUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddStatuIDs(ids...)
}

// AddCredentialIDs adds the "credential" edge to the Credential entity by IDs.
func (cuo *CheckUpdateOne) AddCredentialIDs(ids ...int) *CheckUpdateOne {
	cuo.mutation.AddCredentialIDs(ids...)
	return cuo
}

// AddCredential adds the "credential" edges to the Credential entity.
func (cuo *CheckUpdateOne) AddCredential(c ...*Credential) *CheckUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCredentialIDs(ids...)
}

// Mutation returns the CheckMutation object of the builder.
func (cuo *CheckUpdateOne) Mutation() *CheckMutation {
	return cuo.mutation
}

// ClearStatus clears all "status" edges to the Status entity.
func (cuo *CheckUpdateOne) ClearStatus() *CheckUpdateOne {
	cuo.mutation.ClearStatus()
	return cuo
}

// RemoveStatuIDs removes the "status" edge to Status entities by IDs.
func (cuo *CheckUpdateOne) RemoveStatuIDs(ids ...int) *CheckUpdateOne {
	cuo.mutation.RemoveStatuIDs(ids...)
	return cuo
}

// RemoveStatus removes "status" edges to Status entities.
func (cuo *CheckUpdateOne) RemoveStatus(s ...*Status) *CheckUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveStatuIDs(ids...)
}

// ClearCredential clears all "credential" edges to the Credential entity.
func (cuo *CheckUpdateOne) ClearCredential() *CheckUpdateOne {
	cuo.mutation.ClearCredential()
	return cuo
}

// RemoveCredentialIDs removes the "credential" edge to Credential entities by IDs.
func (cuo *CheckUpdateOne) RemoveCredentialIDs(ids ...int) *CheckUpdateOne {
	cuo.mutation.RemoveCredentialIDs(ids...)
	return cuo
}

// RemoveCredential removes "credential" edges to Credential entities.
func (cuo *CheckUpdateOne) RemoveCredential(c ...*Credential) *CheckUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCredentialIDs(ids...)
}

// Where appends a list predicates to the CheckUpdate builder.
func (cuo *CheckUpdateOne) Where(ps ...predicate.Check) *CheckUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CheckUpdateOne) Select(field string, fields ...string) *CheckUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Check entity.
func (cuo *CheckUpdateOne) Save(ctx context.Context) (*Check, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CheckUpdateOne) SaveX(ctx context.Context) *Check {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CheckUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CheckUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CheckUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := check.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Check.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Weight(); ok {
		if err := check.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Check.weight": %w`, err)}
		}
	}
	return nil
}

func (cuo *CheckUpdateOne) sqlSave(ctx context.Context) (_node *Check, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Check.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, check.FieldID)
		for _, f := range fields {
			if !check.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != check.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(check.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Weight(); ok {
		_spec.SetField(check.FieldWeight, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedWeight(); ok {
		_spec.AddField(check.FieldWeight, field.TypeInt, value)
	}
	if cuo.mutation.StatusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStatusIDs(); len(nodes) > 0 && !cuo.mutation.StatusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StatusIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CredentialCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCredentialIDs(); len(nodes) > 0 && !cuo.mutation.CredentialCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CredentialIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Check{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{check.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}