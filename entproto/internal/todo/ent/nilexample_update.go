// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/todo/ent/nilexample"
	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NilExampleUpdate is the builder for updating NilExample entities.
type NilExampleUpdate struct {
	config
	hooks    []Hook
	mutation *NilExampleMutation
}

// Where appends a list predicates to the NilExampleUpdate builder.
func (neu *NilExampleUpdate) Where(ps ...predicate.NilExample) *NilExampleUpdate {
	neu.mutation.Where(ps...)
	return neu
}

// SetStrNil sets the "str_nil" field.
func (neu *NilExampleUpdate) SetStrNil(s string) *NilExampleUpdate {
	neu.mutation.SetStrNil(s)
	return neu
}

// SetNillableStrNil sets the "str_nil" field if the given value is not nil.
func (neu *NilExampleUpdate) SetNillableStrNil(s *string) *NilExampleUpdate {
	if s != nil {
		neu.SetStrNil(*s)
	}
	return neu
}

// ClearStrNil clears the value of the "str_nil" field.
func (neu *NilExampleUpdate) ClearStrNil() *NilExampleUpdate {
	neu.mutation.ClearStrNil()
	return neu
}

// SetTimeNil sets the "time_nil" field.
func (neu *NilExampleUpdate) SetTimeNil(t time.Time) *NilExampleUpdate {
	neu.mutation.SetTimeNil(t)
	return neu
}

// SetNillableTimeNil sets the "time_nil" field if the given value is not nil.
func (neu *NilExampleUpdate) SetNillableTimeNil(t *time.Time) *NilExampleUpdate {
	if t != nil {
		neu.SetTimeNil(*t)
	}
	return neu
}

// ClearTimeNil clears the value of the "time_nil" field.
func (neu *NilExampleUpdate) ClearTimeNil() *NilExampleUpdate {
	neu.mutation.ClearTimeNil()
	return neu
}

// Mutation returns the NilExampleMutation object of the builder.
func (neu *NilExampleUpdate) Mutation() *NilExampleMutation {
	return neu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (neu *NilExampleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(neu.hooks) == 0 {
		affected, err = neu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NilExampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			neu.mutation = mutation
			affected, err = neu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(neu.hooks) - 1; i >= 0; i-- {
			if neu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = neu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, neu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (neu *NilExampleUpdate) SaveX(ctx context.Context) int {
	affected, err := neu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (neu *NilExampleUpdate) Exec(ctx context.Context) error {
	_, err := neu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (neu *NilExampleUpdate) ExecX(ctx context.Context) {
	if err := neu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (neu *NilExampleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nilexample.Table,
			Columns: nilexample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nilexample.FieldID,
			},
		},
	}
	if ps := neu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := neu.mutation.StrNil(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nilexample.FieldStrNil,
		})
	}
	if neu.mutation.StrNilCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: nilexample.FieldStrNil,
		})
	}
	if value, ok := neu.mutation.TimeNil(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: nilexample.FieldTimeNil,
		})
	}
	if neu.mutation.TimeNilCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: nilexample.FieldTimeNil,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, neu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nilexample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NilExampleUpdateOne is the builder for updating a single NilExample entity.
type NilExampleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NilExampleMutation
}

// SetStrNil sets the "str_nil" field.
func (neuo *NilExampleUpdateOne) SetStrNil(s string) *NilExampleUpdateOne {
	neuo.mutation.SetStrNil(s)
	return neuo
}

// SetNillableStrNil sets the "str_nil" field if the given value is not nil.
func (neuo *NilExampleUpdateOne) SetNillableStrNil(s *string) *NilExampleUpdateOne {
	if s != nil {
		neuo.SetStrNil(*s)
	}
	return neuo
}

// ClearStrNil clears the value of the "str_nil" field.
func (neuo *NilExampleUpdateOne) ClearStrNil() *NilExampleUpdateOne {
	neuo.mutation.ClearStrNil()
	return neuo
}

// SetTimeNil sets the "time_nil" field.
func (neuo *NilExampleUpdateOne) SetTimeNil(t time.Time) *NilExampleUpdateOne {
	neuo.mutation.SetTimeNil(t)
	return neuo
}

// SetNillableTimeNil sets the "time_nil" field if the given value is not nil.
func (neuo *NilExampleUpdateOne) SetNillableTimeNil(t *time.Time) *NilExampleUpdateOne {
	if t != nil {
		neuo.SetTimeNil(*t)
	}
	return neuo
}

// ClearTimeNil clears the value of the "time_nil" field.
func (neuo *NilExampleUpdateOne) ClearTimeNil() *NilExampleUpdateOne {
	neuo.mutation.ClearTimeNil()
	return neuo
}

// Mutation returns the NilExampleMutation object of the builder.
func (neuo *NilExampleUpdateOne) Mutation() *NilExampleMutation {
	return neuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (neuo *NilExampleUpdateOne) Select(field string, fields ...string) *NilExampleUpdateOne {
	neuo.fields = append([]string{field}, fields...)
	return neuo
}

// Save executes the query and returns the updated NilExample entity.
func (neuo *NilExampleUpdateOne) Save(ctx context.Context) (*NilExample, error) {
	var (
		err  error
		node *NilExample
	)
	if len(neuo.hooks) == 0 {
		node, err = neuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NilExampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			neuo.mutation = mutation
			node, err = neuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(neuo.hooks) - 1; i >= 0; i-- {
			if neuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = neuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, neuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (neuo *NilExampleUpdateOne) SaveX(ctx context.Context) *NilExample {
	node, err := neuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (neuo *NilExampleUpdateOne) Exec(ctx context.Context) error {
	_, err := neuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (neuo *NilExampleUpdateOne) ExecX(ctx context.Context) {
	if err := neuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (neuo *NilExampleUpdateOne) sqlSave(ctx context.Context) (_node *NilExample, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nilexample.Table,
			Columns: nilexample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nilexample.FieldID,
			},
		},
	}
	id, ok := neuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NilExample.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := neuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nilexample.FieldID)
		for _, f := range fields {
			if !nilexample.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nilexample.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := neuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := neuo.mutation.StrNil(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nilexample.FieldStrNil,
		})
	}
	if neuo.mutation.StrNilCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: nilexample.FieldStrNil,
		})
	}
	if value, ok := neuo.mutation.TimeNil(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: nilexample.FieldTimeNil,
		})
	}
	if neuo.mutation.TimeNilCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: nilexample.FieldTimeNil,
		})
	}
	_node = &NilExample{config: neuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, neuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nilexample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
