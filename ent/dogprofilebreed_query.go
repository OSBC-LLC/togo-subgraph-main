// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/breed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dog"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofilebreed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/predicate"
	"github.com/google/uuid"
)

// DogProfileBreedQuery is the builder for querying DogProfileBreed entities.
type DogProfileBreedQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DogProfileBreed
	// eager-loading edges.
	withDog   *DogQuery
	withBreed *BreedQuery
	modifiers []func(*sql.Selector)
	loadTotal []func(context.Context, []*DogProfileBreed) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DogProfileBreedQuery builder.
func (dpbq *DogProfileBreedQuery) Where(ps ...predicate.DogProfileBreed) *DogProfileBreedQuery {
	dpbq.predicates = append(dpbq.predicates, ps...)
	return dpbq
}

// Limit adds a limit step to the query.
func (dpbq *DogProfileBreedQuery) Limit(limit int) *DogProfileBreedQuery {
	dpbq.limit = &limit
	return dpbq
}

// Offset adds an offset step to the query.
func (dpbq *DogProfileBreedQuery) Offset(offset int) *DogProfileBreedQuery {
	dpbq.offset = &offset
	return dpbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dpbq *DogProfileBreedQuery) Unique(unique bool) *DogProfileBreedQuery {
	dpbq.unique = &unique
	return dpbq
}

// Order adds an order step to the query.
func (dpbq *DogProfileBreedQuery) Order(o ...OrderFunc) *DogProfileBreedQuery {
	dpbq.order = append(dpbq.order, o...)
	return dpbq
}

// QueryDog chains the current query on the "dog" edge.
func (dpbq *DogProfileBreedQuery) QueryDog() *DogQuery {
	query := &DogQuery{config: dpbq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dpbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dpbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dogprofilebreed.Table, dogprofilebreed.FieldID, selector),
			sqlgraph.To(dog.Table, dog.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dogprofilebreed.DogTable, dogprofilebreed.DogColumn),
		)
		fromU = sqlgraph.SetNeighbors(dpbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBreed chains the current query on the "breed" edge.
func (dpbq *DogProfileBreedQuery) QueryBreed() *BreedQuery {
	query := &BreedQuery{config: dpbq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dpbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dpbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dogprofilebreed.Table, dogprofilebreed.FieldID, selector),
			sqlgraph.To(breed.Table, breed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dogprofilebreed.BreedTable, dogprofilebreed.BreedColumn),
		)
		fromU = sqlgraph.SetNeighbors(dpbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DogProfileBreed entity from the query.
// Returns a *NotFoundError when no DogProfileBreed was found.
func (dpbq *DogProfileBreedQuery) First(ctx context.Context) (*DogProfileBreed, error) {
	nodes, err := dpbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dogprofilebreed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) FirstX(ctx context.Context) *DogProfileBreed {
	node, err := dpbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DogProfileBreed ID from the query.
// Returns a *NotFoundError when no DogProfileBreed ID was found.
func (dpbq *DogProfileBreedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dpbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dogprofilebreed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dpbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DogProfileBreed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DogProfileBreed entity is found.
// Returns a *NotFoundError when no DogProfileBreed entities are found.
func (dpbq *DogProfileBreedQuery) Only(ctx context.Context) (*DogProfileBreed, error) {
	nodes, err := dpbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dogprofilebreed.Label}
	default:
		return nil, &NotSingularError{dogprofilebreed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) OnlyX(ctx context.Context) *DogProfileBreed {
	node, err := dpbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DogProfileBreed ID in the query.
// Returns a *NotSingularError when more than one DogProfileBreed ID is found.
// Returns a *NotFoundError when no entities are found.
func (dpbq *DogProfileBreedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dpbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dogprofilebreed.Label}
	default:
		err = &NotSingularError{dogprofilebreed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dpbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DogProfileBreeds.
func (dpbq *DogProfileBreedQuery) All(ctx context.Context) ([]*DogProfileBreed, error) {
	if err := dpbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dpbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) AllX(ctx context.Context) []*DogProfileBreed {
	nodes, err := dpbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DogProfileBreed IDs.
func (dpbq *DogProfileBreedQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dpbq.Select(dogprofilebreed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dpbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dpbq *DogProfileBreedQuery) Count(ctx context.Context) (int, error) {
	if err := dpbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dpbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) CountX(ctx context.Context) int {
	count, err := dpbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dpbq *DogProfileBreedQuery) Exist(ctx context.Context) (bool, error) {
	if err := dpbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dpbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dpbq *DogProfileBreedQuery) ExistX(ctx context.Context) bool {
	exist, err := dpbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DogProfileBreedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dpbq *DogProfileBreedQuery) Clone() *DogProfileBreedQuery {
	if dpbq == nil {
		return nil
	}
	return &DogProfileBreedQuery{
		config:     dpbq.config,
		limit:      dpbq.limit,
		offset:     dpbq.offset,
		order:      append([]OrderFunc{}, dpbq.order...),
		predicates: append([]predicate.DogProfileBreed{}, dpbq.predicates...),
		withDog:    dpbq.withDog.Clone(),
		withBreed:  dpbq.withBreed.Clone(),
		// clone intermediate query.
		sql:    dpbq.sql.Clone(),
		path:   dpbq.path,
		unique: dpbq.unique,
	}
}

// WithDog tells the query-builder to eager-load the nodes that are connected to
// the "dog" edge. The optional arguments are used to configure the query builder of the edge.
func (dpbq *DogProfileBreedQuery) WithDog(opts ...func(*DogQuery)) *DogProfileBreedQuery {
	query := &DogQuery{config: dpbq.config}
	for _, opt := range opts {
		opt(query)
	}
	dpbq.withDog = query
	return dpbq
}

// WithBreed tells the query-builder to eager-load the nodes that are connected to
// the "breed" edge. The optional arguments are used to configure the query builder of the edge.
func (dpbq *DogProfileBreedQuery) WithBreed(opts ...func(*BreedQuery)) *DogProfileBreedQuery {
	query := &BreedQuery{config: dpbq.config}
	for _, opt := range opts {
		opt(query)
	}
	dpbq.withBreed = query
	return dpbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BreedID uuid.UUID `json:"breed_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DogProfileBreed.Query().
//		GroupBy(dogprofilebreed.FieldBreedID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dpbq *DogProfileBreedQuery) GroupBy(field string, fields ...string) *DogProfileBreedGroupBy {
	grbuild := &DogProfileBreedGroupBy{config: dpbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dpbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dpbq.sqlQuery(ctx), nil
	}
	grbuild.label = dogprofilebreed.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BreedID uuid.UUID `json:"breed_id,omitempty"`
//	}
//
//	client.DogProfileBreed.Query().
//		Select(dogprofilebreed.FieldBreedID).
//		Scan(ctx, &v)
func (dpbq *DogProfileBreedQuery) Select(fields ...string) *DogProfileBreedSelect {
	dpbq.fields = append(dpbq.fields, fields...)
	selbuild := &DogProfileBreedSelect{DogProfileBreedQuery: dpbq}
	selbuild.label = dogprofilebreed.Label
	selbuild.flds, selbuild.scan = &dpbq.fields, selbuild.Scan
	return selbuild
}

func (dpbq *DogProfileBreedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dpbq.fields {
		if !dogprofilebreed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dpbq.path != nil {
		prev, err := dpbq.path(ctx)
		if err != nil {
			return err
		}
		dpbq.sql = prev
	}
	return nil
}

func (dpbq *DogProfileBreedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DogProfileBreed, error) {
	var (
		nodes       = []*DogProfileBreed{}
		_spec       = dpbq.querySpec()
		loadedTypes = [2]bool{
			dpbq.withDog != nil,
			dpbq.withBreed != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DogProfileBreed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DogProfileBreed{config: dpbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dpbq.modifiers) > 0 {
		_spec.Modifiers = dpbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dpbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dpbq.withDog; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*DogProfileBreed)
		for i := range nodes {
			fk := nodes[i].DogID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(dog.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dog_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Dog = n
			}
		}
	}

	if query := dpbq.withBreed; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*DogProfileBreed)
		for i := range nodes {
			fk := nodes[i].BreedID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(breed.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "breed_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Breed = n
			}
		}
	}

	for i := range dpbq.loadTotal {
		if err := dpbq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dpbq *DogProfileBreedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dpbq.querySpec()
	if len(dpbq.modifiers) > 0 {
		_spec.Modifiers = dpbq.modifiers
	}
	_spec.Node.Columns = dpbq.fields
	if len(dpbq.fields) > 0 {
		_spec.Unique = dpbq.unique != nil && *dpbq.unique
	}
	return sqlgraph.CountNodes(ctx, dpbq.driver, _spec)
}

func (dpbq *DogProfileBreedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dpbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dpbq *DogProfileBreedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dogprofilebreed.Table,
			Columns: dogprofilebreed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dogprofilebreed.FieldID,
			},
		},
		From:   dpbq.sql,
		Unique: true,
	}
	if unique := dpbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dpbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dogprofilebreed.FieldID)
		for i := range fields {
			if fields[i] != dogprofilebreed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dpbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dpbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dpbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dpbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dpbq *DogProfileBreedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dpbq.driver.Dialect())
	t1 := builder.Table(dogprofilebreed.Table)
	columns := dpbq.fields
	if len(columns) == 0 {
		columns = dogprofilebreed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dpbq.sql != nil {
		selector = dpbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dpbq.unique != nil && *dpbq.unique {
		selector.Distinct()
	}
	for _, p := range dpbq.predicates {
		p(selector)
	}
	for _, p := range dpbq.order {
		p(selector)
	}
	if offset := dpbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dpbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DogProfileBreedGroupBy is the group-by builder for DogProfileBreed entities.
type DogProfileBreedGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dpbgb *DogProfileBreedGroupBy) Aggregate(fns ...AggregateFunc) *DogProfileBreedGroupBy {
	dpbgb.fns = append(dpbgb.fns, fns...)
	return dpbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dpbgb *DogProfileBreedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dpbgb.path(ctx)
	if err != nil {
		return err
	}
	dpbgb.sql = query
	return dpbgb.sqlScan(ctx, v)
}

func (dpbgb *DogProfileBreedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dpbgb.fields {
		if !dogprofilebreed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dpbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dpbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dpbgb *DogProfileBreedGroupBy) sqlQuery() *sql.Selector {
	selector := dpbgb.sql.Select()
	aggregation := make([]string, 0, len(dpbgb.fns))
	for _, fn := range dpbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dpbgb.fields)+len(dpbgb.fns))
		for _, f := range dpbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dpbgb.fields...)...)
}

// DogProfileBreedSelect is the builder for selecting fields of DogProfileBreed entities.
type DogProfileBreedSelect struct {
	*DogProfileBreedQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dpbs *DogProfileBreedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dpbs.prepareQuery(ctx); err != nil {
		return err
	}
	dpbs.sql = dpbs.DogProfileBreedQuery.sqlQuery(ctx)
	return dpbs.sqlScan(ctx, v)
}

func (dpbs *DogProfileBreedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dpbs.sql.Query()
	if err := dpbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
