// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dog"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofilebreed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofileowner"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/image"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/predicate"
	"github.com/google/uuid"
)

// DogQuery is the builder for querying Dog entities.
type DogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Dog
	// eager-loading edges.
	withImage         *ImageQuery
	withOwnerProfiles *DogProfileOwnerQuery
	withBreedProfiles *DogProfileBreedQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Dog) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DogQuery builder.
func (dq *DogQuery) Where(ps ...predicate.Dog) *DogQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DogQuery) Limit(limit int) *DogQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DogQuery) Offset(offset int) *DogQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DogQuery) Unique(unique bool) *DogQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DogQuery) Order(o ...OrderFunc) *DogQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryImage chains the current query on the "image" edge.
func (dq *DogQuery) QueryImage() *ImageQuery {
	query := &ImageQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dog.Table, dog.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dog.ImageTable, dog.ImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwnerProfiles chains the current query on the "ownerProfiles" edge.
func (dq *DogQuery) QueryOwnerProfiles() *DogProfileOwnerQuery {
	query := &DogProfileOwnerQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dog.Table, dog.FieldID, selector),
			sqlgraph.To(dogprofileowner.Table, dogprofileowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dog.OwnerProfilesTable, dog.OwnerProfilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBreedProfiles chains the current query on the "breedProfiles" edge.
func (dq *DogQuery) QueryBreedProfiles() *DogProfileBreedQuery {
	query := &DogProfileBreedQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dog.Table, dog.FieldID, selector),
			sqlgraph.To(dogprofilebreed.Table, dogprofilebreed.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dog.BreedProfilesTable, dog.BreedProfilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dog entity from the query.
// Returns a *NotFoundError when no Dog was found.
func (dq *DogQuery) First(ctx context.Context) (*Dog, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DogQuery) FirstX(ctx context.Context) *Dog {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dog ID from the query.
// Returns a *NotFoundError when no Dog ID was found.
func (dq *DogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dog entity is found.
// Returns a *NotFoundError when no Dog entities are found.
func (dq *DogQuery) Only(ctx context.Context) (*Dog, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dog.Label}
	default:
		return nil, &NotSingularError{dog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DogQuery) OnlyX(ctx context.Context) *Dog {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dog ID in the query.
// Returns a *NotSingularError when more than one Dog ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dog.Label}
	default:
		err = &NotSingularError{dog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dogs.
func (dq *DogQuery) All(ctx context.Context) ([]*Dog, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DogQuery) AllX(ctx context.Context) []*Dog {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dog IDs.
func (dq *DogQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dq.Select(dog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DogQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DogQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DogQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DogQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DogQuery) Clone() *DogQuery {
	if dq == nil {
		return nil
	}
	return &DogQuery{
		config:            dq.config,
		limit:             dq.limit,
		offset:            dq.offset,
		order:             append([]OrderFunc{}, dq.order...),
		predicates:        append([]predicate.Dog{}, dq.predicates...),
		withImage:         dq.withImage.Clone(),
		withOwnerProfiles: dq.withOwnerProfiles.Clone(),
		withBreedProfiles: dq.withBreedProfiles.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithImage tells the query-builder to eager-load the nodes that are connected to
// the "image" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DogQuery) WithImage(opts ...func(*ImageQuery)) *DogQuery {
	query := &ImageQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withImage = query
	return dq
}

// WithOwnerProfiles tells the query-builder to eager-load the nodes that are connected to
// the "ownerProfiles" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DogQuery) WithOwnerProfiles(opts ...func(*DogProfileOwnerQuery)) *DogQuery {
	query := &DogProfileOwnerQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withOwnerProfiles = query
	return dq
}

// WithBreedProfiles tells the query-builder to eager-load the nodes that are connected to
// the "breedProfiles" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DogQuery) WithBreedProfiles(opts ...func(*DogProfileBreedQuery)) *DogQuery {
	query := &DogProfileBreedQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withBreedProfiles = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FullName string `json:"full_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Dog.Query().
//		GroupBy(dog.FieldFullName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DogQuery) GroupBy(field string, fields ...string) *DogGroupBy {
	grbuild := &DogGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = dog.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FullName string `json:"full_name,omitempty"`
//	}
//
//	client.Dog.Query().
//		Select(dog.FieldFullName).
//		Scan(ctx, &v)
func (dq *DogQuery) Select(fields ...string) *DogSelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DogSelect{DogQuery: dq}
	selbuild.label = dog.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

func (dq *DogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !dog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dog, error) {
	var (
		nodes       = []*Dog{}
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withImage != nil,
			dq.withOwnerProfiles != nil,
			dq.withBreedProfiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Dog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Dog{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withImage; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Dog)
		for i := range nodes {
			fk := nodes[i].DogImgID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(image.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dog_img_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Image = n
			}
		}
	}

	if query := dq.withOwnerProfiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Dog)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OwnerProfiles = []*DogProfileOwner{}
		}
		query.Where(predicate.DogProfileOwner(func(s *sql.Selector) {
			s.Where(sql.InValues(dog.OwnerProfilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.DogID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dog_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.OwnerProfiles = append(node.Edges.OwnerProfiles, n)
		}
	}

	if query := dq.withBreedProfiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Dog)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BreedProfiles = []*DogProfileBreed{}
		}
		query.Where(predicate.DogProfileBreed(func(s *sql.Selector) {
			s.Where(sql.InValues(dog.BreedProfilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.DogID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dog_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.BreedProfiles = append(node.Edges.BreedProfiles, n)
		}
	}

	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dog.Table,
			Columns: dog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dog.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dog.FieldID)
		for i := range fields {
			if fields[i] != dog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dog.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = dog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DogGroupBy is the group-by builder for Dog entities.
type DogGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DogGroupBy) Aggregate(fns ...AggregateFunc) *DogGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !dog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DogGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DogSelect is the builder for selecting fields of Dog entities.
type DogSelect struct {
	*DogQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DogQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
