// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/account"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/predicate"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/tennant"
	"github.com/google/uuid"
)

// TennantUpdate is the builder for updating Tennant entities.
type TennantUpdate struct {
	config
	hooks    []Hook
	mutation *TennantMutation
}

// Where appends a list predicates to the TennantUpdate builder.
func (tu *TennantUpdate) Where(ps ...predicate.Tennant) *TennantUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetExternalId sets the "externalId" field.
func (tu *TennantUpdate) SetExternalId(s string) *TennantUpdate {
	tu.mutation.SetExternalId(s)
	return tu
}

// SetCloud sets the "cloud" field.
func (tu *TennantUpdate) SetCloud(s string) *TennantUpdate {
	tu.mutation.SetCloud(s)
	return tu
}

// SetAccountID sets the "account_id" field.
func (tu *TennantUpdate) SetAccountID(u uuid.UUID) *TennantUpdate {
	tu.mutation.SetAccountID(u)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TennantUpdate) SetCreatedAt(t time.Time) *TennantUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TennantUpdate) SetNillableCreatedAt(t *time.Time) *TennantUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetAccount sets the "account" edge to the Account entity.
func (tu *TennantUpdate) SetAccount(a *Account) *TennantUpdate {
	return tu.SetAccountID(a.ID)
}

// Mutation returns the TennantMutation object of the builder.
func (tu *TennantUpdate) Mutation() *TennantMutation {
	return tu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (tu *TennantUpdate) ClearAccount() *TennantUpdate {
	tu.mutation.ClearAccount()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TennantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TennantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TennantUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TennantUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TennantUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TennantUpdate) check() error {
	if _, ok := tu.mutation.AccountID(); tu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tennant.account"`)
	}
	return nil
}

func (tu *TennantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tennant.Table,
			Columns: tennant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tennant.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.ExternalId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tennant.FieldExternalId,
		})
	}
	if value, ok := tu.mutation.Cloud(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tennant.FieldCloud,
		})
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tennant.FieldCreatedAt,
		})
	}
	if tu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tennant.AccountTable,
			Columns: []string{tennant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tennant.AccountTable,
			Columns: []string{tennant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tennant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TennantUpdateOne is the builder for updating a single Tennant entity.
type TennantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TennantMutation
}

// SetExternalId sets the "externalId" field.
func (tuo *TennantUpdateOne) SetExternalId(s string) *TennantUpdateOne {
	tuo.mutation.SetExternalId(s)
	return tuo
}

// SetCloud sets the "cloud" field.
func (tuo *TennantUpdateOne) SetCloud(s string) *TennantUpdateOne {
	tuo.mutation.SetCloud(s)
	return tuo
}

// SetAccountID sets the "account_id" field.
func (tuo *TennantUpdateOne) SetAccountID(u uuid.UUID) *TennantUpdateOne {
	tuo.mutation.SetAccountID(u)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TennantUpdateOne) SetCreatedAt(t time.Time) *TennantUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TennantUpdateOne) SetNillableCreatedAt(t *time.Time) *TennantUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetAccount sets the "account" edge to the Account entity.
func (tuo *TennantUpdateOne) SetAccount(a *Account) *TennantUpdateOne {
	return tuo.SetAccountID(a.ID)
}

// Mutation returns the TennantMutation object of the builder.
func (tuo *TennantUpdateOne) Mutation() *TennantMutation {
	return tuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (tuo *TennantUpdateOne) ClearAccount() *TennantUpdateOne {
	tuo.mutation.ClearAccount()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TennantUpdateOne) Select(field string, fields ...string) *TennantUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tennant entity.
func (tuo *TennantUpdateOne) Save(ctx context.Context) (*Tennant, error) {
	var (
		err  error
		node *Tennant
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TennantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tennant)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TennantMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TennantUpdateOne) SaveX(ctx context.Context) *Tennant {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TennantUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TennantUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TennantUpdateOne) check() error {
	if _, ok := tuo.mutation.AccountID(); tuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tennant.account"`)
	}
	return nil
}

func (tuo *TennantUpdateOne) sqlSave(ctx context.Context) (_node *Tennant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tennant.Table,
			Columns: tennant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tennant.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tennant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tennant.FieldID)
		for _, f := range fields {
			if !tennant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tennant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.ExternalId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tennant.FieldExternalId,
		})
	}
	if value, ok := tuo.mutation.Cloud(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tennant.FieldCloud,
		})
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tennant.FieldCreatedAt,
		})
	}
	if tuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tennant.AccountTable,
			Columns: []string{tennant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tennant.AccountTable,
			Columns: []string{tennant.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tennant{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tennant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
