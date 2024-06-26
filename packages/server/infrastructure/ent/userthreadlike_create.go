// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userthreadlike"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserThreadLikeCreate is the builder for creating a UserThreadLike entity.
type UserThreadLikeCreate struct {
	config
	mutation *UserThreadLikeMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (utlc *UserThreadLikeCreate) SetUserId(i int) *UserThreadLikeCreate {
	utlc.mutation.SetUserId(i)
	return utlc
}

// SetThreadId sets the "threadId" field.
func (utlc *UserThreadLikeCreate) SetThreadId(i int) *UserThreadLikeCreate {
	utlc.mutation.SetThreadId(i)
	return utlc
}

// SetLikedAt sets the "likedAt" field.
func (utlc *UserThreadLikeCreate) SetLikedAt(t time.Time) *UserThreadLikeCreate {
	utlc.mutation.SetLikedAt(t)
	return utlc
}

// SetNillableLikedAt sets the "likedAt" field if the given value is not nil.
func (utlc *UserThreadLikeCreate) SetNillableLikedAt(t *time.Time) *UserThreadLikeCreate {
	if t != nil {
		utlc.SetLikedAt(*t)
	}
	return utlc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (utlc *UserThreadLikeCreate) SetUserID(id int) *UserThreadLikeCreate {
	utlc.mutation.SetUserID(id)
	return utlc
}

// SetUser sets the "user" edge to the User entity.
func (utlc *UserThreadLikeCreate) SetUser(u *User) *UserThreadLikeCreate {
	return utlc.SetUserID(u.ID)
}

// SetThreadID sets the "thread" edge to the Thread entity by ID.
func (utlc *UserThreadLikeCreate) SetThreadID(id int) *UserThreadLikeCreate {
	utlc.mutation.SetThreadID(id)
	return utlc
}

// SetThread sets the "thread" edge to the Thread entity.
func (utlc *UserThreadLikeCreate) SetThread(t *Thread) *UserThreadLikeCreate {
	return utlc.SetThreadID(t.ID)
}

// Mutation returns the UserThreadLikeMutation object of the builder.
func (utlc *UserThreadLikeCreate) Mutation() *UserThreadLikeMutation {
	return utlc.mutation
}

// Save creates the UserThreadLike in the database.
func (utlc *UserThreadLikeCreate) Save(ctx context.Context) (*UserThreadLike, error) {
	utlc.defaults()
	return withHooks(ctx, utlc.sqlSave, utlc.mutation, utlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (utlc *UserThreadLikeCreate) SaveX(ctx context.Context) *UserThreadLike {
	v, err := utlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utlc *UserThreadLikeCreate) Exec(ctx context.Context) error {
	_, err := utlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utlc *UserThreadLikeCreate) ExecX(ctx context.Context) {
	if err := utlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (utlc *UserThreadLikeCreate) defaults() {
	if _, ok := utlc.mutation.LikedAt(); !ok {
		v := userthreadlike.DefaultLikedAt()
		utlc.mutation.SetLikedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utlc *UserThreadLikeCreate) check() error {
	if _, ok := utlc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "UserThreadLike.userId"`)}
	}
	if _, ok := utlc.mutation.ThreadId(); !ok {
		return &ValidationError{Name: "threadId", err: errors.New(`ent: missing required field "UserThreadLike.threadId"`)}
	}
	if _, ok := utlc.mutation.LikedAt(); !ok {
		return &ValidationError{Name: "likedAt", err: errors.New(`ent: missing required field "UserThreadLike.likedAt"`)}
	}
	if _, ok := utlc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserThreadLike.user"`)}
	}
	if _, ok := utlc.mutation.ThreadID(); !ok {
		return &ValidationError{Name: "thread", err: errors.New(`ent: missing required edge "UserThreadLike.thread"`)}
	}
	return nil
}

func (utlc *UserThreadLikeCreate) sqlSave(ctx context.Context) (*UserThreadLike, error) {
	if err := utlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := utlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, utlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (utlc *UserThreadLikeCreate) createSpec() (*UserThreadLike, *sqlgraph.CreateSpec) {
	var (
		_node = &UserThreadLike{config: utlc.config}
		_spec = sqlgraph.NewCreateSpec(userthreadlike.Table, nil)
	)
	if value, ok := utlc.mutation.LikedAt(); ok {
		_spec.SetField(userthreadlike.FieldLikedAt, field.TypeTime, value)
		_node.LikedAt = value
	}
	if nodes := utlc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userthreadlike.UserTable,
			Columns: []string{userthreadlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := utlc.mutation.ThreadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userthreadlike.ThreadTable,
			Columns: []string{userthreadlike.ThreadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ThreadId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserThreadLikeCreateBulk is the builder for creating many UserThreadLike entities in bulk.
type UserThreadLikeCreateBulk struct {
	config
	err      error
	builders []*UserThreadLikeCreate
}

// Save creates the UserThreadLike entities in the database.
func (utlcb *UserThreadLikeCreateBulk) Save(ctx context.Context) ([]*UserThreadLike, error) {
	if utlcb.err != nil {
		return nil, utlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(utlcb.builders))
	nodes := make([]*UserThreadLike, len(utlcb.builders))
	mutators := make([]Mutator, len(utlcb.builders))
	for i := range utlcb.builders {
		func(i int, root context.Context) {
			builder := utlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserThreadLikeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, utlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, utlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, utlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (utlcb *UserThreadLikeCreateBulk) SaveX(ctx context.Context) []*UserThreadLike {
	v, err := utlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utlcb *UserThreadLikeCreateBulk) Exec(ctx context.Context) error {
	_, err := utlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utlcb *UserThreadLikeCreateBulk) ExecX(ctx context.Context) {
	if err := utlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
