// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"server/infrastructure/ent/forumlike"
	"server/infrastructure/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ForumLikeDelete is the builder for deleting a ForumLike entity.
type ForumLikeDelete struct {
	config
	hooks    []Hook
	mutation *ForumLikeMutation
}

// Where appends a list predicates to the ForumLikeDelete builder.
func (fld *ForumLikeDelete) Where(ps ...predicate.ForumLike) *ForumLikeDelete {
	fld.mutation.Where(ps...)
	return fld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fld *ForumLikeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fld.sqlExec, fld.mutation, fld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fld *ForumLikeDelete) ExecX(ctx context.Context) int {
	n, err := fld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fld *ForumLikeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(forumlike.Table, sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt))
	if ps := fld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fld.mutation.done = true
	return affected, err
}

// ForumLikeDeleteOne is the builder for deleting a single ForumLike entity.
type ForumLikeDeleteOne struct {
	fld *ForumLikeDelete
}

// Where appends a list predicates to the ForumLikeDelete builder.
func (fldo *ForumLikeDeleteOne) Where(ps ...predicate.ForumLike) *ForumLikeDeleteOne {
	fldo.fld.mutation.Where(ps...)
	return fldo
}

// Exec executes the deletion query.
func (fldo *ForumLikeDeleteOne) Exec(ctx context.Context) error {
	n, err := fldo.fld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{forumlike.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fldo *ForumLikeDeleteOne) ExecX(ctx context.Context) {
	if err := fldo.Exec(ctx); err != nil {
		panic(err)
	}
}
