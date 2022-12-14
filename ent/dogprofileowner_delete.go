// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofileowner"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/predicate"
)

// DogProfileOwnerDelete is the builder for deleting a DogProfileOwner entity.
type DogProfileOwnerDelete struct {
	config
	hooks    []Hook
	mutation *DogProfileOwnerMutation
}

// Where appends a list predicates to the DogProfileOwnerDelete builder.
func (dpod *DogProfileOwnerDelete) Where(ps ...predicate.DogProfileOwner) *DogProfileOwnerDelete {
	dpod.mutation.Where(ps...)
	return dpod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dpod *DogProfileOwnerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dpod.hooks) == 0 {
		affected, err = dpod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DogProfileOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dpod.mutation = mutation
			affected, err = dpod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dpod.hooks) - 1; i >= 0; i-- {
			if dpod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpod *DogProfileOwnerDelete) ExecX(ctx context.Context) int {
	n, err := dpod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dpod *DogProfileOwnerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dogprofileowner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dogprofileowner.FieldID,
			},
		},
	}
	if ps := dpod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dpod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DogProfileOwnerDeleteOne is the builder for deleting a single DogProfileOwner entity.
type DogProfileOwnerDeleteOne struct {
	dpod *DogProfileOwnerDelete
}

// Exec executes the deletion query.
func (dpodo *DogProfileOwnerDeleteOne) Exec(ctx context.Context) error {
	n, err := dpodo.dpod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dogprofileowner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dpodo *DogProfileOwnerDeleteOne) ExecX(ctx context.Context) {
	dpodo.dpod.ExecX(ctx)
}
