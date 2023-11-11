// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/predicate"
	"github.com/compscore/compscore/pkg/ent/user"
)

// CredentialQuery is the builder for querying Credential entities.
type CredentialQuery struct {
	config
	ctx        *QueryContext
	order      []credential.OrderOption
	inters     []Interceptor
	predicates []predicate.Credential
	withUser   *UserQuery
	withCheck  *CheckQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Credential) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CredentialQuery builder.
func (cq *CredentialQuery) Where(ps ...predicate.Credential) *CredentialQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CredentialQuery) Limit(limit int) *CredentialQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CredentialQuery) Offset(offset int) *CredentialQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CredentialQuery) Unique(unique bool) *CredentialQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CredentialQuery) Order(o ...credential.OrderOption) *CredentialQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryUser chains the current query on the "user" edge.
func (cq *CredentialQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(credential.Table, credential.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, credential.UserTable, credential.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCheck chains the current query on the "check" edge.
func (cq *CredentialQuery) QueryCheck() *CheckQuery {
	query := (&CheckClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(credential.Table, credential.FieldID, selector),
			sqlgraph.To(check.Table, check.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, credential.CheckTable, credential.CheckColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Credential entity from the query.
// Returns a *NotFoundError when no Credential was found.
func (cq *CredentialQuery) First(ctx context.Context) (*Credential, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{credential.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CredentialQuery) FirstX(ctx context.Context) *Credential {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Credential ID from the query.
// Returns a *NotFoundError when no Credential ID was found.
func (cq *CredentialQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{credential.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CredentialQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Credential entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Credential entity is found.
// Returns a *NotFoundError when no Credential entities are found.
func (cq *CredentialQuery) Only(ctx context.Context) (*Credential, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{credential.Label}
	default:
		return nil, &NotSingularError{credential.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CredentialQuery) OnlyX(ctx context.Context) *Credential {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Credential ID in the query.
// Returns a *NotSingularError when more than one Credential ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CredentialQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{credential.Label}
	default:
		err = &NotSingularError{credential.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CredentialQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Credentials.
func (cq *CredentialQuery) All(ctx context.Context) ([]*Credential, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Credential, *CredentialQuery]()
	return withInterceptors[[]*Credential](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CredentialQuery) AllX(ctx context.Context) []*Credential {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Credential IDs.
func (cq *CredentialQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(credential.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CredentialQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CredentialQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CredentialQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CredentialQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CredentialQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CredentialQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CredentialQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CredentialQuery) Clone() *CredentialQuery {
	if cq == nil {
		return nil
	}
	return &CredentialQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]credential.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Credential{}, cq.predicates...),
		withUser:   cq.withUser.Clone(),
		withCheck:  cq.withCheck.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CredentialQuery) WithUser(opts ...func(*UserQuery)) *CredentialQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withUser = query
	return cq
}

// WithCheck tells the query-builder to eager-load the nodes that are connected to
// the "check" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CredentialQuery) WithCheck(opts ...func(*CheckQuery)) *CredentialQuery {
	query := (&CheckClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCheck = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Credential.Query().
//		GroupBy(credential.FieldPassword).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CredentialQuery) GroupBy(field string, fields ...string) *CredentialGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CredentialGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = credential.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password"`
//	}
//
//	client.Credential.Query().
//		Select(credential.FieldPassword).
//		Scan(ctx, &v)
func (cq *CredentialQuery) Select(fields ...string) *CredentialSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CredentialSelect{CredentialQuery: cq}
	sbuild.label = credential.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CredentialSelect configured with the given aggregations.
func (cq *CredentialQuery) Aggregate(fns ...AggregateFunc) *CredentialSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CredentialQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !credential.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CredentialQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Credential, error) {
	var (
		nodes       = []*Credential{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withUser != nil,
			cq.withCheck != nil,
		}
	)
	if cq.withUser != nil || cq.withCheck != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, credential.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Credential).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Credential{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withUser; query != nil {
		if err := cq.loadUser(ctx, query, nodes, nil,
			func(n *Credential, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCheck; query != nil {
		if err := cq.loadCheck(ctx, query, nodes, nil,
			func(n *Credential, e *Check) { n.Edges.Check = e }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CredentialQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Credential, init func(*Credential), assign func(*Credential, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Credential)
	for i := range nodes {
		if nodes[i].credential_user == nil {
			continue
		}
		fk := *nodes[i].credential_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "credential_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CredentialQuery) loadCheck(ctx context.Context, query *CheckQuery, nodes []*Credential, init func(*Credential), assign func(*Credential, *Check)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Credential)
	for i := range nodes {
		if nodes[i].credential_check == nil {
			continue
		}
		fk := *nodes[i].credential_check
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(check.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "credential_check" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *CredentialQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CredentialQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(credential.Table, credential.Columns, sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, credential.FieldID)
		for i := range fields {
			if fields[i] != credential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CredentialQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(credential.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = credential.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CredentialGroupBy is the group-by builder for Credential entities.
type CredentialGroupBy struct {
	selector
	build *CredentialQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CredentialGroupBy) Aggregate(fns ...AggregateFunc) *CredentialGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CredentialGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CredentialQuery, *CredentialGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CredentialGroupBy) sqlScan(ctx context.Context, root *CredentialQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CredentialSelect is the builder for selecting fields of Credential entities.
type CredentialSelect struct {
	*CredentialQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CredentialSelect) Aggregate(fns ...AggregateFunc) *CredentialSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CredentialSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CredentialQuery, *CredentialSelect](ctx, cs.CredentialQuery, cs, cs.inters, v)
}

func (cs *CredentialSelect) sqlScan(ctx context.Context, root *CredentialQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
