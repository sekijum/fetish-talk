// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/comment"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentlike"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCommentLikeUpdate is the builder for updating UserCommentLike entities.
type UserCommentLikeUpdate struct {
	config
	hooks    []Hook
	mutation *UserCommentLikeMutation
}

// Where appends a list predicates to the UserCommentLikeUpdate builder.
func (uclu *UserCommentLikeUpdate) Where(ps ...predicate.UserCommentLike) *UserCommentLikeUpdate {
	uclu.mutation.Where(ps...)
	return uclu
}

// SetUserId sets the "userId" field.
func (uclu *UserCommentLikeUpdate) SetUserId(i int) *UserCommentLikeUpdate {
	uclu.mutation.SetUserId(i)
	return uclu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (uclu *UserCommentLikeUpdate) SetNillableUserId(i *int) *UserCommentLikeUpdate {
	if i != nil {
		uclu.SetUserId(*i)
	}
	return uclu
}

// SetCommentId sets the "commentId" field.
func (uclu *UserCommentLikeUpdate) SetCommentId(i int) *UserCommentLikeUpdate {
	uclu.mutation.SetCommentId(i)
	return uclu
}

// SetNillableCommentId sets the "commentId" field if the given value is not nil.
func (uclu *UserCommentLikeUpdate) SetNillableCommentId(i *int) *UserCommentLikeUpdate {
	if i != nil {
		uclu.SetCommentId(*i)
	}
	return uclu
}

// SetLikedAt sets the "likedAt" field.
func (uclu *UserCommentLikeUpdate) SetLikedAt(t time.Time) *UserCommentLikeUpdate {
	uclu.mutation.SetLikedAt(t)
	return uclu
}

// SetNillableLikedAt sets the "likedAt" field if the given value is not nil.
func (uclu *UserCommentLikeUpdate) SetNillableLikedAt(t *time.Time) *UserCommentLikeUpdate {
	if t != nil {
		uclu.SetLikedAt(*t)
	}
	return uclu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uclu *UserCommentLikeUpdate) SetUserID(id int) *UserCommentLikeUpdate {
	uclu.mutation.SetUserID(id)
	return uclu
}

// SetUser sets the "user" edge to the User entity.
func (uclu *UserCommentLikeUpdate) SetUser(u *User) *UserCommentLikeUpdate {
	return uclu.SetUserID(u.ID)
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (uclu *UserCommentLikeUpdate) SetCommentID(id int) *UserCommentLikeUpdate {
	uclu.mutation.SetCommentID(id)
	return uclu
}

// SetComment sets the "comment" edge to the Comment entity.
func (uclu *UserCommentLikeUpdate) SetComment(c *Comment) *UserCommentLikeUpdate {
	return uclu.SetCommentID(c.ID)
}

// Mutation returns the UserCommentLikeMutation object of the builder.
func (uclu *UserCommentLikeUpdate) Mutation() *UserCommentLikeMutation {
	return uclu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uclu *UserCommentLikeUpdate) ClearUser() *UserCommentLikeUpdate {
	uclu.mutation.ClearUser()
	return uclu
}

// ClearComment clears the "comment" edge to the Comment entity.
func (uclu *UserCommentLikeUpdate) ClearComment() *UserCommentLikeUpdate {
	uclu.mutation.ClearComment()
	return uclu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uclu *UserCommentLikeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uclu.sqlSave, uclu.mutation, uclu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uclu *UserCommentLikeUpdate) SaveX(ctx context.Context) int {
	affected, err := uclu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uclu *UserCommentLikeUpdate) Exec(ctx context.Context) error {
	_, err := uclu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uclu *UserCommentLikeUpdate) ExecX(ctx context.Context) {
	if err := uclu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uclu *UserCommentLikeUpdate) check() error {
	if _, ok := uclu.mutation.UserID(); uclu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.user"`)
	}
	if _, ok := uclu.mutation.CommentID(); uclu.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.comment"`)
	}
	return nil
}

func (uclu *UserCommentLikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uclu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercommentlike.Table, usercommentlike.Columns, sqlgraph.NewFieldSpec(usercommentlike.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(usercommentlike.FieldCommentId, field.TypeInt))
	if ps := uclu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uclu.mutation.LikedAt(); ok {
		_spec.SetField(usercommentlike.FieldLikedAt, field.TypeTime, value)
	}
	if uclu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uclu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
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
	if uclu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uclu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uclu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercommentlike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uclu.mutation.done = true
	return n, nil
}

// UserCommentLikeUpdateOne is the builder for updating a single UserCommentLike entity.
type UserCommentLikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCommentLikeMutation
}

// SetUserId sets the "userId" field.
func (ucluo *UserCommentLikeUpdateOne) SetUserId(i int) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetUserId(i)
	return ucluo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ucluo *UserCommentLikeUpdateOne) SetNillableUserId(i *int) *UserCommentLikeUpdateOne {
	if i != nil {
		ucluo.SetUserId(*i)
	}
	return ucluo
}

// SetCommentId sets the "commentId" field.
func (ucluo *UserCommentLikeUpdateOne) SetCommentId(i int) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetCommentId(i)
	return ucluo
}

// SetNillableCommentId sets the "commentId" field if the given value is not nil.
func (ucluo *UserCommentLikeUpdateOne) SetNillableCommentId(i *int) *UserCommentLikeUpdateOne {
	if i != nil {
		ucluo.SetCommentId(*i)
	}
	return ucluo
}

// SetLikedAt sets the "likedAt" field.
func (ucluo *UserCommentLikeUpdateOne) SetLikedAt(t time.Time) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetLikedAt(t)
	return ucluo
}

// SetNillableLikedAt sets the "likedAt" field if the given value is not nil.
func (ucluo *UserCommentLikeUpdateOne) SetNillableLikedAt(t *time.Time) *UserCommentLikeUpdateOne {
	if t != nil {
		ucluo.SetLikedAt(*t)
	}
	return ucluo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ucluo *UserCommentLikeUpdateOne) SetUserID(id int) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetUserID(id)
	return ucluo
}

// SetUser sets the "user" edge to the User entity.
func (ucluo *UserCommentLikeUpdateOne) SetUser(u *User) *UserCommentLikeUpdateOne {
	return ucluo.SetUserID(u.ID)
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (ucluo *UserCommentLikeUpdateOne) SetCommentID(id int) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetCommentID(id)
	return ucluo
}

// SetComment sets the "comment" edge to the Comment entity.
func (ucluo *UserCommentLikeUpdateOne) SetComment(c *Comment) *UserCommentLikeUpdateOne {
	return ucluo.SetCommentID(c.ID)
}

// Mutation returns the UserCommentLikeMutation object of the builder.
func (ucluo *UserCommentLikeUpdateOne) Mutation() *UserCommentLikeMutation {
	return ucluo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucluo *UserCommentLikeUpdateOne) ClearUser() *UserCommentLikeUpdateOne {
	ucluo.mutation.ClearUser()
	return ucluo
}

// ClearComment clears the "comment" edge to the Comment entity.
func (ucluo *UserCommentLikeUpdateOne) ClearComment() *UserCommentLikeUpdateOne {
	ucluo.mutation.ClearComment()
	return ucluo
}

// Where appends a list predicates to the UserCommentLikeUpdate builder.
func (ucluo *UserCommentLikeUpdateOne) Where(ps ...predicate.UserCommentLike) *UserCommentLikeUpdateOne {
	ucluo.mutation.Where(ps...)
	return ucluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucluo *UserCommentLikeUpdateOne) Select(field string, fields ...string) *UserCommentLikeUpdateOne {
	ucluo.fields = append([]string{field}, fields...)
	return ucluo
}

// Save executes the query and returns the updated UserCommentLike entity.
func (ucluo *UserCommentLikeUpdateOne) Save(ctx context.Context) (*UserCommentLike, error) {
	return withHooks(ctx, ucluo.sqlSave, ucluo.mutation, ucluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucluo *UserCommentLikeUpdateOne) SaveX(ctx context.Context) *UserCommentLike {
	node, err := ucluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucluo *UserCommentLikeUpdateOne) Exec(ctx context.Context) error {
	_, err := ucluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucluo *UserCommentLikeUpdateOne) ExecX(ctx context.Context) {
	if err := ucluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucluo *UserCommentLikeUpdateOne) check() error {
	if _, ok := ucluo.mutation.UserID(); ucluo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.user"`)
	}
	if _, ok := ucluo.mutation.CommentID(); ucluo.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.comment"`)
	}
	return nil
}

func (ucluo *UserCommentLikeUpdateOne) sqlSave(ctx context.Context) (_node *UserCommentLike, err error) {
	if err := ucluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercommentlike.Table, usercommentlike.Columns, sqlgraph.NewFieldSpec(usercommentlike.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(usercommentlike.FieldCommentId, field.TypeInt))
	if id, ok := ucluo.mutation.UserId(); !ok {
		return nil, &ValidationError{Name: "userId", err: errors.New(`ent: missing "UserCommentLike.userId" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ucluo.mutation.CommentId(); !ok {
		return nil, &ValidationError{Name: "commentId", err: errors.New(`ent: missing "UserCommentLike.commentId" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ucluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !usercommentlike.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ucluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucluo.mutation.LikedAt(); ok {
		_spec.SetField(usercommentlike.FieldLikedAt, field.TypeTime, value)
	}
	if ucluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
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
	if ucluo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucluo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCommentLike{config: ucluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercommentlike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucluo.mutation.done = true
	return _node, nil
}
