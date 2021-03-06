// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/pattapong1/app/ent/migrate"

	"github.com/pattapong1/app/ent/activity"
	"github.com/pattapong1/app/ent/club"
	"github.com/pattapong1/app/ent/clubtypes"
	"github.com/pattapong1/app/ent/location"
	"github.com/pattapong1/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Activity is the client for interacting with the Activity builders.
	Activity *ActivityClient
	// Club is the client for interacting with the Club builders.
	Club *ClubClient
	// ClubTypes is the client for interacting with the ClubTypes builders.
	ClubTypes *ClubTypesClient
	// Location is the client for interacting with the Location builders.
	Location *LocationClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Activity = NewActivityClient(c.config)
	c.Club = NewClubClient(c.config)
	c.ClubTypes = NewClubTypesClient(c.config)
	c.Location = NewLocationClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Activity:  NewActivityClient(cfg),
		Club:      NewClubClient(cfg),
		ClubTypes: NewClubTypesClient(cfg),
		Location:  NewLocationClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:    cfg,
		Activity:  NewActivityClient(cfg),
		Club:      NewClubClient(cfg),
		ClubTypes: NewClubTypesClient(cfg),
		Location:  NewLocationClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Activity.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Activity.Use(hooks...)
	c.Club.Use(hooks...)
	c.ClubTypes.Use(hooks...)
	c.Location.Use(hooks...)
	c.User.Use(hooks...)
}

// ActivityClient is a client for the Activity schema.
type ActivityClient struct {
	config
}

// NewActivityClient returns a client for the Activity from the given config.
func NewActivityClient(c config) *ActivityClient {
	return &ActivityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `activity.Hooks(f(g(h())))`.
func (c *ActivityClient) Use(hooks ...Hook) {
	c.hooks.Activity = append(c.hooks.Activity, hooks...)
}

// Create returns a create builder for Activity.
func (c *ActivityClient) Create() *ActivityCreate {
	mutation := newActivityMutation(c.config, OpCreate)
	return &ActivityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Activity.
func (c *ActivityClient) Update() *ActivityUpdate {
	mutation := newActivityMutation(c.config, OpUpdate)
	return &ActivityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActivityClient) UpdateOne(a *Activity) *ActivityUpdateOne {
	mutation := newActivityMutation(c.config, OpUpdateOne, withActivity(a))
	return &ActivityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActivityClient) UpdateOneID(id int) *ActivityUpdateOne {
	mutation := newActivityMutation(c.config, OpUpdateOne, withActivityID(id))
	return &ActivityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Activity.
func (c *ActivityClient) Delete() *ActivityDelete {
	mutation := newActivityMutation(c.config, OpDelete)
	return &ActivityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ActivityClient) DeleteOne(a *Activity) *ActivityDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ActivityClient) DeleteOneID(id int) *ActivityDeleteOne {
	builder := c.Delete().Where(activity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActivityDeleteOne{builder}
}

// Create returns a query builder for Activity.
func (c *ActivityClient) Query() *ActivityQuery {
	return &ActivityQuery{config: c.config}
}

// Get returns a Activity entity by its id.
func (c *ActivityClient) Get(ctx context.Context, id int) (*Activity, error) {
	return c.Query().Where(activity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActivityClient) GetX(ctx context.Context, id int) *Activity {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// QueryClub queries the club edge of a Activity.
func (c *ActivityClient) QueryClub(a *Activity) *ClubQuery {
	query := &ClubQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(activity.Table, activity.FieldID, id),
			sqlgraph.To(club.Table, club.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, activity.ClubTable, activity.ClubColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActivityClient) Hooks() []Hook {
	return c.hooks.Activity
}

// ClubClient is a client for the Club schema.
type ClubClient struct {
	config
}

// NewClubClient returns a client for the Club from the given config.
func NewClubClient(c config) *ClubClient {
	return &ClubClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `club.Hooks(f(g(h())))`.
func (c *ClubClient) Use(hooks ...Hook) {
	c.hooks.Club = append(c.hooks.Club, hooks...)
}

// Create returns a create builder for Club.
func (c *ClubClient) Create() *ClubCreate {
	mutation := newClubMutation(c.config, OpCreate)
	return &ClubCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Club.
func (c *ClubClient) Update() *ClubUpdate {
	mutation := newClubMutation(c.config, OpUpdate)
	return &ClubUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClubClient) UpdateOne(cl *Club) *ClubUpdateOne {
	mutation := newClubMutation(c.config, OpUpdateOne, withClub(cl))
	return &ClubUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClubClient) UpdateOneID(id int) *ClubUpdateOne {
	mutation := newClubMutation(c.config, OpUpdateOne, withClubID(id))
	return &ClubUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Club.
func (c *ClubClient) Delete() *ClubDelete {
	mutation := newClubMutation(c.config, OpDelete)
	return &ClubDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ClubClient) DeleteOne(cl *Club) *ClubDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ClubClient) DeleteOneID(id int) *ClubDeleteOne {
	builder := c.Delete().Where(club.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClubDeleteOne{builder}
}

// Create returns a query builder for Club.
func (c *ClubClient) Query() *ClubQuery {
	return &ClubQuery{config: c.config}
}

// Get returns a Club entity by its id.
func (c *ClubClient) Get(ctx context.Context, id int) (*Club, error) {
	return c.Query().Where(club.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClubClient) GetX(ctx context.Context, id int) *Club {
	cl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return cl
}

// QueryLocation queries the location edge of a Club.
func (c *ClubClient) QueryLocation(cl *Club) *LocationQuery {
	query := &LocationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(club.Table, club.FieldID, id),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, club.LocationTable, club.LocationColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryClubtypes queries the clubtypes edge of a Club.
func (c *ClubClient) QueryClubtypes(cl *Club) *ClubTypesQuery {
	query := &ClubTypesQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(club.Table, club.FieldID, id),
			sqlgraph.To(clubtypes.Table, clubtypes.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, club.ClubtypesTable, club.ClubtypesColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Club.
func (c *ClubClient) QueryUser(cl *Club) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(club.Table, club.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, club.UserTable, club.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryActivity queries the activity edge of a Club.
func (c *ClubClient) QueryActivity(cl *Club) *ActivityQuery {
	query := &ActivityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(club.Table, club.FieldID, id),
			sqlgraph.To(activity.Table, activity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, club.ActivityTable, club.ActivityColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClubClient) Hooks() []Hook {
	return c.hooks.Club
}

// ClubTypesClient is a client for the ClubTypes schema.
type ClubTypesClient struct {
	config
}

// NewClubTypesClient returns a client for the ClubTypes from the given config.
func NewClubTypesClient(c config) *ClubTypesClient {
	return &ClubTypesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `clubtypes.Hooks(f(g(h())))`.
func (c *ClubTypesClient) Use(hooks ...Hook) {
	c.hooks.ClubTypes = append(c.hooks.ClubTypes, hooks...)
}

// Create returns a create builder for ClubTypes.
func (c *ClubTypesClient) Create() *ClubTypesCreate {
	mutation := newClubTypesMutation(c.config, OpCreate)
	return &ClubTypesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for ClubTypes.
func (c *ClubTypesClient) Update() *ClubTypesUpdate {
	mutation := newClubTypesMutation(c.config, OpUpdate)
	return &ClubTypesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClubTypesClient) UpdateOne(ct *ClubTypes) *ClubTypesUpdateOne {
	mutation := newClubTypesMutation(c.config, OpUpdateOne, withClubTypes(ct))
	return &ClubTypesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClubTypesClient) UpdateOneID(id int) *ClubTypesUpdateOne {
	mutation := newClubTypesMutation(c.config, OpUpdateOne, withClubTypesID(id))
	return &ClubTypesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ClubTypes.
func (c *ClubTypesClient) Delete() *ClubTypesDelete {
	mutation := newClubTypesMutation(c.config, OpDelete)
	return &ClubTypesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ClubTypesClient) DeleteOne(ct *ClubTypes) *ClubTypesDeleteOne {
	return c.DeleteOneID(ct.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ClubTypesClient) DeleteOneID(id int) *ClubTypesDeleteOne {
	builder := c.Delete().Where(clubtypes.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClubTypesDeleteOne{builder}
}

// Create returns a query builder for ClubTypes.
func (c *ClubTypesClient) Query() *ClubTypesQuery {
	return &ClubTypesQuery{config: c.config}
}

// Get returns a ClubTypes entity by its id.
func (c *ClubTypesClient) Get(ctx context.Context, id int) (*ClubTypes, error) {
	return c.Query().Where(clubtypes.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClubTypesClient) GetX(ctx context.Context, id int) *ClubTypes {
	ct, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ct
}

// QueryClub queries the club edge of a ClubTypes.
func (c *ClubTypesClient) QueryClub(ct *ClubTypes) *ClubQuery {
	query := &ClubQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ct.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(clubtypes.Table, clubtypes.FieldID, id),
			sqlgraph.To(club.Table, club.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clubtypes.ClubTable, clubtypes.ClubColumn),
		)
		fromV = sqlgraph.Neighbors(ct.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClubTypesClient) Hooks() []Hook {
	return c.hooks.ClubTypes
}

// LocationClient is a client for the Location schema.
type LocationClient struct {
	config
}

// NewLocationClient returns a client for the Location from the given config.
func NewLocationClient(c config) *LocationClient {
	return &LocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `location.Hooks(f(g(h())))`.
func (c *LocationClient) Use(hooks ...Hook) {
	c.hooks.Location = append(c.hooks.Location, hooks...)
}

// Create returns a create builder for Location.
func (c *LocationClient) Create() *LocationCreate {
	mutation := newLocationMutation(c.config, OpCreate)
	return &LocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Location.
func (c *LocationClient) Update() *LocationUpdate {
	mutation := newLocationMutation(c.config, OpUpdate)
	return &LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LocationClient) UpdateOne(l *Location) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocation(l))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LocationClient) UpdateOneID(id int) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocationID(id))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Location.
func (c *LocationClient) Delete() *LocationDelete {
	mutation := newLocationMutation(c.config, OpDelete)
	return &LocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LocationClient) DeleteOne(l *Location) *LocationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LocationClient) DeleteOneID(id int) *LocationDeleteOne {
	builder := c.Delete().Where(location.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LocationDeleteOne{builder}
}

// Create returns a query builder for Location.
func (c *LocationClient) Query() *LocationQuery {
	return &LocationQuery{config: c.config}
}

// Get returns a Location entity by its id.
func (c *LocationClient) Get(ctx context.Context, id int) (*Location, error) {
	return c.Query().Where(location.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LocationClient) GetX(ctx context.Context, id int) *Location {
	l, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return l
}

// QueryClub queries the club edge of a Location.
func (c *LocationClient) QueryClub(l *Location) *ClubQuery {
	query := &ClubQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, id),
			sqlgraph.To(club.Table, club.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, location.ClubTable, location.ClubColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LocationClient) Hooks() []Hook {
	return c.hooks.Location
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryClub queries the club edge of a User.
func (c *UserClient) QueryClub(u *User) *ClubQuery {
	query := &ClubQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(club.Table, club.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ClubTable, user.ClubColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
