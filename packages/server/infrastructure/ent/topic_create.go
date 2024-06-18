// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/comment"
	"server/infrastructure/ent/forum"
	"server/infrastructure/ent/topic"
	"server/infrastructure/ent/topiclike"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usertopicnotification"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicCreate is the builder for creating a Topic entity.
type TopicCreate struct {
	config
	mutation *TopicMutation
	hooks    []Hook
}

// SetForumID sets the "forum_id" field.
func (tc *TopicCreate) SetForumID(i int) *TopicCreate {
	tc.mutation.SetForumID(i)
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TopicCreate) SetUserID(i int) *TopicCreate {
	tc.mutation.SetUserID(i)
	return tc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tc *TopicCreate) SetNillableUserID(i *int) *TopicCreate {
	if i != nil {
		tc.SetUserID(*i)
	}
	return tc
}

// SetTitle sets the "title" field.
func (tc *TopicCreate) SetTitle(s string) *TopicCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetContent sets the "content" field.
func (tc *TopicCreate) SetContent(s string) *TopicCreate {
	tc.mutation.SetContent(s)
	return tc
}

// SetIsDefault sets the "is_default" field.
func (tc *TopicCreate) SetIsDefault(b bool) *TopicCreate {
	tc.mutation.SetIsDefault(b)
	return tc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (tc *TopicCreate) SetNillableIsDefault(b *bool) *TopicCreate {
	if b != nil {
		tc.SetIsDefault(*b)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TopicCreate) SetStatus(t topic.Status) *TopicCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TopicCreate) SetNillableStatus(t *topic.Status) *TopicCreate {
	if t != nil {
		tc.SetStatus(*t)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TopicCreate) SetCreatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableCreatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TopicCreate) SetUpdatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableUpdatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TopicCreate) SetID(i int) *TopicCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetForum sets the "forum" edge to the Forum entity.
func (tc *TopicCreate) SetForum(f *Forum) *TopicCreate {
	return tc.SetForumID(f.ID)
}

// SetUser sets the "user" edge to the User entity.
func (tc *TopicCreate) SetUser(u *User) *TopicCreate {
	return tc.SetUserID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (tc *TopicCreate) AddCommentIDs(ids ...int) *TopicCreate {
	tc.mutation.AddCommentIDs(ids...)
	return tc
}

// AddComments adds the "comments" edges to the Comment entity.
func (tc *TopicCreate) AddComments(c ...*Comment) *TopicCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCommentIDs(ids...)
}

// AddTopicLikeIDs adds the "topic_likes" edge to the TopicLike entity by IDs.
func (tc *TopicCreate) AddTopicLikeIDs(ids ...int) *TopicCreate {
	tc.mutation.AddTopicLikeIDs(ids...)
	return tc
}

// AddTopicLikes adds the "topic_likes" edges to the TopicLike entity.
func (tc *TopicCreate) AddTopicLikes(t ...*TopicLike) *TopicCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTopicLikeIDs(ids...)
}

// AddUserTopicNotificationIDs adds the "user_topic_notifications" edge to the UserTopicNotification entity by IDs.
func (tc *TopicCreate) AddUserTopicNotificationIDs(ids ...int) *TopicCreate {
	tc.mutation.AddUserTopicNotificationIDs(ids...)
	return tc
}

// AddUserTopicNotifications adds the "user_topic_notifications" edges to the UserTopicNotification entity.
func (tc *TopicCreate) AddUserTopicNotifications(u ...*UserTopicNotification) *TopicCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUserTopicNotificationIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tc *TopicCreate) Mutation() *TopicMutation {
	return tc.mutation
}

// Save creates the Topic in the database.
func (tc *TopicCreate) Save(ctx context.Context) (*Topic, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TopicCreate) SaveX(ctx context.Context) *Topic {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TopicCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TopicCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TopicCreate) defaults() {
	if _, ok := tc.mutation.IsDefault(); !ok {
		v := topic.DefaultIsDefault
		tc.mutation.SetIsDefault(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := topic.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := topic.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := topic.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TopicCreate) check() error {
	if _, ok := tc.mutation.ForumID(); !ok {
		return &ValidationError{Name: "forum_id", err: errors.New(`ent: missing required field "Topic.forum_id"`)}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Topic.title"`)}
	}
	if _, ok := tc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Topic.content"`)}
	}
	if _, ok := tc.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`ent: missing required field "Topic.is_default"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Topic.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := topic.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Topic.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Topic.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Topic.updated_at"`)}
	}
	if _, ok := tc.mutation.ForumID(); !ok {
		return &ValidationError{Name: "forum", err: errors.New(`ent: missing required edge "Topic.forum"`)}
	}
	return nil
}

func (tc *TopicCreate) sqlSave(ctx context.Context) (*Topic, error) {
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

func (tc *TopicCreate) createSpec() (*Topic, *sqlgraph.CreateSpec) {
	var (
		_node = &Topic{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(topic.Table, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(topic.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Content(); ok {
		_spec.SetField(topic.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := tc.mutation.IsDefault(); ok {
		_spec.SetField(topic.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(topic.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(topic.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.ForumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.ForumTable,
			Columns: []string{topic.ForumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ForumID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.UserTable,
			Columns: []string{topic.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.CommentsTable,
			Columns: []string{topic.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TopicLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicLikesTable,
			Columns: []string{topic.TopicLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topiclike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserTopicNotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.UserTopicNotificationsTable,
			Columns: []string{topic.UserTopicNotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertopicnotification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TopicCreateBulk is the builder for creating many Topic entities in bulk.
type TopicCreateBulk struct {
	config
	err      error
	builders []*TopicCreate
}

// Save creates the Topic entities in the database.
func (tcb *TopicCreateBulk) Save(ctx context.Context) ([]*Topic, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Topic, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopicMutation)
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
func (tcb *TopicCreateBulk) SaveX(ctx context.Context) []*Topic {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TopicCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TopicCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
