// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/forum"
	"server/infrastructure/ent/forumlike"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/topic"
	"server/infrastructure/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ForumUpdate is the builder for updating Forum entities.
type ForumUpdate struct {
	config
	hooks    []Hook
	mutation *ForumMutation
}

// Where appends a list predicates to the ForumUpdate builder.
func (fu *ForumUpdate) Where(ps ...predicate.Forum) *ForumUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *ForumUpdate) SetUserID(i int) *ForumUpdate {
	fu.mutation.SetUserID(i)
	return fu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fu *ForumUpdate) SetNillableUserID(i *int) *ForumUpdate {
	if i != nil {
		fu.SetUserID(*i)
	}
	return fu
}

// SetName sets the "name" field.
func (fu *ForumUpdate) SetName(s string) *ForumUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *ForumUpdate) SetNillableName(s *string) *ForumUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetDescription sets the "description" field.
func (fu *ForumUpdate) SetDescription(s string) *ForumUpdate {
	fu.mutation.SetDescription(s)
	return fu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fu *ForumUpdate) SetNillableDescription(s *string) *ForumUpdate {
	if s != nil {
		fu.SetDescription(*s)
	}
	return fu
}

// ClearDescription clears the value of the "description" field.
func (fu *ForumUpdate) ClearDescription() *ForumUpdate {
	fu.mutation.ClearDescription()
	return fu
}

// SetStatus sets the "status" field.
func (fu *ForumUpdate) SetStatus(f forum.Status) *ForumUpdate {
	fu.mutation.SetStatus(f)
	return fu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fu *ForumUpdate) SetNillableStatus(f *forum.Status) *ForumUpdate {
	if f != nil {
		fu.SetStatus(*f)
	}
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *ForumUpdate) SetCreatedAt(t time.Time) *ForumUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *ForumUpdate) SetNillableCreatedAt(t *time.Time) *ForumUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *ForumUpdate) SetUpdatedAt(t time.Time) *ForumUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetUser sets the "user" edge to the User entity.
func (fu *ForumUpdate) SetUser(u *User) *ForumUpdate {
	return fu.SetUserID(u.ID)
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (fu *ForumUpdate) AddTopicIDs(ids ...int) *ForumUpdate {
	fu.mutation.AddTopicIDs(ids...)
	return fu
}

// AddTopics adds the "topics" edges to the Topic entity.
func (fu *ForumUpdate) AddTopics(t ...*Topic) *ForumUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fu.AddTopicIDs(ids...)
}

// AddForumLikeIDs adds the "forum_likes" edge to the ForumLike entity by IDs.
func (fu *ForumUpdate) AddForumLikeIDs(ids ...int) *ForumUpdate {
	fu.mutation.AddForumLikeIDs(ids...)
	return fu
}

// AddForumLikes adds the "forum_likes" edges to the ForumLike entity.
func (fu *ForumUpdate) AddForumLikes(f ...*ForumLike) *ForumUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddForumLikeIDs(ids...)
}

// Mutation returns the ForumMutation object of the builder.
func (fu *ForumUpdate) Mutation() *ForumMutation {
	return fu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fu *ForumUpdate) ClearUser() *ForumUpdate {
	fu.mutation.ClearUser()
	return fu
}

// ClearTopics clears all "topics" edges to the Topic entity.
func (fu *ForumUpdate) ClearTopics() *ForumUpdate {
	fu.mutation.ClearTopics()
	return fu
}

// RemoveTopicIDs removes the "topics" edge to Topic entities by IDs.
func (fu *ForumUpdate) RemoveTopicIDs(ids ...int) *ForumUpdate {
	fu.mutation.RemoveTopicIDs(ids...)
	return fu
}

// RemoveTopics removes "topics" edges to Topic entities.
func (fu *ForumUpdate) RemoveTopics(t ...*Topic) *ForumUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fu.RemoveTopicIDs(ids...)
}

// ClearForumLikes clears all "forum_likes" edges to the ForumLike entity.
func (fu *ForumUpdate) ClearForumLikes() *ForumUpdate {
	fu.mutation.ClearForumLikes()
	return fu
}

// RemoveForumLikeIDs removes the "forum_likes" edge to ForumLike entities by IDs.
func (fu *ForumUpdate) RemoveForumLikeIDs(ids ...int) *ForumUpdate {
	fu.mutation.RemoveForumLikeIDs(ids...)
	return fu
}

// RemoveForumLikes removes "forum_likes" edges to ForumLike entities.
func (fu *ForumUpdate) RemoveForumLikes(f ...*ForumLike) *ForumUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveForumLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *ForumUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *ForumUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *ForumUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *ForumUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *ForumUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := forum.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *ForumUpdate) check() error {
	if v, ok := fu.mutation.Status(); ok {
		if err := forum.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Forum.status": %w`, err)}
		}
	}
	if _, ok := fu.mutation.UserID(); fu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Forum.user"`)
	}
	return nil
}

func (fu *ForumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(forum.Table, forum.Columns, sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(forum.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.Description(); ok {
		_spec.SetField(forum.FieldDescription, field.TypeString, value)
	}
	if fu.mutation.DescriptionCleared() {
		_spec.ClearField(forum.FieldDescription, field.TypeString)
	}
	if value, ok := fu.mutation.Status(); ok {
		_spec.SetField(forum.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.SetField(forum.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(forum.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forum.UserTable,
			Columns: []string{forum.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forum.UserTable,
			Columns: []string{forum.UserColumn},
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
	if fu.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.TopicsTable,
			Columns: []string{forum.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedTopicsIDs(); len(nodes) > 0 && !fu.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.TopicsTable,
			Columns: []string{forum.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.TopicsTable,
			Columns: []string{forum.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.ForumLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.ForumLikesTable,
			Columns: []string{forum.ForumLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedForumLikesIDs(); len(nodes) > 0 && !fu.mutation.ForumLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.ForumLikesTable,
			Columns: []string{forum.ForumLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ForumLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.ForumLikesTable,
			Columns: []string{forum.ForumLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{forum.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// ForumUpdateOne is the builder for updating a single Forum entity.
type ForumUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ForumMutation
}

// SetUserID sets the "user_id" field.
func (fuo *ForumUpdateOne) SetUserID(i int) *ForumUpdateOne {
	fuo.mutation.SetUserID(i)
	return fuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fuo *ForumUpdateOne) SetNillableUserID(i *int) *ForumUpdateOne {
	if i != nil {
		fuo.SetUserID(*i)
	}
	return fuo
}

// SetName sets the "name" field.
func (fuo *ForumUpdateOne) SetName(s string) *ForumUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *ForumUpdateOne) SetNillableName(s *string) *ForumUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetDescription sets the "description" field.
func (fuo *ForumUpdateOne) SetDescription(s string) *ForumUpdateOne {
	fuo.mutation.SetDescription(s)
	return fuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fuo *ForumUpdateOne) SetNillableDescription(s *string) *ForumUpdateOne {
	if s != nil {
		fuo.SetDescription(*s)
	}
	return fuo
}

// ClearDescription clears the value of the "description" field.
func (fuo *ForumUpdateOne) ClearDescription() *ForumUpdateOne {
	fuo.mutation.ClearDescription()
	return fuo
}

// SetStatus sets the "status" field.
func (fuo *ForumUpdateOne) SetStatus(f forum.Status) *ForumUpdateOne {
	fuo.mutation.SetStatus(f)
	return fuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fuo *ForumUpdateOne) SetNillableStatus(f *forum.Status) *ForumUpdateOne {
	if f != nil {
		fuo.SetStatus(*f)
	}
	return fuo
}

// SetCreatedAt sets the "created_at" field.
func (fuo *ForumUpdateOne) SetCreatedAt(t time.Time) *ForumUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *ForumUpdateOne) SetNillableCreatedAt(t *time.Time) *ForumUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *ForumUpdateOne) SetUpdatedAt(t time.Time) *ForumUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetUser sets the "user" edge to the User entity.
func (fuo *ForumUpdateOne) SetUser(u *User) *ForumUpdateOne {
	return fuo.SetUserID(u.ID)
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (fuo *ForumUpdateOne) AddTopicIDs(ids ...int) *ForumUpdateOne {
	fuo.mutation.AddTopicIDs(ids...)
	return fuo
}

// AddTopics adds the "topics" edges to the Topic entity.
func (fuo *ForumUpdateOne) AddTopics(t ...*Topic) *ForumUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fuo.AddTopicIDs(ids...)
}

// AddForumLikeIDs adds the "forum_likes" edge to the ForumLike entity by IDs.
func (fuo *ForumUpdateOne) AddForumLikeIDs(ids ...int) *ForumUpdateOne {
	fuo.mutation.AddForumLikeIDs(ids...)
	return fuo
}

// AddForumLikes adds the "forum_likes" edges to the ForumLike entity.
func (fuo *ForumUpdateOne) AddForumLikes(f ...*ForumLike) *ForumUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddForumLikeIDs(ids...)
}

// Mutation returns the ForumMutation object of the builder.
func (fuo *ForumUpdateOne) Mutation() *ForumMutation {
	return fuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fuo *ForumUpdateOne) ClearUser() *ForumUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// ClearTopics clears all "topics" edges to the Topic entity.
func (fuo *ForumUpdateOne) ClearTopics() *ForumUpdateOne {
	fuo.mutation.ClearTopics()
	return fuo
}

// RemoveTopicIDs removes the "topics" edge to Topic entities by IDs.
func (fuo *ForumUpdateOne) RemoveTopicIDs(ids ...int) *ForumUpdateOne {
	fuo.mutation.RemoveTopicIDs(ids...)
	return fuo
}

// RemoveTopics removes "topics" edges to Topic entities.
func (fuo *ForumUpdateOne) RemoveTopics(t ...*Topic) *ForumUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fuo.RemoveTopicIDs(ids...)
}

// ClearForumLikes clears all "forum_likes" edges to the ForumLike entity.
func (fuo *ForumUpdateOne) ClearForumLikes() *ForumUpdateOne {
	fuo.mutation.ClearForumLikes()
	return fuo
}

// RemoveForumLikeIDs removes the "forum_likes" edge to ForumLike entities by IDs.
func (fuo *ForumUpdateOne) RemoveForumLikeIDs(ids ...int) *ForumUpdateOne {
	fuo.mutation.RemoveForumLikeIDs(ids...)
	return fuo
}

// RemoveForumLikes removes "forum_likes" edges to ForumLike entities.
func (fuo *ForumUpdateOne) RemoveForumLikes(f ...*ForumLike) *ForumUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveForumLikeIDs(ids...)
}

// Where appends a list predicates to the ForumUpdate builder.
func (fuo *ForumUpdateOne) Where(ps ...predicate.Forum) *ForumUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *ForumUpdateOne) Select(field string, fields ...string) *ForumUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Forum entity.
func (fuo *ForumUpdateOne) Save(ctx context.Context) (*Forum, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *ForumUpdateOne) SaveX(ctx context.Context) *Forum {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *ForumUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *ForumUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *ForumUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := forum.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *ForumUpdateOne) check() error {
	if v, ok := fuo.mutation.Status(); ok {
		if err := forum.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Forum.status": %w`, err)}
		}
	}
	if _, ok := fuo.mutation.UserID(); fuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Forum.user"`)
	}
	return nil
}

func (fuo *ForumUpdateOne) sqlSave(ctx context.Context) (_node *Forum, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(forum.Table, forum.Columns, sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Forum.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, forum.FieldID)
		for _, f := range fields {
			if !forum.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != forum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(forum.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Description(); ok {
		_spec.SetField(forum.FieldDescription, field.TypeString, value)
	}
	if fuo.mutation.DescriptionCleared() {
		_spec.ClearField(forum.FieldDescription, field.TypeString)
	}
	if value, ok := fuo.mutation.Status(); ok {
		_spec.SetField(forum.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.SetField(forum.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(forum.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forum.UserTable,
			Columns: []string{forum.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forum.UserTable,
			Columns: []string{forum.UserColumn},
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
	if fuo.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.TopicsTable,
			Columns: []string{forum.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedTopicsIDs(); len(nodes) > 0 && !fuo.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.TopicsTable,
			Columns: []string{forum.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.TopicsTable,
			Columns: []string{forum.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.ForumLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.ForumLikesTable,
			Columns: []string{forum.ForumLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedForumLikesIDs(); len(nodes) > 0 && !fuo.mutation.ForumLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.ForumLikesTable,
			Columns: []string{forum.ForumLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ForumLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.ForumLikesTable,
			Columns: []string{forum.ForumLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Forum{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{forum.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
