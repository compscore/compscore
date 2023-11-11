// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CheckEdge is the edge representation of Check.
type CheckEdge struct {
	Node   *Check `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// CheckConnection is the connection containing edges to Check.
type CheckConnection struct {
	Edges      []*CheckEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *CheckConnection) build(nodes []*Check, pager *checkPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Check
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Check {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Check {
			return nodes[i]
		}
	}
	c.Edges = make([]*CheckEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CheckEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CheckPaginateOption enables pagination customization.
type CheckPaginateOption func(*checkPager) error

// WithCheckOrder configures pagination ordering.
func WithCheckOrder(order *CheckOrder) CheckPaginateOption {
	if order == nil {
		order = DefaultCheckOrder
	}
	o := *order
	return func(pager *checkPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCheckOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCheckFilter configures pagination filter.
func WithCheckFilter(filter func(*CheckQuery) (*CheckQuery, error)) CheckPaginateOption {
	return func(pager *checkPager) error {
		if filter == nil {
			return errors.New("CheckQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type checkPager struct {
	reverse bool
	order   *CheckOrder
	filter  func(*CheckQuery) (*CheckQuery, error)
}

func newCheckPager(opts []CheckPaginateOption, reverse bool) (*checkPager, error) {
	pager := &checkPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCheckOrder
	}
	return pager, nil
}

func (p *checkPager) applyFilter(query *CheckQuery) (*CheckQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *checkPager) toCursor(c *Check) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *checkPager) applyCursors(query *CheckQuery, after, before *Cursor) (*CheckQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCheckOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *checkPager) applyOrder(query *CheckQuery) *CheckQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCheckOrder.Field {
		query = query.Order(DefaultCheckOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *checkPager) orderExpr(query *CheckQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCheckOrder.Field {
			b.Comma().Ident(DefaultCheckOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Check.
func (c *CheckQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CheckPaginateOption,
) (*CheckConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCheckPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CheckConnection{Edges: []*CheckEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = c.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CheckOrderField defines the ordering field of Check.
type CheckOrderField struct {
	// Value extracts the ordering value from the given Check.
	Value    func(*Check) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) check.OrderOption
	toCursor func(*Check) Cursor
}

// CheckOrder defines the ordering of Check.
type CheckOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *CheckOrderField `json:"field"`
}

// DefaultCheckOrder is the default ordering of Check.
var DefaultCheckOrder = &CheckOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CheckOrderField{
		Value: func(c *Check) (ent.Value, error) {
			return c.ID, nil
		},
		column: check.FieldID,
		toTerm: check.ByID,
		toCursor: func(c *Check) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Check into CheckEdge.
func (c *Check) ToEdge(order *CheckOrder) *CheckEdge {
	if order == nil {
		order = DefaultCheckOrder
	}
	return &CheckEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// CredentialEdge is the edge representation of Credential.
type CredentialEdge struct {
	Node   *Credential `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// CredentialConnection is the connection containing edges to Credential.
type CredentialConnection struct {
	Edges      []*CredentialEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *CredentialConnection) build(nodes []*Credential, pager *credentialPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Credential
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Credential {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Credential {
			return nodes[i]
		}
	}
	c.Edges = make([]*CredentialEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CredentialEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CredentialPaginateOption enables pagination customization.
type CredentialPaginateOption func(*credentialPager) error

// WithCredentialOrder configures pagination ordering.
func WithCredentialOrder(order *CredentialOrder) CredentialPaginateOption {
	if order == nil {
		order = DefaultCredentialOrder
	}
	o := *order
	return func(pager *credentialPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCredentialOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCredentialFilter configures pagination filter.
func WithCredentialFilter(filter func(*CredentialQuery) (*CredentialQuery, error)) CredentialPaginateOption {
	return func(pager *credentialPager) error {
		if filter == nil {
			return errors.New("CredentialQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type credentialPager struct {
	reverse bool
	order   *CredentialOrder
	filter  func(*CredentialQuery) (*CredentialQuery, error)
}

func newCredentialPager(opts []CredentialPaginateOption, reverse bool) (*credentialPager, error) {
	pager := &credentialPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCredentialOrder
	}
	return pager, nil
}

func (p *credentialPager) applyFilter(query *CredentialQuery) (*CredentialQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *credentialPager) toCursor(c *Credential) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *credentialPager) applyCursors(query *CredentialQuery, after, before *Cursor) (*CredentialQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCredentialOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *credentialPager) applyOrder(query *CredentialQuery) *CredentialQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCredentialOrder.Field {
		query = query.Order(DefaultCredentialOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *credentialPager) orderExpr(query *CredentialQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCredentialOrder.Field {
			b.Comma().Ident(DefaultCredentialOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Credential.
func (c *CredentialQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CredentialPaginateOption,
) (*CredentialConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCredentialPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CredentialConnection{Edges: []*CredentialEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = c.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CredentialOrderField defines the ordering field of Credential.
type CredentialOrderField struct {
	// Value extracts the ordering value from the given Credential.
	Value    func(*Credential) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) credential.OrderOption
	toCursor func(*Credential) Cursor
}

// CredentialOrder defines the ordering of Credential.
type CredentialOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *CredentialOrderField `json:"field"`
}

// DefaultCredentialOrder is the default ordering of Credential.
var DefaultCredentialOrder = &CredentialOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CredentialOrderField{
		Value: func(c *Credential) (ent.Value, error) {
			return c.ID, nil
		},
		column: credential.FieldID,
		toTerm: credential.ByID,
		toCursor: func(c *Credential) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Credential into CredentialEdge.
func (c *Credential) ToEdge(order *CredentialOrder) *CredentialEdge {
	if order == nil {
		order = DefaultCredentialOrder
	}
	return &CredentialEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// RoundEdge is the edge representation of Round.
type RoundEdge struct {
	Node   *Round `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// RoundConnection is the connection containing edges to Round.
type RoundConnection struct {
	Edges      []*RoundEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *RoundConnection) build(nodes []*Round, pager *roundPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Round
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Round {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Round {
			return nodes[i]
		}
	}
	c.Edges = make([]*RoundEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &RoundEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// RoundPaginateOption enables pagination customization.
type RoundPaginateOption func(*roundPager) error

// WithRoundOrder configures pagination ordering.
func WithRoundOrder(order *RoundOrder) RoundPaginateOption {
	if order == nil {
		order = DefaultRoundOrder
	}
	o := *order
	return func(pager *roundPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultRoundOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithRoundFilter configures pagination filter.
func WithRoundFilter(filter func(*RoundQuery) (*RoundQuery, error)) RoundPaginateOption {
	return func(pager *roundPager) error {
		if filter == nil {
			return errors.New("RoundQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type roundPager struct {
	reverse bool
	order   *RoundOrder
	filter  func(*RoundQuery) (*RoundQuery, error)
}

func newRoundPager(opts []RoundPaginateOption, reverse bool) (*roundPager, error) {
	pager := &roundPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRoundOrder
	}
	return pager, nil
}

func (p *roundPager) applyFilter(query *RoundQuery) (*RoundQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *roundPager) toCursor(r *Round) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *roundPager) applyCursors(query *RoundQuery, after, before *Cursor) (*RoundQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultRoundOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *roundPager) applyOrder(query *RoundQuery) *RoundQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultRoundOrder.Field {
		query = query.Order(DefaultRoundOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *roundPager) orderExpr(query *RoundQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultRoundOrder.Field {
			b.Comma().Ident(DefaultRoundOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Round.
func (r *RoundQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...RoundPaginateOption,
) (*RoundConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newRoundPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}
	conn := &RoundConnection{Edges: []*RoundEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = r.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if r, err = pager.applyCursors(r, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		r.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	r = pager.applyOrder(r)
	nodes, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// RoundOrderField defines the ordering field of Round.
type RoundOrderField struct {
	// Value extracts the ordering value from the given Round.
	Value    func(*Round) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) round.OrderOption
	toCursor func(*Round) Cursor
}

// RoundOrder defines the ordering of Round.
type RoundOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *RoundOrderField `json:"field"`
}

// DefaultRoundOrder is the default ordering of Round.
var DefaultRoundOrder = &RoundOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &RoundOrderField{
		Value: func(r *Round) (ent.Value, error) {
			return r.ID, nil
		},
		column: round.FieldID,
		toTerm: round.ByID,
		toCursor: func(r *Round) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Round into RoundEdge.
func (r *Round) ToEdge(order *RoundOrder) *RoundEdge {
	if order == nil {
		order = DefaultRoundOrder
	}
	return &RoundEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}

// StatusEdge is the edge representation of Status.
type StatusEdge struct {
	Node   *Status `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// StatusConnection is the connection containing edges to Status.
type StatusConnection struct {
	Edges      []*StatusEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *StatusConnection) build(nodes []*Status, pager *statusPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Status
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Status {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Status {
			return nodes[i]
		}
	}
	c.Edges = make([]*StatusEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &StatusEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// StatusPaginateOption enables pagination customization.
type StatusPaginateOption func(*statusPager) error

// WithStatusOrder configures pagination ordering.
func WithStatusOrder(order *StatusOrder) StatusPaginateOption {
	if order == nil {
		order = DefaultStatusOrder
	}
	o := *order
	return func(pager *statusPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultStatusOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithStatusFilter configures pagination filter.
func WithStatusFilter(filter func(*StatusQuery) (*StatusQuery, error)) StatusPaginateOption {
	return func(pager *statusPager) error {
		if filter == nil {
			return errors.New("StatusQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type statusPager struct {
	reverse bool
	order   *StatusOrder
	filter  func(*StatusQuery) (*StatusQuery, error)
}

func newStatusPager(opts []StatusPaginateOption, reverse bool) (*statusPager, error) {
	pager := &statusPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultStatusOrder
	}
	return pager, nil
}

func (p *statusPager) applyFilter(query *StatusQuery) (*StatusQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *statusPager) toCursor(s *Status) Cursor {
	return p.order.Field.toCursor(s)
}

func (p *statusPager) applyCursors(query *StatusQuery, after, before *Cursor) (*StatusQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultStatusOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *statusPager) applyOrder(query *StatusQuery) *StatusQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultStatusOrder.Field {
		query = query.Order(DefaultStatusOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *statusPager) orderExpr(query *StatusQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultStatusOrder.Field {
			b.Comma().Ident(DefaultStatusOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Status.
func (s *StatusQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...StatusPaginateOption,
) (*StatusConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newStatusPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if s, err = pager.applyFilter(s); err != nil {
		return nil, err
	}
	conn := &StatusConnection{Edges: []*StatusEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = s.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if s, err = pager.applyCursors(s, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		s.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := s.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	s = pager.applyOrder(s)
	nodes, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// StatusOrderField defines the ordering field of Status.
type StatusOrderField struct {
	// Value extracts the ordering value from the given Status.
	Value    func(*Status) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) status.OrderOption
	toCursor func(*Status) Cursor
}

// StatusOrder defines the ordering of Status.
type StatusOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *StatusOrderField `json:"field"`
}

// DefaultStatusOrder is the default ordering of Status.
var DefaultStatusOrder = &StatusOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &StatusOrderField{
		Value: func(s *Status) (ent.Value, error) {
			return s.ID, nil
		},
		column: status.FieldID,
		toTerm: status.ByID,
		toCursor: func(s *Status) Cursor {
			return Cursor{ID: s.ID}
		},
	},
}

// ToEdge converts Status into StatusEdge.
func (s *Status) ToEdge(order *StatusOrder) *StatusEdge {
	if order == nil {
		order = DefaultStatusOrder
	}
	return &StatusEdge{
		Node:   s,
		Cursor: order.Field.toCursor(s),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}