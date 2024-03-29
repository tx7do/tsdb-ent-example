// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tsdb-ent-example/ent/predicate"
	"tsdb-ent-example/ent/tskvlatest"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TsKvLatestDelete is the builder for deleting a TsKvLatest entity.
type TsKvLatestDelete struct {
	config
	hooks    []Hook
	mutation *TsKvLatestMutation
}

// Where appends a list predicates to the TsKvLatestDelete builder.
func (tkld *TsKvLatestDelete) Where(ps ...predicate.TsKvLatest) *TsKvLatestDelete {
	tkld.mutation.Where(ps...)
	return tkld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tkld *TsKvLatestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tkld.hooks) == 0 {
		affected, err = tkld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TsKvLatestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tkld.mutation = mutation
			affected, err = tkld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tkld.hooks) - 1; i >= 0; i-- {
			if tkld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tkld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tkld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tkld *TsKvLatestDelete) ExecX(ctx context.Context) int {
	n, err := tkld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tkld *TsKvLatestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tskvlatest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tskvlatest.FieldID,
			},
		},
	}
	if ps := tkld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tkld.driver, _spec)
}

// TsKvLatestDeleteOne is the builder for deleting a single TsKvLatest entity.
type TsKvLatestDeleteOne struct {
	tkld *TsKvLatestDelete
}

// Exec executes the deletion query.
func (tkldo *TsKvLatestDeleteOne) Exec(ctx context.Context) error {
	n, err := tkldo.tkld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tskvlatest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tkldo *TsKvLatestDeleteOne) ExecX(ctx context.Context) {
	tkldo.tkld.ExecX(ctx)
}
