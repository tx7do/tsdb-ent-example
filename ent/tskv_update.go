// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tsdb-ent-example/ent/predicate"
	"tsdb-ent-example/ent/tskv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TsKvUpdate is the builder for updating TsKv entities.
type TsKvUpdate struct {
	config
	hooks    []Hook
	mutation *TsKvMutation
}

// Where appends a list predicates to the TsKvUpdate builder.
func (tku *TsKvUpdate) Where(ps ...predicate.TsKv) *TsKvUpdate {
	tku.mutation.Where(ps...)
	return tku
}

// SetEntityID sets the "entity_id" field.
func (tku *TsKvUpdate) SetEntityID(u uuid.UUID) *TsKvUpdate {
	tku.mutation.SetEntityID(u)
	return tku
}

// SetKey sets the "key" field.
func (tku *TsKvUpdate) SetKey(i int) *TsKvUpdate {
	tku.mutation.ResetKey()
	tku.mutation.SetKey(i)
	return tku
}

// AddKey adds i to the "key" field.
func (tku *TsKvUpdate) AddKey(i int) *TsKvUpdate {
	tku.mutation.AddKey(i)
	return tku
}

// SetBoolV sets the "bool_v" field.
func (tku *TsKvUpdate) SetBoolV(b bool) *TsKvUpdate {
	tku.mutation.SetBoolV(b)
	return tku
}

// SetNillableBoolV sets the "bool_v" field if the given value is not nil.
func (tku *TsKvUpdate) SetNillableBoolV(b *bool) *TsKvUpdate {
	if b != nil {
		tku.SetBoolV(*b)
	}
	return tku
}

// ClearBoolV clears the value of the "bool_v" field.
func (tku *TsKvUpdate) ClearBoolV() *TsKvUpdate {
	tku.mutation.ClearBoolV()
	return tku
}

// SetStrV sets the "str_v" field.
func (tku *TsKvUpdate) SetStrV(s string) *TsKvUpdate {
	tku.mutation.SetStrV(s)
	return tku
}

// SetNillableStrV sets the "str_v" field if the given value is not nil.
func (tku *TsKvUpdate) SetNillableStrV(s *string) *TsKvUpdate {
	if s != nil {
		tku.SetStrV(*s)
	}
	return tku
}

// ClearStrV clears the value of the "str_v" field.
func (tku *TsKvUpdate) ClearStrV() *TsKvUpdate {
	tku.mutation.ClearStrV()
	return tku
}

// SetLongV sets the "long_v" field.
func (tku *TsKvUpdate) SetLongV(i int64) *TsKvUpdate {
	tku.mutation.ResetLongV()
	tku.mutation.SetLongV(i)
	return tku
}

// SetNillableLongV sets the "long_v" field if the given value is not nil.
func (tku *TsKvUpdate) SetNillableLongV(i *int64) *TsKvUpdate {
	if i != nil {
		tku.SetLongV(*i)
	}
	return tku
}

// AddLongV adds i to the "long_v" field.
func (tku *TsKvUpdate) AddLongV(i int64) *TsKvUpdate {
	tku.mutation.AddLongV(i)
	return tku
}

// ClearLongV clears the value of the "long_v" field.
func (tku *TsKvUpdate) ClearLongV() *TsKvUpdate {
	tku.mutation.ClearLongV()
	return tku
}

// SetDblV sets the "dbl_v" field.
func (tku *TsKvUpdate) SetDblV(f float64) *TsKvUpdate {
	tku.mutation.ResetDblV()
	tku.mutation.SetDblV(f)
	return tku
}

// SetNillableDblV sets the "dbl_v" field if the given value is not nil.
func (tku *TsKvUpdate) SetNillableDblV(f *float64) *TsKvUpdate {
	if f != nil {
		tku.SetDblV(*f)
	}
	return tku
}

// AddDblV adds f to the "dbl_v" field.
func (tku *TsKvUpdate) AddDblV(f float64) *TsKvUpdate {
	tku.mutation.AddDblV(f)
	return tku
}

// ClearDblV clears the value of the "dbl_v" field.
func (tku *TsKvUpdate) ClearDblV() *TsKvUpdate {
	tku.mutation.ClearDblV()
	return tku
}

// SetJSONV sets the "json_v" field.
func (tku *TsKvUpdate) SetJSONV(s string) *TsKvUpdate {
	tku.mutation.SetJSONV(s)
	return tku
}

// SetNillableJSONV sets the "json_v" field if the given value is not nil.
func (tku *TsKvUpdate) SetNillableJSONV(s *string) *TsKvUpdate {
	if s != nil {
		tku.SetJSONV(*s)
	}
	return tku
}

// ClearJSONV clears the value of the "json_v" field.
func (tku *TsKvUpdate) ClearJSONV() *TsKvUpdate {
	tku.mutation.ClearJSONV()
	return tku
}

// SetTs sets the "ts" field.
func (tku *TsKvUpdate) SetTs(i int64) *TsKvUpdate {
	tku.mutation.ResetTs()
	tku.mutation.SetTs(i)
	return tku
}

// AddTs adds i to the "ts" field.
func (tku *TsKvUpdate) AddTs(i int64) *TsKvUpdate {
	tku.mutation.AddTs(i)
	return tku
}

// Mutation returns the TsKvMutation object of the builder.
func (tku *TsKvUpdate) Mutation() *TsKvMutation {
	return tku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tku *TsKvUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tku.hooks) == 0 {
		if err = tku.check(); err != nil {
			return 0, err
		}
		affected, err = tku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TsKvMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tku.check(); err != nil {
				return 0, err
			}
			tku.mutation = mutation
			affected, err = tku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tku.hooks) - 1; i >= 0; i-- {
			if tku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tku *TsKvUpdate) SaveX(ctx context.Context) int {
	affected, err := tku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tku *TsKvUpdate) Exec(ctx context.Context) error {
	_, err := tku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tku *TsKvUpdate) ExecX(ctx context.Context) {
	if err := tku.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tku *TsKvUpdate) check() error {
	if v, ok := tku.mutation.StrV(); ok {
		if err := tskv.StrVValidator(v); err != nil {
			return &ValidationError{Name: "str_v", err: fmt.Errorf(`ent: validator failed for field "TsKv.str_v": %w`, err)}
		}
	}
	return nil
}

func (tku *TsKvUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tskv.Table,
			Columns: tskv.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tskv.FieldID,
			},
		},
	}
	if ps := tku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tku.mutation.EntityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tskv.FieldEntityID,
		})
	}
	if value, ok := tku.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tskv.FieldKey,
		})
	}
	if value, ok := tku.mutation.AddedKey(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tskv.FieldKey,
		})
	}
	if value, ok := tku.mutation.BoolV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tskv.FieldBoolV,
		})
	}
	if tku.mutation.BoolVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: tskv.FieldBoolV,
		})
	}
	if value, ok := tku.mutation.StrV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tskv.FieldStrV,
		})
	}
	if tku.mutation.StrVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tskv.FieldStrV,
		})
	}
	if value, ok := tku.mutation.LongV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldLongV,
		})
	}
	if value, ok := tku.mutation.AddedLongV(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldLongV,
		})
	}
	if tku.mutation.LongVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tskv.FieldLongV,
		})
	}
	if value, ok := tku.mutation.DblV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tskv.FieldDblV,
		})
	}
	if value, ok := tku.mutation.AddedDblV(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tskv.FieldDblV,
		})
	}
	if tku.mutation.DblVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: tskv.FieldDblV,
		})
	}
	if value, ok := tku.mutation.JSONV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tskv.FieldJSONV,
		})
	}
	if tku.mutation.JSONVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tskv.FieldJSONV,
		})
	}
	if value, ok := tku.mutation.Ts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldTs,
		})
	}
	if value, ok := tku.mutation.AddedTs(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldTs,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tskv.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TsKvUpdateOne is the builder for updating a single TsKv entity.
type TsKvUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TsKvMutation
}

// SetEntityID sets the "entity_id" field.
func (tkuo *TsKvUpdateOne) SetEntityID(u uuid.UUID) *TsKvUpdateOne {
	tkuo.mutation.SetEntityID(u)
	return tkuo
}

// SetKey sets the "key" field.
func (tkuo *TsKvUpdateOne) SetKey(i int) *TsKvUpdateOne {
	tkuo.mutation.ResetKey()
	tkuo.mutation.SetKey(i)
	return tkuo
}

// AddKey adds i to the "key" field.
func (tkuo *TsKvUpdateOne) AddKey(i int) *TsKvUpdateOne {
	tkuo.mutation.AddKey(i)
	return tkuo
}

// SetBoolV sets the "bool_v" field.
func (tkuo *TsKvUpdateOne) SetBoolV(b bool) *TsKvUpdateOne {
	tkuo.mutation.SetBoolV(b)
	return tkuo
}

// SetNillableBoolV sets the "bool_v" field if the given value is not nil.
func (tkuo *TsKvUpdateOne) SetNillableBoolV(b *bool) *TsKvUpdateOne {
	if b != nil {
		tkuo.SetBoolV(*b)
	}
	return tkuo
}

// ClearBoolV clears the value of the "bool_v" field.
func (tkuo *TsKvUpdateOne) ClearBoolV() *TsKvUpdateOne {
	tkuo.mutation.ClearBoolV()
	return tkuo
}

// SetStrV sets the "str_v" field.
func (tkuo *TsKvUpdateOne) SetStrV(s string) *TsKvUpdateOne {
	tkuo.mutation.SetStrV(s)
	return tkuo
}

// SetNillableStrV sets the "str_v" field if the given value is not nil.
func (tkuo *TsKvUpdateOne) SetNillableStrV(s *string) *TsKvUpdateOne {
	if s != nil {
		tkuo.SetStrV(*s)
	}
	return tkuo
}

// ClearStrV clears the value of the "str_v" field.
func (tkuo *TsKvUpdateOne) ClearStrV() *TsKvUpdateOne {
	tkuo.mutation.ClearStrV()
	return tkuo
}

// SetLongV sets the "long_v" field.
func (tkuo *TsKvUpdateOne) SetLongV(i int64) *TsKvUpdateOne {
	tkuo.mutation.ResetLongV()
	tkuo.mutation.SetLongV(i)
	return tkuo
}

// SetNillableLongV sets the "long_v" field if the given value is not nil.
func (tkuo *TsKvUpdateOne) SetNillableLongV(i *int64) *TsKvUpdateOne {
	if i != nil {
		tkuo.SetLongV(*i)
	}
	return tkuo
}

// AddLongV adds i to the "long_v" field.
func (tkuo *TsKvUpdateOne) AddLongV(i int64) *TsKvUpdateOne {
	tkuo.mutation.AddLongV(i)
	return tkuo
}

// ClearLongV clears the value of the "long_v" field.
func (tkuo *TsKvUpdateOne) ClearLongV() *TsKvUpdateOne {
	tkuo.mutation.ClearLongV()
	return tkuo
}

// SetDblV sets the "dbl_v" field.
func (tkuo *TsKvUpdateOne) SetDblV(f float64) *TsKvUpdateOne {
	tkuo.mutation.ResetDblV()
	tkuo.mutation.SetDblV(f)
	return tkuo
}

// SetNillableDblV sets the "dbl_v" field if the given value is not nil.
func (tkuo *TsKvUpdateOne) SetNillableDblV(f *float64) *TsKvUpdateOne {
	if f != nil {
		tkuo.SetDblV(*f)
	}
	return tkuo
}

// AddDblV adds f to the "dbl_v" field.
func (tkuo *TsKvUpdateOne) AddDblV(f float64) *TsKvUpdateOne {
	tkuo.mutation.AddDblV(f)
	return tkuo
}

// ClearDblV clears the value of the "dbl_v" field.
func (tkuo *TsKvUpdateOne) ClearDblV() *TsKvUpdateOne {
	tkuo.mutation.ClearDblV()
	return tkuo
}

// SetJSONV sets the "json_v" field.
func (tkuo *TsKvUpdateOne) SetJSONV(s string) *TsKvUpdateOne {
	tkuo.mutation.SetJSONV(s)
	return tkuo
}

// SetNillableJSONV sets the "json_v" field if the given value is not nil.
func (tkuo *TsKvUpdateOne) SetNillableJSONV(s *string) *TsKvUpdateOne {
	if s != nil {
		tkuo.SetJSONV(*s)
	}
	return tkuo
}

// ClearJSONV clears the value of the "json_v" field.
func (tkuo *TsKvUpdateOne) ClearJSONV() *TsKvUpdateOne {
	tkuo.mutation.ClearJSONV()
	return tkuo
}

// SetTs sets the "ts" field.
func (tkuo *TsKvUpdateOne) SetTs(i int64) *TsKvUpdateOne {
	tkuo.mutation.ResetTs()
	tkuo.mutation.SetTs(i)
	return tkuo
}

// AddTs adds i to the "ts" field.
func (tkuo *TsKvUpdateOne) AddTs(i int64) *TsKvUpdateOne {
	tkuo.mutation.AddTs(i)
	return tkuo
}

// Mutation returns the TsKvMutation object of the builder.
func (tkuo *TsKvUpdateOne) Mutation() *TsKvMutation {
	return tkuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tkuo *TsKvUpdateOne) Select(field string, fields ...string) *TsKvUpdateOne {
	tkuo.fields = append([]string{field}, fields...)
	return tkuo
}

// Save executes the query and returns the updated TsKv entity.
func (tkuo *TsKvUpdateOne) Save(ctx context.Context) (*TsKv, error) {
	var (
		err  error
		node *TsKv
	)
	if len(tkuo.hooks) == 0 {
		if err = tkuo.check(); err != nil {
			return nil, err
		}
		node, err = tkuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TsKvMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tkuo.check(); err != nil {
				return nil, err
			}
			tkuo.mutation = mutation
			node, err = tkuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tkuo.hooks) - 1; i >= 0; i-- {
			if tkuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tkuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tkuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tkuo *TsKvUpdateOne) SaveX(ctx context.Context) *TsKv {
	node, err := tkuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tkuo *TsKvUpdateOne) Exec(ctx context.Context) error {
	_, err := tkuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tkuo *TsKvUpdateOne) ExecX(ctx context.Context) {
	if err := tkuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tkuo *TsKvUpdateOne) check() error {
	if v, ok := tkuo.mutation.StrV(); ok {
		if err := tskv.StrVValidator(v); err != nil {
			return &ValidationError{Name: "str_v", err: fmt.Errorf(`ent: validator failed for field "TsKv.str_v": %w`, err)}
		}
	}
	return nil
}

func (tkuo *TsKvUpdateOne) sqlSave(ctx context.Context) (_node *TsKv, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tskv.Table,
			Columns: tskv.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tskv.FieldID,
			},
		},
	}
	id, ok := tkuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TsKv.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tkuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tskv.FieldID)
		for _, f := range fields {
			if !tskv.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tskv.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tkuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tkuo.mutation.EntityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tskv.FieldEntityID,
		})
	}
	if value, ok := tkuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tskv.FieldKey,
		})
	}
	if value, ok := tkuo.mutation.AddedKey(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tskv.FieldKey,
		})
	}
	if value, ok := tkuo.mutation.BoolV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tskv.FieldBoolV,
		})
	}
	if tkuo.mutation.BoolVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: tskv.FieldBoolV,
		})
	}
	if value, ok := tkuo.mutation.StrV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tskv.FieldStrV,
		})
	}
	if tkuo.mutation.StrVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tskv.FieldStrV,
		})
	}
	if value, ok := tkuo.mutation.LongV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldLongV,
		})
	}
	if value, ok := tkuo.mutation.AddedLongV(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldLongV,
		})
	}
	if tkuo.mutation.LongVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tskv.FieldLongV,
		})
	}
	if value, ok := tkuo.mutation.DblV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tskv.FieldDblV,
		})
	}
	if value, ok := tkuo.mutation.AddedDblV(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tskv.FieldDblV,
		})
	}
	if tkuo.mutation.DblVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: tskv.FieldDblV,
		})
	}
	if value, ok := tkuo.mutation.JSONV(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tskv.FieldJSONV,
		})
	}
	if tkuo.mutation.JSONVCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tskv.FieldJSONV,
		})
	}
	if value, ok := tkuo.mutation.Ts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldTs,
		})
	}
	if value, ok := tkuo.mutation.AddedTs(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tskv.FieldTs,
		})
	}
	_node = &TsKv{config: tkuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tkuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tskv.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
