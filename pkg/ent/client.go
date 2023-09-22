// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/compscore/compscore/pkg/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/team"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Check is the client for interacting with the Check builders.
	Check *CheckClient
	// Credential is the client for interacting with the Credential builders.
	Credential *CredentialClient
	// Round is the client for interacting with the Round builders.
	Round *RoundClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Check = NewCheckClient(c.config)
	c.Credential = NewCredentialClient(c.config)
	c.Round = NewRoundClient(c.config)
	c.Status = NewStatusClient(c.config)
	c.Team = NewTeamClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Check:      NewCheckClient(cfg),
		Credential: NewCredentialClient(cfg),
		Round:      NewRoundClient(cfg),
		Status:     NewStatusClient(cfg),
		Team:       NewTeamClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Check:      NewCheckClient(cfg),
		Credential: NewCredentialClient(cfg),
		Round:      NewRoundClient(cfg),
		Status:     NewStatusClient(cfg),
		Team:       NewTeamClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Check.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Check.Use(hooks...)
	c.Credential.Use(hooks...)
	c.Round.Use(hooks...)
	c.Status.Use(hooks...)
	c.Team.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Check.Intercept(interceptors...)
	c.Credential.Intercept(interceptors...)
	c.Round.Intercept(interceptors...)
	c.Status.Intercept(interceptors...)
	c.Team.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CheckMutation:
		return c.Check.mutate(ctx, m)
	case *CredentialMutation:
		return c.Credential.mutate(ctx, m)
	case *RoundMutation:
		return c.Round.mutate(ctx, m)
	case *StatusMutation:
		return c.Status.mutate(ctx, m)
	case *TeamMutation:
		return c.Team.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CheckClient is a client for the Check schema.
type CheckClient struct {
	config
}

// NewCheckClient returns a client for the Check from the given config.
func NewCheckClient(c config) *CheckClient {
	return &CheckClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `check.Hooks(f(g(h())))`.
func (c *CheckClient) Use(hooks ...Hook) {
	c.hooks.Check = append(c.hooks.Check, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `check.Intercept(f(g(h())))`.
func (c *CheckClient) Intercept(interceptors ...Interceptor) {
	c.inters.Check = append(c.inters.Check, interceptors...)
}

// Create returns a builder for creating a Check entity.
func (c *CheckClient) Create() *CheckCreate {
	mutation := newCheckMutation(c.config, OpCreate)
	return &CheckCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Check entities.
func (c *CheckClient) CreateBulk(builders ...*CheckCreate) *CheckCreateBulk {
	return &CheckCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Check.
func (c *CheckClient) Update() *CheckUpdate {
	mutation := newCheckMutation(c.config, OpUpdate)
	return &CheckUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CheckClient) UpdateOne(ch *Check) *CheckUpdateOne {
	mutation := newCheckMutation(c.config, OpUpdateOne, withCheck(ch))
	return &CheckUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CheckClient) UpdateOneID(id int) *CheckUpdateOne {
	mutation := newCheckMutation(c.config, OpUpdateOne, withCheckID(id))
	return &CheckUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Check.
func (c *CheckClient) Delete() *CheckDelete {
	mutation := newCheckMutation(c.config, OpDelete)
	return &CheckDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CheckClient) DeleteOne(ch *Check) *CheckDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CheckClient) DeleteOneID(id int) *CheckDeleteOne {
	builder := c.Delete().Where(check.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CheckDeleteOne{builder}
}

// Query returns a query builder for Check.
func (c *CheckClient) Query() *CheckQuery {
	return &CheckQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCheck},
		inters: c.Interceptors(),
	}
}

// Get returns a Check entity by its id.
func (c *CheckClient) Get(ctx context.Context, id int) (*Check, error) {
	return c.Query().Where(check.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CheckClient) GetX(ctx context.Context, id int) *Check {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStatus queries the status edge of a Check.
func (c *CheckClient) QueryStatus(ch *Check) *StatusQuery {
	query := (&StatusClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(check.Table, check.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, check.StatusTable, check.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCredential queries the credential edge of a Check.
func (c *CheckClient) QueryCredential(ch *Check) *CredentialQuery {
	query := (&CredentialClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(check.Table, check.FieldID, id),
			sqlgraph.To(credential.Table, credential.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, check.CredentialTable, check.CredentialColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CheckClient) Hooks() []Hook {
	return c.hooks.Check
}

// Interceptors returns the client interceptors.
func (c *CheckClient) Interceptors() []Interceptor {
	return c.inters.Check
}

func (c *CheckClient) mutate(ctx context.Context, m *CheckMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CheckCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CheckUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CheckUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CheckDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Check mutation op: %q", m.Op())
	}
}

// CredentialClient is a client for the Credential schema.
type CredentialClient struct {
	config
}

// NewCredentialClient returns a client for the Credential from the given config.
func NewCredentialClient(c config) *CredentialClient {
	return &CredentialClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `credential.Hooks(f(g(h())))`.
func (c *CredentialClient) Use(hooks ...Hook) {
	c.hooks.Credential = append(c.hooks.Credential, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `credential.Intercept(f(g(h())))`.
func (c *CredentialClient) Intercept(interceptors ...Interceptor) {
	c.inters.Credential = append(c.inters.Credential, interceptors...)
}

// Create returns a builder for creating a Credential entity.
func (c *CredentialClient) Create() *CredentialCreate {
	mutation := newCredentialMutation(c.config, OpCreate)
	return &CredentialCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Credential entities.
func (c *CredentialClient) CreateBulk(builders ...*CredentialCreate) *CredentialCreateBulk {
	return &CredentialCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Credential.
func (c *CredentialClient) Update() *CredentialUpdate {
	mutation := newCredentialMutation(c.config, OpUpdate)
	return &CredentialUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CredentialClient) UpdateOne(cr *Credential) *CredentialUpdateOne {
	mutation := newCredentialMutation(c.config, OpUpdateOne, withCredential(cr))
	return &CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CredentialClient) UpdateOneID(id int) *CredentialUpdateOne {
	mutation := newCredentialMutation(c.config, OpUpdateOne, withCredentialID(id))
	return &CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Credential.
func (c *CredentialClient) Delete() *CredentialDelete {
	mutation := newCredentialMutation(c.config, OpDelete)
	return &CredentialDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CredentialClient) DeleteOne(cr *Credential) *CredentialDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CredentialClient) DeleteOneID(id int) *CredentialDeleteOne {
	builder := c.Delete().Where(credential.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CredentialDeleteOne{builder}
}

// Query returns a query builder for Credential.
func (c *CredentialClient) Query() *CredentialQuery {
	return &CredentialQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCredential},
		inters: c.Interceptors(),
	}
}

// Get returns a Credential entity by its id.
func (c *CredentialClient) Get(ctx context.Context, id int) (*Credential, error) {
	return c.Query().Where(credential.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CredentialClient) GetX(ctx context.Context, id int) *Credential {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCheck queries the check edge of a Credential.
func (c *CredentialClient) QueryCheck(cr *Credential) *CheckQuery {
	query := (&CheckClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(credential.Table, credential.FieldID, id),
			sqlgraph.To(check.Table, check.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, credential.CheckTable, credential.CheckColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTeam queries the team edge of a Credential.
func (c *CredentialClient) QueryTeam(cr *Credential) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(credential.Table, credential.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, credential.TeamTable, credential.TeamColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CredentialClient) Hooks() []Hook {
	return c.hooks.Credential
}

// Interceptors returns the client interceptors.
func (c *CredentialClient) Interceptors() []Interceptor {
	return c.inters.Credential
}

func (c *CredentialClient) mutate(ctx context.Context, m *CredentialMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CredentialCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CredentialUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CredentialDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Credential mutation op: %q", m.Op())
	}
}

// RoundClient is a client for the Round schema.
type RoundClient struct {
	config
}

// NewRoundClient returns a client for the Round from the given config.
func NewRoundClient(c config) *RoundClient {
	return &RoundClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `round.Hooks(f(g(h())))`.
func (c *RoundClient) Use(hooks ...Hook) {
	c.hooks.Round = append(c.hooks.Round, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `round.Intercept(f(g(h())))`.
func (c *RoundClient) Intercept(interceptors ...Interceptor) {
	c.inters.Round = append(c.inters.Round, interceptors...)
}

// Create returns a builder for creating a Round entity.
func (c *RoundClient) Create() *RoundCreate {
	mutation := newRoundMutation(c.config, OpCreate)
	return &RoundCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Round entities.
func (c *RoundClient) CreateBulk(builders ...*RoundCreate) *RoundCreateBulk {
	return &RoundCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Round.
func (c *RoundClient) Update() *RoundUpdate {
	mutation := newRoundMutation(c.config, OpUpdate)
	return &RoundUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoundClient) UpdateOne(r *Round) *RoundUpdateOne {
	mutation := newRoundMutation(c.config, OpUpdateOne, withRound(r))
	return &RoundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoundClient) UpdateOneID(id int) *RoundUpdateOne {
	mutation := newRoundMutation(c.config, OpUpdateOne, withRoundID(id))
	return &RoundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Round.
func (c *RoundClient) Delete() *RoundDelete {
	mutation := newRoundMutation(c.config, OpDelete)
	return &RoundDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoundClient) DeleteOne(r *Round) *RoundDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoundClient) DeleteOneID(id int) *RoundDeleteOne {
	builder := c.Delete().Where(round.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoundDeleteOne{builder}
}

// Query returns a query builder for Round.
func (c *RoundClient) Query() *RoundQuery {
	return &RoundQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRound},
		inters: c.Interceptors(),
	}
}

// Get returns a Round entity by its id.
func (c *RoundClient) Get(ctx context.Context, id int) (*Round, error) {
	return c.Query().Where(round.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoundClient) GetX(ctx context.Context, id int) *Round {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStatus queries the status edge of a Round.
func (c *RoundClient) QueryStatus(r *Round) *StatusQuery {
	query := (&StatusClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(round.Table, round.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, round.StatusTable, round.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoundClient) Hooks() []Hook {
	return c.hooks.Round
}

// Interceptors returns the client interceptors.
func (c *RoundClient) Interceptors() []Interceptor {
	return c.inters.Round
}

func (c *RoundClient) mutate(ctx context.Context, m *RoundMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoundCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoundUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoundDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Round mutation op: %q", m.Op())
	}
}

// StatusClient is a client for the Status schema.
type StatusClient struct {
	config
}

// NewStatusClient returns a client for the Status from the given config.
func NewStatusClient(c config) *StatusClient {
	return &StatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `status.Hooks(f(g(h())))`.
func (c *StatusClient) Use(hooks ...Hook) {
	c.hooks.Status = append(c.hooks.Status, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `status.Intercept(f(g(h())))`.
func (c *StatusClient) Intercept(interceptors ...Interceptor) {
	c.inters.Status = append(c.inters.Status, interceptors...)
}

// Create returns a builder for creating a Status entity.
func (c *StatusClient) Create() *StatusCreate {
	mutation := newStatusMutation(c.config, OpCreate)
	return &StatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Status entities.
func (c *StatusClient) CreateBulk(builders ...*StatusCreate) *StatusCreateBulk {
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Status.
func (c *StatusClient) Update() *StatusUpdate {
	mutation := newStatusMutation(c.config, OpUpdate)
	return &StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusClient) UpdateOne(s *Status) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatus(s))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusClient) UpdateOneID(id int) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatusID(id))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Status.
func (c *StatusClient) Delete() *StatusDelete {
	mutation := newStatusMutation(c.config, OpDelete)
	return &StatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatusClient) DeleteOne(s *Status) *StatusDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StatusClient) DeleteOneID(id int) *StatusDeleteOne {
	builder := c.Delete().Where(status.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusDeleteOne{builder}
}

// Query returns a query builder for Status.
func (c *StatusClient) Query() *StatusQuery {
	return &StatusQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStatus},
		inters: c.Interceptors(),
	}
}

// Get returns a Status entity by its id.
func (c *StatusClient) Get(ctx context.Context, id int) (*Status, error) {
	return c.Query().Where(status.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusClient) GetX(ctx context.Context, id int) *Status {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCheck queries the check edge of a Status.
func (c *StatusClient) QueryCheck(s *Status) *CheckQuery {
	query := (&CheckClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(check.Table, check.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, status.CheckTable, status.CheckColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTeam queries the team edge of a Status.
func (c *StatusClient) QueryTeam(s *Status) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, status.TeamTable, status.TeamColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRound queries the round edge of a Status.
func (c *StatusClient) QueryRound(s *Status) *RoundQuery {
	query := (&RoundClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(round.Table, round.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, status.RoundTable, status.RoundColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatusClient) Hooks() []Hook {
	return c.hooks.Status
}

// Interceptors returns the client interceptors.
func (c *StatusClient) Interceptors() []Interceptor {
	return c.inters.Status
}

func (c *StatusClient) mutate(ctx context.Context, m *StatusMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StatusCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StatusDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Status mutation op: %q", m.Op())
	}
}

// TeamClient is a client for the Team schema.
type TeamClient struct {
	config
}

// NewTeamClient returns a client for the Team from the given config.
func NewTeamClient(c config) *TeamClient {
	return &TeamClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `team.Hooks(f(g(h())))`.
func (c *TeamClient) Use(hooks ...Hook) {
	c.hooks.Team = append(c.hooks.Team, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `team.Intercept(f(g(h())))`.
func (c *TeamClient) Intercept(interceptors ...Interceptor) {
	c.inters.Team = append(c.inters.Team, interceptors...)
}

// Create returns a builder for creating a Team entity.
func (c *TeamClient) Create() *TeamCreate {
	mutation := newTeamMutation(c.config, OpCreate)
	return &TeamCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Team entities.
func (c *TeamClient) CreateBulk(builders ...*TeamCreate) *TeamCreateBulk {
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Team.
func (c *TeamClient) Update() *TeamUpdate {
	mutation := newTeamMutation(c.config, OpUpdate)
	return &TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeamClient) UpdateOne(t *Team) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeam(t))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeamClient) UpdateOneID(id int) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeamID(id))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Team.
func (c *TeamClient) Delete() *TeamDelete {
	mutation := newTeamMutation(c.config, OpDelete)
	return &TeamDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TeamClient) DeleteOne(t *Team) *TeamDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TeamClient) DeleteOneID(id int) *TeamDeleteOne {
	builder := c.Delete().Where(team.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeamDeleteOne{builder}
}

// Query returns a query builder for Team.
func (c *TeamClient) Query() *TeamQuery {
	return &TeamQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTeam},
		inters: c.Interceptors(),
	}
}

// Get returns a Team entity by its id.
func (c *TeamClient) Get(ctx context.Context, id int) (*Team, error) {
	return c.Query().Where(team.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeamClient) GetX(ctx context.Context, id int) *Team {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStatus queries the status edge of a Team.
func (c *TeamClient) QueryStatus(t *Team) *StatusQuery {
	query := (&StatusClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, team.StatusTable, team.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCredential queries the credential edge of a Team.
func (c *TeamClient) QueryCredential(t *Team) *CredentialQuery {
	query := (&CredentialClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(credential.Table, credential.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, team.CredentialTable, team.CredentialColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TeamClient) Hooks() []Hook {
	return c.hooks.Team
}

// Interceptors returns the client interceptors.
func (c *TeamClient) Interceptors() []Interceptor {
	return c.inters.Team
}

func (c *TeamClient) mutate(ctx context.Context, m *TeamMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TeamCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TeamDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Team mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Check, Credential, Round, Status, Team []ent.Hook
	}
	inters struct {
		Check, Credential, Round, Status, Team []ent.Interceptor
	}
)
