// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThreadCommentAttachmentCreate is the builder for creating a ThreadCommentAttachment entity.
type ThreadCommentAttachmentCreate struct {
	config
	mutation *ThreadCommentAttachmentMutation
	hooks    []Hook
}

// SetCommentId sets the "commentId" field.
func (tcac *ThreadCommentAttachmentCreate) SetCommentId(i int) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetCommentId(i)
	return tcac
}

// SetURL sets the "url" field.
func (tcac *ThreadCommentAttachmentCreate) SetURL(s string) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetURL(s)
	return tcac
}

// SetDisplayOrder sets the "display_order" field.
func (tcac *ThreadCommentAttachmentCreate) SetDisplayOrder(i int) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetDisplayOrder(i)
	return tcac
}

// SetNillableDisplayOrder sets the "display_order" field if the given value is not nil.
func (tcac *ThreadCommentAttachmentCreate) SetNillableDisplayOrder(i *int) *ThreadCommentAttachmentCreate {
	if i != nil {
		tcac.SetDisplayOrder(*i)
	}
	return tcac
}

// SetType sets the "type" field.
func (tcac *ThreadCommentAttachmentCreate) SetType(i int) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetType(i)
	return tcac
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tcac *ThreadCommentAttachmentCreate) SetNillableType(i *int) *ThreadCommentAttachmentCreate {
	if i != nil {
		tcac.SetType(*i)
	}
	return tcac
}

// SetCreatedAt sets the "createdAt" field.
func (tcac *ThreadCommentAttachmentCreate) SetCreatedAt(t time.Time) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetCreatedAt(t)
	return tcac
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tcac *ThreadCommentAttachmentCreate) SetNillableCreatedAt(t *time.Time) *ThreadCommentAttachmentCreate {
	if t != nil {
		tcac.SetCreatedAt(*t)
	}
	return tcac
}

// SetID sets the "id" field.
func (tcac *ThreadCommentAttachmentCreate) SetID(i int) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetID(i)
	return tcac
}

// SetCommentID sets the "comment" edge to the ThreadComment entity by ID.
func (tcac *ThreadCommentAttachmentCreate) SetCommentID(id int) *ThreadCommentAttachmentCreate {
	tcac.mutation.SetCommentID(id)
	return tcac
}

// SetComment sets the "comment" edge to the ThreadComment entity.
func (tcac *ThreadCommentAttachmentCreate) SetComment(t *ThreadComment) *ThreadCommentAttachmentCreate {
	return tcac.SetCommentID(t.ID)
}

// Mutation returns the ThreadCommentAttachmentMutation object of the builder.
func (tcac *ThreadCommentAttachmentCreate) Mutation() *ThreadCommentAttachmentMutation {
	return tcac.mutation
}

// Save creates the ThreadCommentAttachment in the database.
func (tcac *ThreadCommentAttachmentCreate) Save(ctx context.Context) (*ThreadCommentAttachment, error) {
	tcac.defaults()
	return withHooks(ctx, tcac.sqlSave, tcac.mutation, tcac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tcac *ThreadCommentAttachmentCreate) SaveX(ctx context.Context) *ThreadCommentAttachment {
	v, err := tcac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcac *ThreadCommentAttachmentCreate) Exec(ctx context.Context) error {
	_, err := tcac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcac *ThreadCommentAttachmentCreate) ExecX(ctx context.Context) {
	if err := tcac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcac *ThreadCommentAttachmentCreate) defaults() {
	if _, ok := tcac.mutation.DisplayOrder(); !ok {
		v := threadcommentattachment.DefaultDisplayOrder
		tcac.mutation.SetDisplayOrder(v)
	}
	if _, ok := tcac.mutation.GetType(); !ok {
		v := threadcommentattachment.DefaultType
		tcac.mutation.SetType(v)
	}
	if _, ok := tcac.mutation.CreatedAt(); !ok {
		v := threadcommentattachment.DefaultCreatedAt()
		tcac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcac *ThreadCommentAttachmentCreate) check() error {
	if _, ok := tcac.mutation.CommentId(); !ok {
		return &ValidationError{Name: "commentId", err: errors.New(`ent: missing required field "ThreadCommentAttachment.commentId"`)}
	}
	if _, ok := tcac.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "ThreadCommentAttachment.url"`)}
	}
	if _, ok := tcac.mutation.DisplayOrder(); !ok {
		return &ValidationError{Name: "display_order", err: errors.New(`ent: missing required field "ThreadCommentAttachment.display_order"`)}
	}
	if _, ok := tcac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ThreadCommentAttachment.type"`)}
	}
	if _, ok := tcac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "ThreadCommentAttachment.createdAt"`)}
	}
	if _, ok := tcac.mutation.CommentID(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required edge "ThreadCommentAttachment.comment"`)}
	}
	return nil
}

func (tcac *ThreadCommentAttachmentCreate) sqlSave(ctx context.Context) (*ThreadCommentAttachment, error) {
	if err := tcac.check(); err != nil {
		return nil, err
	}
	_node, _spec := tcac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tcac.mutation.id = &_node.ID
	tcac.mutation.done = true
	return _node, nil
}

func (tcac *ThreadCommentAttachmentCreate) createSpec() (*ThreadCommentAttachment, *sqlgraph.CreateSpec) {
	var (
		_node = &ThreadCommentAttachment{config: tcac.config}
		_spec = sqlgraph.NewCreateSpec(threadcommentattachment.Table, sqlgraph.NewFieldSpec(threadcommentattachment.FieldID, field.TypeInt))
	)
	if id, ok := tcac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tcac.mutation.URL(); ok {
		_spec.SetField(threadcommentattachment.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := tcac.mutation.DisplayOrder(); ok {
		_spec.SetField(threadcommentattachment.FieldDisplayOrder, field.TypeInt, value)
		_node.DisplayOrder = value
	}
	if value, ok := tcac.mutation.GetType(); ok {
		_spec.SetField(threadcommentattachment.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := tcac.mutation.CreatedAt(); ok {
		_spec.SetField(threadcommentattachment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := tcac.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   threadcommentattachment.CommentTable,
			Columns: []string{threadcommentattachment.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommentId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThreadCommentAttachmentCreateBulk is the builder for creating many ThreadCommentAttachment entities in bulk.
type ThreadCommentAttachmentCreateBulk struct {
	config
	err      error
	builders []*ThreadCommentAttachmentCreate
}

// Save creates the ThreadCommentAttachment entities in the database.
func (tcacb *ThreadCommentAttachmentCreateBulk) Save(ctx context.Context) ([]*ThreadCommentAttachment, error) {
	if tcacb.err != nil {
		return nil, tcacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcacb.builders))
	nodes := make([]*ThreadCommentAttachment, len(tcacb.builders))
	mutators := make([]Mutator, len(tcacb.builders))
	for i := range tcacb.builders {
		func(i int, root context.Context) {
			builder := tcacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadCommentAttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, tcacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcacb *ThreadCommentAttachmentCreateBulk) SaveX(ctx context.Context) []*ThreadCommentAttachment {
	v, err := tcacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcacb *ThreadCommentAttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := tcacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcacb *ThreadCommentAttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := tcacb.Exec(ctx); err != nil {
		panic(err)
	}
}