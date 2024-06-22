// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BoardCreate is the builder for creating a Board entity.
type BoardCreate struct {
	config
	mutation *BoardMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (bc *BoardCreate) SetUserId(i int) *BoardCreate {
	bc.mutation.SetUserId(i)
	return bc
}

// SetTitle sets the "title" field.
func (bc *BoardCreate) SetTitle(s string) *BoardCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BoardCreate) SetDescription(s string) *BoardCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bc *BoardCreate) SetNillableDescription(s *string) *BoardCreate {
	if s != nil {
		bc.SetDescription(*s)
	}
	return bc
}

// SetThumbnailUrl sets the "thumbnailUrl" field.
func (bc *BoardCreate) SetThumbnailUrl(s string) *BoardCreate {
	bc.mutation.SetThumbnailUrl(s)
	return bc
}

// SetNillableThumbnailUrl sets the "thumbnailUrl" field if the given value is not nil.
func (bc *BoardCreate) SetNillableThumbnailUrl(s *string) *BoardCreate {
	if s != nil {
		bc.SetThumbnailUrl(*s)
	}
	return bc
}

// SetOrder sets the "order" field.
func (bc *BoardCreate) SetOrder(i int) *BoardCreate {
	bc.mutation.SetOrder(i)
	return bc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (bc *BoardCreate) SetNillableOrder(i *int) *BoardCreate {
	if i != nil {
		bc.SetOrder(*i)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BoardCreate) SetStatus(b board.Status) *BoardCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bc *BoardCreate) SetNillableStatus(b *board.Status) *BoardCreate {
	if b != nil {
		bc.SetStatus(*b)
	}
	return bc
}

// SetCreatedAt sets the "createdAt" field.
func (bc *BoardCreate) SetCreatedAt(t time.Time) *BoardCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bc *BoardCreate) SetNillableCreatedAt(t *time.Time) *BoardCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updatedAt" field.
func (bc *BoardCreate) SetUpdatedAt(t time.Time) *BoardCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bc *BoardCreate) SetNillableUpdatedAt(t *time.Time) *BoardCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BoardCreate) SetID(i int) *BoardCreate {
	bc.mutation.SetID(i)
	return bc
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (bc *BoardCreate) AddLikedUserIDs(ids ...int) *BoardCreate {
	bc.mutation.AddLikedUserIDs(ids...)
	return bc
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (bc *BoardCreate) AddLikedUsers(u ...*User) *BoardCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddLikedUserIDs(ids...)
}

// AddSubscribedUserIDs adds the "subscribed_users" edge to the User entity by IDs.
func (bc *BoardCreate) AddSubscribedUserIDs(ids ...int) *BoardCreate {
	bc.mutation.AddSubscribedUserIDs(ids...)
	return bc
}

// AddSubscribedUsers adds the "subscribed_users" edges to the User entity.
func (bc *BoardCreate) AddSubscribedUsers(u ...*User) *BoardCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddSubscribedUserIDs(ids...)
}

// AddThreadIDs adds the "threads" edge to the Thread entity by IDs.
func (bc *BoardCreate) AddThreadIDs(ids ...int) *BoardCreate {
	bc.mutation.AddThreadIDs(ids...)
	return bc
}

// AddThreads adds the "threads" edges to the Thread entity.
func (bc *BoardCreate) AddThreads(t ...*Thread) *BoardCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddThreadIDs(ids...)
}

// Mutation returns the BoardMutation object of the builder.
func (bc *BoardCreate) Mutation() *BoardMutation {
	return bc.mutation
}

// Save creates the Board in the database.
func (bc *BoardCreate) Save(ctx context.Context) (*Board, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BoardCreate) SaveX(ctx context.Context) *Board {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BoardCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BoardCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BoardCreate) defaults() {
	if _, ok := bc.mutation.Order(); !ok {
		v := board.DefaultOrder
		bc.mutation.SetOrder(v)
	}
	if _, ok := bc.mutation.Status(); !ok {
		v := board.DefaultStatus
		bc.mutation.SetStatus(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := board.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := board.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BoardCreate) check() error {
	if _, ok := bc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "Board.userId"`)}
	}
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Board.title"`)}
	}
	if v, ok := bc.mutation.Title(); ok {
		if err := board.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Board.title": %w`, err)}
		}
	}
	if v, ok := bc.mutation.Description(); ok {
		if err := board.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Board.description": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Board.order"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Board.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := board.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Board.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Board.createdAt"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Board.updatedAt"`)}
	}
	return nil
}

func (bc *BoardCreate) sqlSave(ctx context.Context) (*Board, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BoardCreate) createSpec() (*Board, *sqlgraph.CreateSpec) {
	var (
		_node = &Board{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(board.Table, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.UserId(); ok {
		_spec.SetField(board.FieldUserId, field.TypeInt, value)
		_node.UserId = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.SetField(board.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := bc.mutation.ThumbnailUrl(); ok {
		_spec.SetField(board.FieldThumbnailUrl, field.TypeString, value)
		_node.ThumbnailUrl = value
	}
	if value, ok := bc.mutation.Order(); ok {
		_spec.SetField(board.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(board.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardSubscriptionCreate{config: bc.config, mutation: newUserBoardSubscriptionMutation(bc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.SubscribedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardLikeCreate{config: bc.config, mutation: newUserBoardLikeMutation(bc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BoardCreateBulk is the builder for creating many Board entities in bulk.
type BoardCreateBulk struct {
	config
	err      error
	builders []*BoardCreate
}

// Save creates the Board entities in the database.
func (bcb *BoardCreateBulk) Save(ctx context.Context) ([]*Board, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Board, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BoardMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BoardCreateBulk) SaveX(ctx context.Context) []*Board {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BoardCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BoardCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}