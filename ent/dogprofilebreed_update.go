// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofilebreed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/predicate"
)

// DogProfileBreedUpdate is the builder for updating DogProfileBreed entities.
type DogProfileBreedUpdate struct {
	config
	hooks    []Hook
	mutation *DogProfileBreedMutation
}

// Where appends a list predicates to the DogProfileBreedUpdate builder.
func (dpbu *DogProfileBreedUpdate) Where(ps ...predicate.DogProfileBreed) *DogProfileBreedUpdate {
	dpbu.mutation.Where(ps...)
	return dpbu
}

// Mutation returns the DogProfileBreedMutation object of the builder.
func (dpbu *DogProfileBreedUpdate) Mutation() *DogProfileBreedMutation {
	return dpbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dpbu *DogProfileBreedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dpbu.hooks) == 0 {
		affected, err = dpbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DogProfileBreedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dpbu.mutation = mutation
			affected, err = dpbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dpbu.hooks) - 1; i >= 0; i-- {
			if dpbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dpbu *DogProfileBreedUpdate) SaveX(ctx context.Context) int {
	affected, err := dpbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dpbu *DogProfileBreedUpdate) Exec(ctx context.Context) error {
	_, err := dpbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpbu *DogProfileBreedUpdate) ExecX(ctx context.Context) {
	if err := dpbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dpbu *DogProfileBreedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dogprofilebreed.Table,
			Columns: dogprofilebreed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dogprofilebreed.FieldID,
			},
		},
	}
	if ps := dpbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dpbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dogprofilebreed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DogProfileBreedUpdateOne is the builder for updating a single DogProfileBreed entity.
type DogProfileBreedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DogProfileBreedMutation
}

// Mutation returns the DogProfileBreedMutation object of the builder.
func (dpbuo *DogProfileBreedUpdateOne) Mutation() *DogProfileBreedMutation {
	return dpbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dpbuo *DogProfileBreedUpdateOne) Select(field string, fields ...string) *DogProfileBreedUpdateOne {
	dpbuo.fields = append([]string{field}, fields...)
	return dpbuo
}

// Save executes the query and returns the updated DogProfileBreed entity.
func (dpbuo *DogProfileBreedUpdateOne) Save(ctx context.Context) (*DogProfileBreed, error) {
	var (
		err  error
		node *DogProfileBreed
	)
	if len(dpbuo.hooks) == 0 {
		node, err = dpbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DogProfileBreedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dpbuo.mutation = mutation
			node, err = dpbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dpbuo.hooks) - 1; i >= 0; i-- {
			if dpbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpbuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dpbuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DogProfileBreed)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DogProfileBreedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dpbuo *DogProfileBreedUpdateOne) SaveX(ctx context.Context) *DogProfileBreed {
	node, err := dpbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dpbuo *DogProfileBreedUpdateOne) Exec(ctx context.Context) error {
	_, err := dpbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpbuo *DogProfileBreedUpdateOne) ExecX(ctx context.Context) {
	if err := dpbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dpbuo *DogProfileBreedUpdateOne) sqlSave(ctx context.Context) (_node *DogProfileBreed, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dogprofilebreed.Table,
			Columns: dogprofilebreed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dogprofilebreed.FieldID,
			},
		},
	}
	id, ok := dpbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DogProfileBreed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dpbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dogprofilebreed.FieldID)
		for _, f := range fields {
			if !dogprofilebreed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dogprofilebreed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dpbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &DogProfileBreed{config: dpbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dpbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dogprofilebreed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
