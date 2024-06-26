// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThreadCreate is the builder for creating a Thread entity.
type ThreadCreate struct {
	config
	mutation *ThreadMutation
	hooks    []Hook
}

// SetBoardId sets the "boardId" field.
func (tc *ThreadCreate) SetBoardId(i int) *ThreadCreate {
	tc.mutation.SetBoardId(i)
	return tc
}

// SetUserId sets the "userId" field.
func (tc *ThreadCreate) SetUserId(i int) *ThreadCreate {
	tc.mutation.SetUserId(i)
	return tc
}

// SetTitle sets the "title" field.
func (tc *ThreadCreate) SetTitle(s string) *ThreadCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *ThreadCreate) SetDescription(s string) *ThreadCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableDescription(s *string) *ThreadCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetThumbnailUrl sets the "thumbnailUrl" field.
func (tc *ThreadCreate) SetThumbnailUrl(s string) *ThreadCreate {
	tc.mutation.SetThumbnailUrl(s)
	return tc
}

// SetNillableThumbnailUrl sets the "thumbnailUrl" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableThumbnailUrl(s *string) *ThreadCreate {
	if s != nil {
		tc.SetThumbnailUrl(*s)
	}
	return tc
}

// SetIPAddress sets the "ip_address" field.
func (tc *ThreadCreate) SetIPAddress(s string) *ThreadCreate {
	tc.mutation.SetIPAddress(s)
	return tc
}

// SetStatus sets the "status" field.
func (tc *ThreadCreate) SetStatus(i int) *ThreadCreate {
	tc.mutation.SetStatus(i)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableStatus(i *int) *ThreadCreate {
	if i != nil {
		tc.SetStatus(*i)
	}
	return tc
}

// SetCreatedAt sets the "createdAt" field.
func (tc *ThreadCreate) SetCreatedAt(t time.Time) *ThreadCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableCreatedAt(t *time.Time) *ThreadCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updatedAt" field.
func (tc *ThreadCreate) SetUpdatedAt(t time.Time) *ThreadCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableUpdatedAt(t *time.Time) *ThreadCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *ThreadCreate) SetID(i int) *ThreadCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetBoardID sets the "board" edge to the Board entity by ID.
func (tc *ThreadCreate) SetBoardID(id int) *ThreadCreate {
	tc.mutation.SetBoardID(id)
	return tc
}

// SetBoard sets the "board" edge to the Board entity.
func (tc *ThreadCreate) SetBoard(b *Board) *ThreadCreate {
	return tc.SetBoardID(b.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (tc *ThreadCreate) SetOwnerID(id int) *ThreadCreate {
	tc.mutation.SetOwnerID(id)
	return tc
}

// SetOwner sets the "owner" edge to the User entity.
func (tc *ThreadCreate) SetOwner(u *User) *ThreadCreate {
	return tc.SetOwnerID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the ThreadComment entity by IDs.
func (tc *ThreadCreate) AddCommentIDs(ids ...int) *ThreadCreate {
	tc.mutation.AddCommentIDs(ids...)
	return tc
}

// AddComments adds the "comments" edges to the ThreadComment entity.
func (tc *ThreadCreate) AddComments(t ...*ThreadComment) *ThreadCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tc *ThreadCreate) AddTagIDs(ids ...int) *ThreadCreate {
	tc.mutation.AddTagIDs(ids...)
	return tc
}

// AddTags adds the "tags" edges to the Tag entity.
func (tc *ThreadCreate) AddTags(t ...*Tag) *ThreadCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTagIDs(ids...)
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tc *ThreadCreate) AddLikedUserIDs(ids ...int) *ThreadCreate {
	tc.mutation.AddLikedUserIDs(ids...)
	return tc
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tc *ThreadCreate) AddLikedUsers(u ...*User) *ThreadCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddLikedUserIDs(ids...)
}

// AddSubscribedUserIDs adds the "subscribed_users" edge to the User entity by IDs.
func (tc *ThreadCreate) AddSubscribedUserIDs(ids ...int) *ThreadCreate {
	tc.mutation.AddSubscribedUserIDs(ids...)
	return tc
}

// AddSubscribedUsers adds the "subscribed_users" edges to the User entity.
func (tc *ThreadCreate) AddSubscribedUsers(u ...*User) *ThreadCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddSubscribedUserIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tc *ThreadCreate) Mutation() *ThreadMutation {
	return tc.mutation
}

// Save creates the Thread in the database.
func (tc *ThreadCreate) Save(ctx context.Context) (*Thread, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ThreadCreate) SaveX(ctx context.Context) *Thread {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ThreadCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ThreadCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *ThreadCreate) defaults() {
	if _, ok := tc.mutation.Status(); !ok {
		v := thread.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := thread.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := thread.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ThreadCreate) check() error {
	if _, ok := tc.mutation.BoardId(); !ok {
		return &ValidationError{Name: "boardId", err: errors.New(`ent: missing required field "Thread.boardId"`)}
	}
	if _, ok := tc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "Thread.userId"`)}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Thread.title"`)}
	}
	if v, ok := tc.mutation.Title(); ok {
		if err := thread.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Thread.title": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Description(); ok {
		if err := thread.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Thread.description": %w`, err)}
		}
	}
	if _, ok := tc.mutation.IPAddress(); !ok {
		return &ValidationError{Name: "ip_address", err: errors.New(`ent: missing required field "Thread.ip_address"`)}
	}
	if v, ok := tc.mutation.IPAddress(); ok {
		if err := thread.IPAddressValidator(v); err != nil {
			return &ValidationError{Name: "ip_address", err: fmt.Errorf(`ent: validator failed for field "Thread.ip_address": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Thread.status"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Thread.createdAt"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Thread.updatedAt"`)}
	}
	if _, ok := tc.mutation.BoardID(); !ok {
		return &ValidationError{Name: "board", err: errors.New(`ent: missing required edge "Thread.board"`)}
	}
	if _, ok := tc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Thread.owner"`)}
	}
	return nil
}

func (tc *ThreadCreate) sqlSave(ctx context.Context) (*Thread, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ThreadCreate) createSpec() (*Thread, *sqlgraph.CreateSpec) {
	var (
		_node = &Thread{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(thread.Table, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(thread.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(thread.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.ThumbnailUrl(); ok {
		_spec.SetField(thread.FieldThumbnailUrl, field.TypeString, value)
		_node.ThumbnailUrl = value
	}
	if value, ok := tc.mutation.IPAddress(); ok {
		_spec.SetField(thread.FieldIPAddress, field.TypeString, value)
		_node.IPAddress = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(thread.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(thread.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(thread.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.BoardTable,
			Columns: []string{thread.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BoardId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.OwnerTable,
			Columns: []string{thread.OwnerColumn},
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
	if nodes := tc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.CommentsTable,
			Columns: []string{thread.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thread.TagsTable,
			Columns: thread.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   thread.LikedUsersTable,
			Columns: thread.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserThreadLikeCreate{config: tc.config, mutation: newUserThreadLikeMutation(tc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SubscribedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   thread.SubscribedUsersTable,
			Columns: thread.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserThreadSubscriptionCreate{config: tc.config, mutation: newUserThreadSubscriptionMutation(tc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThreadCreateBulk is the builder for creating many Thread entities in bulk.
type ThreadCreateBulk struct {
	config
	err      error
	builders []*ThreadCreate
}

// Save creates the Thread entities in the database.
func (tcb *ThreadCreateBulk) Save(ctx context.Context) ([]*Thread, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Thread, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ThreadCreateBulk) SaveX(ctx context.Context) []*Thread {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ThreadCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ThreadCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
