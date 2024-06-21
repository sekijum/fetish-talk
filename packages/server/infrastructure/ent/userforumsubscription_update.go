// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/forum"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userforumsubscription"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserForumSubscriptionUpdate is the builder for updating UserForumSubscription entities.
type UserForumSubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *UserForumSubscriptionMutation
}

// Where appends a list predicates to the UserForumSubscriptionUpdate builder.
func (ufsu *UserForumSubscriptionUpdate) Where(ps ...predicate.UserForumSubscription) *UserForumSubscriptionUpdate {
	ufsu.mutation.Where(ps...)
	return ufsu
}

// SetUserId sets the "userId" field.
func (ufsu *UserForumSubscriptionUpdate) SetUserId(i int) *UserForumSubscriptionUpdate {
	ufsu.mutation.SetUserId(i)
	return ufsu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ufsu *UserForumSubscriptionUpdate) SetNillableUserId(i *int) *UserForumSubscriptionUpdate {
	if i != nil {
		ufsu.SetUserId(*i)
	}
	return ufsu
}

// SetForumId sets the "forumId" field.
func (ufsu *UserForumSubscriptionUpdate) SetForumId(i int) *UserForumSubscriptionUpdate {
	ufsu.mutation.SetForumId(i)
	return ufsu
}

// SetNillableForumId sets the "forumId" field if the given value is not nil.
func (ufsu *UserForumSubscriptionUpdate) SetNillableForumId(i *int) *UserForumSubscriptionUpdate {
	if i != nil {
		ufsu.SetForumId(*i)
	}
	return ufsu
}

// SetIsNotified sets the "isNotified" field.
func (ufsu *UserForumSubscriptionUpdate) SetIsNotified(b bool) *UserForumSubscriptionUpdate {
	ufsu.mutation.SetIsNotified(b)
	return ufsu
}

// SetNillableIsNotified sets the "isNotified" field if the given value is not nil.
func (ufsu *UserForumSubscriptionUpdate) SetNillableIsNotified(b *bool) *UserForumSubscriptionUpdate {
	if b != nil {
		ufsu.SetIsNotified(*b)
	}
	return ufsu
}

// SetSubscribedAt sets the "subscribedAt" field.
func (ufsu *UserForumSubscriptionUpdate) SetSubscribedAt(t time.Time) *UserForumSubscriptionUpdate {
	ufsu.mutation.SetSubscribedAt(t)
	return ufsu
}

// SetNillableSubscribedAt sets the "subscribedAt" field if the given value is not nil.
func (ufsu *UserForumSubscriptionUpdate) SetNillableSubscribedAt(t *time.Time) *UserForumSubscriptionUpdate {
	if t != nil {
		ufsu.SetSubscribedAt(*t)
	}
	return ufsu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ufsu *UserForumSubscriptionUpdate) SetUserID(id int) *UserForumSubscriptionUpdate {
	ufsu.mutation.SetUserID(id)
	return ufsu
}

// SetUser sets the "user" edge to the User entity.
func (ufsu *UserForumSubscriptionUpdate) SetUser(u *User) *UserForumSubscriptionUpdate {
	return ufsu.SetUserID(u.ID)
}

// SetForumID sets the "forum" edge to the Forum entity by ID.
func (ufsu *UserForumSubscriptionUpdate) SetForumID(id int) *UserForumSubscriptionUpdate {
	ufsu.mutation.SetForumID(id)
	return ufsu
}

// SetForum sets the "forum" edge to the Forum entity.
func (ufsu *UserForumSubscriptionUpdate) SetForum(f *Forum) *UserForumSubscriptionUpdate {
	return ufsu.SetForumID(f.ID)
}

// Mutation returns the UserForumSubscriptionMutation object of the builder.
func (ufsu *UserForumSubscriptionUpdate) Mutation() *UserForumSubscriptionMutation {
	return ufsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ufsu *UserForumSubscriptionUpdate) ClearUser() *UserForumSubscriptionUpdate {
	ufsu.mutation.ClearUser()
	return ufsu
}

// ClearForum clears the "forum" edge to the Forum entity.
func (ufsu *UserForumSubscriptionUpdate) ClearForum() *UserForumSubscriptionUpdate {
	ufsu.mutation.ClearForum()
	return ufsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ufsu *UserForumSubscriptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ufsu.sqlSave, ufsu.mutation, ufsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufsu *UserForumSubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := ufsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ufsu *UserForumSubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := ufsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufsu *UserForumSubscriptionUpdate) ExecX(ctx context.Context) {
	if err := ufsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufsu *UserForumSubscriptionUpdate) check() error {
	if _, ok := ufsu.mutation.UserID(); ufsu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserForumSubscription.user"`)
	}
	if _, ok := ufsu.mutation.ForumID(); ufsu.mutation.ForumCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserForumSubscription.forum"`)
	}
	return nil
}

func (ufsu *UserForumSubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ufsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userforumsubscription.Table, userforumsubscription.Columns, sqlgraph.NewFieldSpec(userforumsubscription.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(userforumsubscription.FieldForumId, field.TypeInt))
	if ps := ufsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufsu.mutation.IsNotified(); ok {
		_spec.SetField(userforumsubscription.FieldIsNotified, field.TypeBool, value)
	}
	if value, ok := ufsu.mutation.SubscribedAt(); ok {
		_spec.SetField(userforumsubscription.FieldSubscribedAt, field.TypeTime, value)
	}
	if ufsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.UserTable,
			Columns: []string{userforumsubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.UserTable,
			Columns: []string{userforumsubscription.UserColumn},
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
	if ufsu.mutation.ForumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.ForumTable,
			Columns: []string{userforumsubscription.ForumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufsu.mutation.ForumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.ForumTable,
			Columns: []string{userforumsubscription.ForumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ufsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userforumsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ufsu.mutation.done = true
	return n, nil
}

// UserForumSubscriptionUpdateOne is the builder for updating a single UserForumSubscription entity.
type UserForumSubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserForumSubscriptionMutation
}

// SetUserId sets the "userId" field.
func (ufsuo *UserForumSubscriptionUpdateOne) SetUserId(i int) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.SetUserId(i)
	return ufsuo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ufsuo *UserForumSubscriptionUpdateOne) SetNillableUserId(i *int) *UserForumSubscriptionUpdateOne {
	if i != nil {
		ufsuo.SetUserId(*i)
	}
	return ufsuo
}

// SetForumId sets the "forumId" field.
func (ufsuo *UserForumSubscriptionUpdateOne) SetForumId(i int) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.SetForumId(i)
	return ufsuo
}

// SetNillableForumId sets the "forumId" field if the given value is not nil.
func (ufsuo *UserForumSubscriptionUpdateOne) SetNillableForumId(i *int) *UserForumSubscriptionUpdateOne {
	if i != nil {
		ufsuo.SetForumId(*i)
	}
	return ufsuo
}

// SetIsNotified sets the "isNotified" field.
func (ufsuo *UserForumSubscriptionUpdateOne) SetIsNotified(b bool) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.SetIsNotified(b)
	return ufsuo
}

// SetNillableIsNotified sets the "isNotified" field if the given value is not nil.
func (ufsuo *UserForumSubscriptionUpdateOne) SetNillableIsNotified(b *bool) *UserForumSubscriptionUpdateOne {
	if b != nil {
		ufsuo.SetIsNotified(*b)
	}
	return ufsuo
}

// SetSubscribedAt sets the "subscribedAt" field.
func (ufsuo *UserForumSubscriptionUpdateOne) SetSubscribedAt(t time.Time) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.SetSubscribedAt(t)
	return ufsuo
}

// SetNillableSubscribedAt sets the "subscribedAt" field if the given value is not nil.
func (ufsuo *UserForumSubscriptionUpdateOne) SetNillableSubscribedAt(t *time.Time) *UserForumSubscriptionUpdateOne {
	if t != nil {
		ufsuo.SetSubscribedAt(*t)
	}
	return ufsuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ufsuo *UserForumSubscriptionUpdateOne) SetUserID(id int) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.SetUserID(id)
	return ufsuo
}

// SetUser sets the "user" edge to the User entity.
func (ufsuo *UserForumSubscriptionUpdateOne) SetUser(u *User) *UserForumSubscriptionUpdateOne {
	return ufsuo.SetUserID(u.ID)
}

// SetForumID sets the "forum" edge to the Forum entity by ID.
func (ufsuo *UserForumSubscriptionUpdateOne) SetForumID(id int) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.SetForumID(id)
	return ufsuo
}

// SetForum sets the "forum" edge to the Forum entity.
func (ufsuo *UserForumSubscriptionUpdateOne) SetForum(f *Forum) *UserForumSubscriptionUpdateOne {
	return ufsuo.SetForumID(f.ID)
}

// Mutation returns the UserForumSubscriptionMutation object of the builder.
func (ufsuo *UserForumSubscriptionUpdateOne) Mutation() *UserForumSubscriptionMutation {
	return ufsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ufsuo *UserForumSubscriptionUpdateOne) ClearUser() *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.ClearUser()
	return ufsuo
}

// ClearForum clears the "forum" edge to the Forum entity.
func (ufsuo *UserForumSubscriptionUpdateOne) ClearForum() *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.ClearForum()
	return ufsuo
}

// Where appends a list predicates to the UserForumSubscriptionUpdate builder.
func (ufsuo *UserForumSubscriptionUpdateOne) Where(ps ...predicate.UserForumSubscription) *UserForumSubscriptionUpdateOne {
	ufsuo.mutation.Where(ps...)
	return ufsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ufsuo *UserForumSubscriptionUpdateOne) Select(field string, fields ...string) *UserForumSubscriptionUpdateOne {
	ufsuo.fields = append([]string{field}, fields...)
	return ufsuo
}

// Save executes the query and returns the updated UserForumSubscription entity.
func (ufsuo *UserForumSubscriptionUpdateOne) Save(ctx context.Context) (*UserForumSubscription, error) {
	return withHooks(ctx, ufsuo.sqlSave, ufsuo.mutation, ufsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufsuo *UserForumSubscriptionUpdateOne) SaveX(ctx context.Context) *UserForumSubscription {
	node, err := ufsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ufsuo *UserForumSubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := ufsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufsuo *UserForumSubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := ufsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufsuo *UserForumSubscriptionUpdateOne) check() error {
	if _, ok := ufsuo.mutation.UserID(); ufsuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserForumSubscription.user"`)
	}
	if _, ok := ufsuo.mutation.ForumID(); ufsuo.mutation.ForumCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserForumSubscription.forum"`)
	}
	return nil
}

func (ufsuo *UserForumSubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *UserForumSubscription, err error) {
	if err := ufsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userforumsubscription.Table, userforumsubscription.Columns, sqlgraph.NewFieldSpec(userforumsubscription.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(userforumsubscription.FieldForumId, field.TypeInt))
	if id, ok := ufsuo.mutation.UserId(); !ok {
		return nil, &ValidationError{Name: "userId", err: errors.New(`ent: missing "UserForumSubscription.userId" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ufsuo.mutation.ForumId(); !ok {
		return nil, &ValidationError{Name: "forumId", err: errors.New(`ent: missing "UserForumSubscription.forumId" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ufsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !userforumsubscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ufsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufsuo.mutation.IsNotified(); ok {
		_spec.SetField(userforumsubscription.FieldIsNotified, field.TypeBool, value)
	}
	if value, ok := ufsuo.mutation.SubscribedAt(); ok {
		_spec.SetField(userforumsubscription.FieldSubscribedAt, field.TypeTime, value)
	}
	if ufsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.UserTable,
			Columns: []string{userforumsubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.UserTable,
			Columns: []string{userforumsubscription.UserColumn},
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
	if ufsuo.mutation.ForumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.ForumTable,
			Columns: []string{userforumsubscription.ForumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufsuo.mutation.ForumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userforumsubscription.ForumTable,
			Columns: []string{userforumsubscription.ForumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserForumSubscription{config: ufsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ufsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userforumsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ufsuo.mutation.done = true
	return _node, nil
}
