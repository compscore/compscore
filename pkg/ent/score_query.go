// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compscore/compscore/pkg/ent/predicate"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/score"
	"github.com/compscore/compscore/pkg/ent/user"
)

// ScoreQuery is the builder for querying Score entities.
type ScoreQuery struct {
	config
	ctx        *QueryContext
	order      []score.OrderOption
	inters     []Interceptor
	predicates []predicate.Score
	withRound  *RoundQuery
	withUser   *UserQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Score) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScoreQuery builder.
func (sq *ScoreQuery) Where(ps ...predicate.Score) *ScoreQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *ScoreQuery) Limit(limit int) *ScoreQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *ScoreQuery) Offset(offset int) *ScoreQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ScoreQuery) Unique(unique bool) *ScoreQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *ScoreQuery) Order(o ...score.OrderOption) *ScoreQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryRound chains the current query on the "round" edge.
func (sq *ScoreQuery) QueryRound() *RoundQuery {
	query := (&RoundClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(score.Table, score.FieldID, selector),
			sqlgraph.To(round.Table, round.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, score.RoundTable, score.RoundColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (sq *ScoreQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(score.Table, score.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, score.UserTable, score.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Score entity from the query.
// Returns a *NotFoundError when no Score was found.
func (sq *ScoreQuery) First(ctx context.Context) (*Score, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{score.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ScoreQuery) FirstX(ctx context.Context) *Score {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Score ID from the query.
// Returns a *NotFoundError when no Score ID was found.
func (sq *ScoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{score.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ScoreQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Score entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Score entity is found.
// Returns a *NotFoundError when no Score entities are found.
func (sq *ScoreQuery) Only(ctx context.Context) (*Score, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{score.Label}
	default:
		return nil, &NotSingularError{score.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ScoreQuery) OnlyX(ctx context.Context) *Score {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Score ID in the query.
// Returns a *NotSingularError when more than one Score ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ScoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{score.Label}
	default:
		err = &NotSingularError{score.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ScoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Scores.
func (sq *ScoreQuery) All(ctx context.Context) ([]*Score, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Score, *ScoreQuery]()
	return withInterceptors[[]*Score](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *ScoreQuery) AllX(ctx context.Context) []*Score {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Score IDs.
func (sq *ScoreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(score.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ScoreQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ScoreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*ScoreQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ScoreQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ScoreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ScoreQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ScoreQuery) Clone() *ScoreQuery {
	if sq == nil {
		return nil
	}
	return &ScoreQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]score.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Score{}, sq.predicates...),
		withRound:  sq.withRound.Clone(),
		withUser:   sq.withUser.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithRound tells the query-builder to eager-load the nodes that are connected to
// the "round" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScoreQuery) WithRound(opts ...func(*RoundQuery)) *ScoreQuery {
	query := (&RoundClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withRound = query
	return sq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScoreQuery) WithUser(opts ...func(*UserQuery)) *ScoreQuery {
	query := (&UserClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withUser = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Score.Query().
//		GroupBy(score.FieldScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *ScoreQuery) GroupBy(field string, fields ...string) *ScoreGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScoreGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = score.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score"`
//	}
//
//	client.Score.Query().
//		Select(score.FieldScore).
//		Scan(ctx, &v)
func (sq *ScoreQuery) Select(fields ...string) *ScoreSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &ScoreSelect{ScoreQuery: sq}
	sbuild.label = score.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScoreSelect configured with the given aggregations.
func (sq *ScoreQuery) Aggregate(fns ...AggregateFunc) *ScoreSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *ScoreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !score.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ScoreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Score, error) {
	var (
		nodes       = []*Score{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withRound != nil,
			sq.withUser != nil,
		}
	)
	if sq.withRound != nil || sq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, score.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Score).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Score{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withRound; query != nil {
		if err := sq.loadRound(ctx, query, nodes, nil,
			func(n *Score, e *Round) { n.Edges.Round = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withUser; query != nil {
		if err := sq.loadUser(ctx, query, nodes, nil,
			func(n *Score, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *ScoreQuery) loadRound(ctx context.Context, query *RoundQuery, nodes []*Score, init func(*Score), assign func(*Score, *Round)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Score)
	for i := range nodes {
		if nodes[i].round_scores == nil {
			continue
		}
		fk := *nodes[i].round_scores
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(round.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "round_scores" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *ScoreQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Score, init func(*Score), assign func(*Score, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Score)
	for i := range nodes {
		if nodes[i].user_scores == nil {
			continue
		}
		fk := *nodes[i].user_scores
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
			return fmt.Errorf(`unexpected foreign-key "user_scores" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *ScoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ScoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(score.Table, score.Columns, sqlgraph.NewFieldSpec(score.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, score.FieldID)
		for i := range fields {
			if fields[i] != score.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ScoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(score.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = score.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScoreGroupBy is the group-by builder for Score entities.
type ScoreGroupBy struct {
	selector
	build *ScoreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ScoreGroupBy) Aggregate(fns ...AggregateFunc) *ScoreGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *ScoreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScoreQuery, *ScoreGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *ScoreGroupBy) sqlScan(ctx context.Context, root *ScoreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScoreSelect is the builder for selecting fields of Score entities.
type ScoreSelect struct {
	*ScoreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *ScoreSelect) Aggregate(fns ...AggregateFunc) *ScoreSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ScoreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScoreQuery, *ScoreSelect](ctx, ss.ScoreQuery, ss, ss.inters, v)
}

func (ss *ScoreSelect) sqlScan(ctx context.Context, root *ScoreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
