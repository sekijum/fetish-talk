// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/threadtag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThreadTagDelete is the builder for deleting a ThreadTag entity.
type ThreadTagDelete struct {
	config
	hooks    []Hook
	mutation *ThreadTagMutation
}

// Where appends a list predicates to the ThreadTagDelete builder.
func (ttd *ThreadTagDelete) Where(ps ...predicate.ThreadTag) *ThreadTagDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *ThreadTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ttd.sqlExec, ttd.mutation, ttd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *ThreadTagDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *ThreadTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(threadtag.Table, sqlgraph.NewFieldSpec(threadtag.FieldID, field.TypeInt))
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ttd.mutation.done = true
	return affected, err
}

// ThreadTagDeleteOne is the builder for deleting a single ThreadTag entity.
type ThreadTagDeleteOne struct {
	ttd *ThreadTagDelete
}

// Where appends a list predicates to the ThreadTagDelete builder.
func (ttdo *ThreadTagDeleteOne) Where(ps ...predicate.ThreadTag) *ThreadTagDeleteOne {
	ttdo.ttd.mutation.Where(ps...)
	return ttdo
}

// Exec executes the deletion query.
func (ttdo *ThreadTagDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{threadtag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *ThreadTagDeleteOne) ExecX(ctx context.Context) {
	if err := ttdo.Exec(ctx); err != nil {
		panic(err)
	}
}