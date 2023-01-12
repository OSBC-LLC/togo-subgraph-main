// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/OSBC-LLC/togo-subgraph-main/ent/migrate"
	"github.com/google/uuid"

	"github.com/OSBC-LLC/togo-subgraph-main/ent/breed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dog"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofilebreed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofileowner"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/image"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/profile"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Breed is the client for interacting with the Breed builders.
	Breed *BreedClient
	// Dog is the client for interacting with the Dog builders.
	Dog *DogClient
	// DogProfileBreed is the client for interacting with the DogProfileBreed builders.
	DogProfileBreed *DogProfileBreedClient
	// DogProfileOwner is the client for interacting with the DogProfileOwner builders.
	DogProfileOwner *DogProfileOwnerClient
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
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
	c.Breed = NewBreedClient(c.config)
	c.Dog = NewDogClient(c.config)
	c.DogProfileBreed = NewDogProfileBreedClient(c.config)
	c.DogProfileOwner = NewDogProfileOwnerClient(c.config)
	c.Image = NewImageClient(c.config)
	c.Profile = NewProfileClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Breed:           NewBreedClient(cfg),
		Dog:             NewDogClient(cfg),
		DogProfileBreed: NewDogProfileBreedClient(cfg),
		DogProfileOwner: NewDogProfileOwnerClient(cfg),
		Image:           NewImageClient(cfg),
		Profile:         NewProfileClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:             ctx,
		config:          cfg,
		Breed:           NewBreedClient(cfg),
		Dog:             NewDogClient(cfg),
		DogProfileBreed: NewDogProfileBreedClient(cfg),
		DogProfileOwner: NewDogProfileOwnerClient(cfg),
		Image:           NewImageClient(cfg),
		Profile:         NewProfileClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Breed.
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
	c.Breed.Use(hooks...)
	c.Dog.Use(hooks...)
	c.DogProfileBreed.Use(hooks...)
	c.DogProfileOwner.Use(hooks...)
	c.Image.Use(hooks...)
	c.Profile.Use(hooks...)
	c.User.Use(hooks...)
}

// BreedClient is a client for the Breed schema.
type BreedClient struct {
	config
}

// NewBreedClient returns a client for the Breed from the given config.
func NewBreedClient(c config) *BreedClient {
	return &BreedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `breed.Hooks(f(g(h())))`.
func (c *BreedClient) Use(hooks ...Hook) {
	c.hooks.Breed = append(c.hooks.Breed, hooks...)
}

// Create returns a builder for creating a Breed entity.
func (c *BreedClient) Create() *BreedCreate {
	mutation := newBreedMutation(c.config, OpCreate)
	return &BreedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Breed entities.
func (c *BreedClient) CreateBulk(builders ...*BreedCreate) *BreedCreateBulk {
	return &BreedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Breed.
func (c *BreedClient) Update() *BreedUpdate {
	mutation := newBreedMutation(c.config, OpUpdate)
	return &BreedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BreedClient) UpdateOne(b *Breed) *BreedUpdateOne {
	mutation := newBreedMutation(c.config, OpUpdateOne, withBreed(b))
	return &BreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BreedClient) UpdateOneID(id uuid.UUID) *BreedUpdateOne {
	mutation := newBreedMutation(c.config, OpUpdateOne, withBreedID(id))
	return &BreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Breed.
func (c *BreedClient) Delete() *BreedDelete {
	mutation := newBreedMutation(c.config, OpDelete)
	return &BreedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BreedClient) DeleteOne(b *Breed) *BreedDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BreedClient) DeleteOneID(id uuid.UUID) *BreedDeleteOne {
	builder := c.Delete().Where(breed.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BreedDeleteOne{builder}
}

// Query returns a query builder for Breed.
func (c *BreedClient) Query() *BreedQuery {
	return &BreedQuery{
		config: c.config,
	}
}

// Get returns a Breed entity by its id.
func (c *BreedClient) Get(ctx context.Context, id uuid.UUID) (*Breed, error) {
	return c.Query().Where(breed.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BreedClient) GetX(ctx context.Context, id uuid.UUID) *Breed {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BreedClient) Hooks() []Hook {
	return c.hooks.Breed
}

// DogClient is a client for the Dog schema.
type DogClient struct {
	config
}

// NewDogClient returns a client for the Dog from the given config.
func NewDogClient(c config) *DogClient {
	return &DogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dog.Hooks(f(g(h())))`.
func (c *DogClient) Use(hooks ...Hook) {
	c.hooks.Dog = append(c.hooks.Dog, hooks...)
}

// Create returns a builder for creating a Dog entity.
func (c *DogClient) Create() *DogCreate {
	mutation := newDogMutation(c.config, OpCreate)
	return &DogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dog entities.
func (c *DogClient) CreateBulk(builders ...*DogCreate) *DogCreateBulk {
	return &DogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dog.
func (c *DogClient) Update() *DogUpdate {
	mutation := newDogMutation(c.config, OpUpdate)
	return &DogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DogClient) UpdateOne(d *Dog) *DogUpdateOne {
	mutation := newDogMutation(c.config, OpUpdateOne, withDog(d))
	return &DogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DogClient) UpdateOneID(id uuid.UUID) *DogUpdateOne {
	mutation := newDogMutation(c.config, OpUpdateOne, withDogID(id))
	return &DogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dog.
func (c *DogClient) Delete() *DogDelete {
	mutation := newDogMutation(c.config, OpDelete)
	return &DogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DogClient) DeleteOne(d *Dog) *DogDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DogClient) DeleteOneID(id uuid.UUID) *DogDeleteOne {
	builder := c.Delete().Where(dog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DogDeleteOne{builder}
}

// Query returns a query builder for Dog.
func (c *DogClient) Query() *DogQuery {
	return &DogQuery{
		config: c.config,
	}
}

// Get returns a Dog entity by its id.
func (c *DogClient) Get(ctx context.Context, id uuid.UUID) (*Dog, error) {
	return c.Query().Where(dog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DogClient) GetX(ctx context.Context, id uuid.UUID) *Dog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwnerProfiles queries the ownerProfiles edge of a Dog.
func (c *DogClient) QueryOwnerProfiles(d *Dog) *DogProfileOwnerQuery {
	query := &DogProfileOwnerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dog.Table, dog.FieldID, id),
			sqlgraph.To(dogprofileowner.Table, dogprofileowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dog.OwnerProfilesTable, dog.OwnerProfilesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DogClient) Hooks() []Hook {
	return c.hooks.Dog
}

// DogProfileBreedClient is a client for the DogProfileBreed schema.
type DogProfileBreedClient struct {
	config
}

// NewDogProfileBreedClient returns a client for the DogProfileBreed from the given config.
func NewDogProfileBreedClient(c config) *DogProfileBreedClient {
	return &DogProfileBreedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dogprofilebreed.Hooks(f(g(h())))`.
func (c *DogProfileBreedClient) Use(hooks ...Hook) {
	c.hooks.DogProfileBreed = append(c.hooks.DogProfileBreed, hooks...)
}

// Create returns a builder for creating a DogProfileBreed entity.
func (c *DogProfileBreedClient) Create() *DogProfileBreedCreate {
	mutation := newDogProfileBreedMutation(c.config, OpCreate)
	return &DogProfileBreedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DogProfileBreed entities.
func (c *DogProfileBreedClient) CreateBulk(builders ...*DogProfileBreedCreate) *DogProfileBreedCreateBulk {
	return &DogProfileBreedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DogProfileBreed.
func (c *DogProfileBreedClient) Update() *DogProfileBreedUpdate {
	mutation := newDogProfileBreedMutation(c.config, OpUpdate)
	return &DogProfileBreedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DogProfileBreedClient) UpdateOne(dpb *DogProfileBreed) *DogProfileBreedUpdateOne {
	mutation := newDogProfileBreedMutation(c.config, OpUpdateOne, withDogProfileBreed(dpb))
	return &DogProfileBreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DogProfileBreedClient) UpdateOneID(id uuid.UUID) *DogProfileBreedUpdateOne {
	mutation := newDogProfileBreedMutation(c.config, OpUpdateOne, withDogProfileBreedID(id))
	return &DogProfileBreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DogProfileBreed.
func (c *DogProfileBreedClient) Delete() *DogProfileBreedDelete {
	mutation := newDogProfileBreedMutation(c.config, OpDelete)
	return &DogProfileBreedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DogProfileBreedClient) DeleteOne(dpb *DogProfileBreed) *DogProfileBreedDeleteOne {
	return c.DeleteOneID(dpb.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DogProfileBreedClient) DeleteOneID(id uuid.UUID) *DogProfileBreedDeleteOne {
	builder := c.Delete().Where(dogprofilebreed.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DogProfileBreedDeleteOne{builder}
}

// Query returns a query builder for DogProfileBreed.
func (c *DogProfileBreedClient) Query() *DogProfileBreedQuery {
	return &DogProfileBreedQuery{
		config: c.config,
	}
}

// Get returns a DogProfileBreed entity by its id.
func (c *DogProfileBreedClient) Get(ctx context.Context, id uuid.UUID) (*DogProfileBreed, error) {
	return c.Query().Where(dogprofilebreed.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DogProfileBreedClient) GetX(ctx context.Context, id uuid.UUID) *DogProfileBreed {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DogProfileBreedClient) Hooks() []Hook {
	return c.hooks.DogProfileBreed
}

// DogProfileOwnerClient is a client for the DogProfileOwner schema.
type DogProfileOwnerClient struct {
	config
}

// NewDogProfileOwnerClient returns a client for the DogProfileOwner from the given config.
func NewDogProfileOwnerClient(c config) *DogProfileOwnerClient {
	return &DogProfileOwnerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dogprofileowner.Hooks(f(g(h())))`.
func (c *DogProfileOwnerClient) Use(hooks ...Hook) {
	c.hooks.DogProfileOwner = append(c.hooks.DogProfileOwner, hooks...)
}

// Create returns a builder for creating a DogProfileOwner entity.
func (c *DogProfileOwnerClient) Create() *DogProfileOwnerCreate {
	mutation := newDogProfileOwnerMutation(c.config, OpCreate)
	return &DogProfileOwnerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DogProfileOwner entities.
func (c *DogProfileOwnerClient) CreateBulk(builders ...*DogProfileOwnerCreate) *DogProfileOwnerCreateBulk {
	return &DogProfileOwnerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DogProfileOwner.
func (c *DogProfileOwnerClient) Update() *DogProfileOwnerUpdate {
	mutation := newDogProfileOwnerMutation(c.config, OpUpdate)
	return &DogProfileOwnerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DogProfileOwnerClient) UpdateOne(dpo *DogProfileOwner) *DogProfileOwnerUpdateOne {
	mutation := newDogProfileOwnerMutation(c.config, OpUpdateOne, withDogProfileOwner(dpo))
	return &DogProfileOwnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DogProfileOwnerClient) UpdateOneID(id uuid.UUID) *DogProfileOwnerUpdateOne {
	mutation := newDogProfileOwnerMutation(c.config, OpUpdateOne, withDogProfileOwnerID(id))
	return &DogProfileOwnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DogProfileOwner.
func (c *DogProfileOwnerClient) Delete() *DogProfileOwnerDelete {
	mutation := newDogProfileOwnerMutation(c.config, OpDelete)
	return &DogProfileOwnerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DogProfileOwnerClient) DeleteOne(dpo *DogProfileOwner) *DogProfileOwnerDeleteOne {
	return c.DeleteOneID(dpo.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DogProfileOwnerClient) DeleteOneID(id uuid.UUID) *DogProfileOwnerDeleteOne {
	builder := c.Delete().Where(dogprofileowner.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DogProfileOwnerDeleteOne{builder}
}

// Query returns a query builder for DogProfileOwner.
func (c *DogProfileOwnerClient) Query() *DogProfileOwnerQuery {
	return &DogProfileOwnerQuery{
		config: c.config,
	}
}

// Get returns a DogProfileOwner entity by its id.
func (c *DogProfileOwnerClient) Get(ctx context.Context, id uuid.UUID) (*DogProfileOwner, error) {
	return c.Query().Where(dogprofileowner.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DogProfileOwnerClient) GetX(ctx context.Context, id uuid.UUID) *DogProfileOwner {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a DogProfileOwner.
func (c *DogProfileOwnerClient) QueryOwner(dpo *DogProfileOwner) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dpo.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dogprofileowner.Table, dogprofileowner.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dogprofileowner.OwnerTable, dogprofileowner.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(dpo.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDog queries the dog edge of a DogProfileOwner.
func (c *DogProfileOwnerClient) QueryDog(dpo *DogProfileOwner) *DogQuery {
	query := &DogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dpo.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dogprofileowner.Table, dogprofileowner.FieldID, id),
			sqlgraph.To(dog.Table, dog.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dogprofileowner.DogTable, dogprofileowner.DogColumn),
		)
		fromV = sqlgraph.Neighbors(dpo.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DogProfileOwnerClient) Hooks() []Hook {
	return c.hooks.DogProfileOwner
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Create returns a builder for creating a Image entity.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id uuid.UUID) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ImageClient) DeleteOneID(id uuid.UUID) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id uuid.UUID) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id uuid.UUID) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// ProfileClient is a client for the Profile schema.
type ProfileClient struct {
	config
}

// NewProfileClient returns a client for the Profile from the given config.
func NewProfileClient(c config) *ProfileClient {
	return &ProfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profile.Hooks(f(g(h())))`.
func (c *ProfileClient) Use(hooks ...Hook) {
	c.hooks.Profile = append(c.hooks.Profile, hooks...)
}

// Create returns a builder for creating a Profile entity.
func (c *ProfileClient) Create() *ProfileCreate {
	mutation := newProfileMutation(c.config, OpCreate)
	return &ProfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profile entities.
func (c *ProfileClient) CreateBulk(builders ...*ProfileCreate) *ProfileCreateBulk {
	return &ProfileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profile.
func (c *ProfileClient) Update() *ProfileUpdate {
	mutation := newProfileMutation(c.config, OpUpdate)
	return &ProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfileClient) UpdateOne(pr *Profile) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfile(pr))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfileClient) UpdateOneID(id uuid.UUID) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfileID(id))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profile.
func (c *ProfileClient) Delete() *ProfileDelete {
	mutation := newProfileMutation(c.config, OpDelete)
	return &ProfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProfileClient) DeleteOne(pr *Profile) *ProfileDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ProfileClient) DeleteOneID(id uuid.UUID) *ProfileDeleteOne {
	builder := c.Delete().Where(profile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfileDeleteOne{builder}
}

// Query returns a query builder for Profile.
func (c *ProfileClient) Query() *ProfileQuery {
	return &ProfileQuery{
		config: c.config,
	}
}

// Get returns a Profile entity by its id.
func (c *ProfileClient) Get(ctx context.Context, id uuid.UUID) (*Profile, error) {
	return c.Query().Where(profile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfileClient) GetX(ctx context.Context, id uuid.UUID) *Profile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Profile.
func (c *ProfileClient) QueryUsers(pr *Profile) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, profile.UsersTable, profile.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProfileClient) Hooks() []Hook {
	return c.hooks.Profile
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

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProfile queries the profile edge of a User.
func (c *UserClient) QueryProfile(u *User) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.ProfileTable, user.ProfileColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDogProfiles queries the dogProfiles edge of a User.
func (c *UserClient) QueryDogProfiles(u *User) *DogProfileOwnerQuery {
	query := &DogProfileOwnerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(dogprofileowner.Table, dogprofileowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DogProfilesTable, user.DogProfilesColumn),
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
