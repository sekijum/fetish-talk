// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userboardlike"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBoardLikeUpdate is the builder for updating UserBoardLike entities.
type UserBoardLikeUpdate struct {
	config
	hooks    []Hook
	mutation *UserBoardLikeMutation
}

// Where appends a list predicates to the UserBoardLikeUpdate builder.
func (ublu *UserBoardLikeUpdate) Where(ps ...predicate.UserBoardLike) *UserBoardLikeUpdate {
	ublu.mutation.Where(ps...)
	return ublu
}

// SetUserId sets the "userId" field.
func (ublu *UserBoardLikeUpdate) SetUserId(i int) *UserBoardLikeUpdate {
	ublu.mutation.SetUserId(i)
	return ublu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ublu *UserBoardLikeUpdate) SetNillableUserId(i *int) *UserBoardLikeUpdate {
	if i != nil {
		ublu.SetUserId(*i)
	}
	return ublu
}

// SetBoardId sets the "boardId" field.
func (ublu *UserBoardLikeUpdate) SetBoardId(i int) *UserBoardLikeUpdate {
	ublu.mutation.SetBoardId(i)
	return ublu
}

// SetNillableBoardId sets the "boardId" field if the given value is not nil.
func (ublu *UserBoardLikeUpdate) SetNillableBoardId(i *int) *UserBoardLikeUpdate {
	if i != nil {
		ublu.SetBoardId(*i)
	}
	return ublu
}

// SetLikedAt sets the "likedAt" field.
func (ublu *UserBoardLikeUpdate) SetLikedAt(t time.Time) *UserBoardLikeUpdate {
	ublu.mutation.SetLikedAt(t)
	return ublu
}

// SetNillableLikedAt sets the "likedAt" field if the given value is not nil.
func (ublu *UserBoardLikeUpdate) SetNillableLikedAt(t *time.Time) *UserBoardLikeUpdate {
	if t != nil {
		ublu.SetLikedAt(*t)
	}
	return ublu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ublu *UserBoardLikeUpdate) SetUserID(id int) *UserBoardLikeUpdate {
	ublu.mutation.SetUserID(id)
	return ublu
}

// SetUser sets the "user" edge to the User entity.
func (ublu *UserBoardLikeUpdate) SetUser(u *User) *UserBoardLikeUpdate {
	return ublu.SetUserID(u.ID)
}

// SetBoardID sets the "board" edge to the Board entity by ID.
func (ublu *UserBoardLikeUpdate) SetBoardID(id int) *UserBoardLikeUpdate {
	ublu.mutation.SetBoardID(id)
	return ublu
}

// SetBoard sets the "board" edge to the Board entity.
func (ublu *UserBoardLikeUpdate) SetBoard(b *Board) *UserBoardLikeUpdate {
	return ublu.SetBoardID(b.ID)
}

// Mutation returns the UserBoardLikeMutation object of the builder.
func (ublu *UserBoardLikeUpdate) Mutation() *UserBoardLikeMutation {
	return ublu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ublu *UserBoardLikeUpdate) ClearUser() *UserBoardLikeUpdate {
	ublu.mutation.ClearUser()
	return ublu
}

// ClearBoard clears the "board" edge to the Board entity.
func (ublu *UserBoardLikeUpdate) ClearBoard() *UserBoardLikeUpdate {
	ublu.mutation.ClearBoard()
	return ublu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ublu *UserBoardLikeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ublu.sqlSave, ublu.mutation, ublu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ublu *UserBoardLikeUpdate) SaveX(ctx context.Context) int {
	affected, err := ublu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ublu *UserBoardLikeUpdate) Exec(ctx context.Context) error {
	_, err := ublu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ublu *UserBoardLikeUpdate) ExecX(ctx context.Context) {
	if err := ublu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ublu *UserBoardLikeUpdate) check() error {
	if _, ok := ublu.mutation.UserID(); ublu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardLike.user"`)
	}
	if _, ok := ublu.mutation.BoardID(); ublu.mutation.BoardCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardLike.board"`)
	}
	return nil
}

func (ublu *UserBoardLikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ublu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userboardlike.Table, userboardlike.Columns, sqlgraph.NewFieldSpec(userboardlike.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(userboardlike.FieldBoardId, field.TypeInt))
	if ps := ublu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ublu.mutation.LikedAt(); ok {
		_spec.SetField(userboardlike.FieldLikedAt, field.TypeTime, value)
	}
	if ublu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.UserTable,
			Columns: []string{userboardlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ublu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.UserTable,
			Columns: []string{userboardlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ublu.mutation.BoardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.BoardTable,
			Columns: []string{userboardlike.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ublu.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.BoardTable,
			Columns: []string{userboardlike.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ublu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userboardlike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ublu.mutation.done = true
	return n, nil
}

// UserBoardLikeUpdateOne is the builder for updating a single UserBoardLike entity.
type UserBoardLikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserBoardLikeMutation
}

// SetUserId sets the "userId" field.
func (ubluo *UserBoardLikeUpdateOne) SetUserId(i int) *UserBoardLikeUpdateOne {
	ubluo.mutation.SetUserId(i)
	return ubluo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ubluo *UserBoardLikeUpdateOne) SetNillableUserId(i *int) *UserBoardLikeUpdateOne {
	if i != nil {
		ubluo.SetUserId(*i)
	}
	return ubluo
}

// SetBoardId sets the "boardId" field.
func (ubluo *UserBoardLikeUpdateOne) SetBoardId(i int) *UserBoardLikeUpdateOne {
	ubluo.mutation.SetBoardId(i)
	return ubluo
}

// SetNillableBoardId sets the "boardId" field if the given value is not nil.
func (ubluo *UserBoardLikeUpdateOne) SetNillableBoardId(i *int) *UserBoardLikeUpdateOne {
	if i != nil {
		ubluo.SetBoardId(*i)
	}
	return ubluo
}

// SetLikedAt sets the "likedAt" field.
func (ubluo *UserBoardLikeUpdateOne) SetLikedAt(t time.Time) *UserBoardLikeUpdateOne {
	ubluo.mutation.SetLikedAt(t)
	return ubluo
}

// SetNillableLikedAt sets the "likedAt" field if the given value is not nil.
func (ubluo *UserBoardLikeUpdateOne) SetNillableLikedAt(t *time.Time) *UserBoardLikeUpdateOne {
	if t != nil {
		ubluo.SetLikedAt(*t)
	}
	return ubluo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ubluo *UserBoardLikeUpdateOne) SetUserID(id int) *UserBoardLikeUpdateOne {
	ubluo.mutation.SetUserID(id)
	return ubluo
}

// SetUser sets the "user" edge to the User entity.
func (ubluo *UserBoardLikeUpdateOne) SetUser(u *User) *UserBoardLikeUpdateOne {
	return ubluo.SetUserID(u.ID)
}

// SetBoardID sets the "board" edge to the Board entity by ID.
func (ubluo *UserBoardLikeUpdateOne) SetBoardID(id int) *UserBoardLikeUpdateOne {
	ubluo.mutation.SetBoardID(id)
	return ubluo
}

// SetBoard sets the "board" edge to the Board entity.
func (ubluo *UserBoardLikeUpdateOne) SetBoard(b *Board) *UserBoardLikeUpdateOne {
	return ubluo.SetBoardID(b.ID)
}

// Mutation returns the UserBoardLikeMutation object of the builder.
func (ubluo *UserBoardLikeUpdateOne) Mutation() *UserBoardLikeMutation {
	return ubluo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ubluo *UserBoardLikeUpdateOne) ClearUser() *UserBoardLikeUpdateOne {
	ubluo.mutation.ClearUser()
	return ubluo
}

// ClearBoard clears the "board" edge to the Board entity.
func (ubluo *UserBoardLikeUpdateOne) ClearBoard() *UserBoardLikeUpdateOne {
	ubluo.mutation.ClearBoard()
	return ubluo
}

// Where appends a list predicates to the UserBoardLikeUpdate builder.
func (ubluo *UserBoardLikeUpdateOne) Where(ps ...predicate.UserBoardLike) *UserBoardLikeUpdateOne {
	ubluo.mutation.Where(ps...)
	return ubluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ubluo *UserBoardLikeUpdateOne) Select(field string, fields ...string) *UserBoardLikeUpdateOne {
	ubluo.fields = append([]string{field}, fields...)
	return ubluo
}

// Save executes the query and returns the updated UserBoardLike entity.
func (ubluo *UserBoardLikeUpdateOne) Save(ctx context.Context) (*UserBoardLike, error) {
	return withHooks(ctx, ubluo.sqlSave, ubluo.mutation, ubluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ubluo *UserBoardLikeUpdateOne) SaveX(ctx context.Context) *UserBoardLike {
	node, err := ubluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ubluo *UserBoardLikeUpdateOne) Exec(ctx context.Context) error {
	_, err := ubluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubluo *UserBoardLikeUpdateOne) ExecX(ctx context.Context) {
	if err := ubluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubluo *UserBoardLikeUpdateOne) check() error {
	if _, ok := ubluo.mutation.UserID(); ubluo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardLike.user"`)
	}
	if _, ok := ubluo.mutation.BoardID(); ubluo.mutation.BoardCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardLike.board"`)
	}
	return nil
}

func (ubluo *UserBoardLikeUpdateOne) sqlSave(ctx context.Context) (_node *UserBoardLike, err error) {
	if err := ubluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userboardlike.Table, userboardlike.Columns, sqlgraph.NewFieldSpec(userboardlike.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(userboardlike.FieldBoardId, field.TypeInt))
	if id, ok := ubluo.mutation.UserId(); !ok {
		return nil, &ValidationError{Name: "userId", err: errors.New(`ent: missing "UserBoardLike.userId" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ubluo.mutation.BoardId(); !ok {
		return nil, &ValidationError{Name: "boardId", err: errors.New(`ent: missing "UserBoardLike.boardId" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ubluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !userboardlike.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ubluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ubluo.mutation.LikedAt(); ok {
		_spec.SetField(userboardlike.FieldLikedAt, field.TypeTime, value)
	}
	if ubluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.UserTable,
			Columns: []string{userboardlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.UserTable,
			Columns: []string{userboardlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ubluo.mutation.BoardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.BoardTable,
			Columns: []string{userboardlike.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubluo.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardlike.BoardTable,
			Columns: []string{userboardlike.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserBoardLike{config: ubluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ubluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userboardlike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ubluo.mutation.done = true
	return _node, nil
}
