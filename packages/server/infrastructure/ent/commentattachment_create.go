// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/comment"
	"server/infrastructure/ent/commentattachment"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentAttachmentCreate is the builder for creating a CommentAttachment entity.
type CommentAttachmentCreate struct {
	config
	mutation *CommentAttachmentMutation
	hooks    []Hook
}

// SetCommentId sets the "commentId" field.
func (cac *CommentAttachmentCreate) SetCommentId(i int) *CommentAttachmentCreate {
	cac.mutation.SetCommentId(i)
	return cac
}

// SetURL sets the "url" field.
func (cac *CommentAttachmentCreate) SetURL(s string) *CommentAttachmentCreate {
	cac.mutation.SetURL(s)
	return cac
}

// SetOrder sets the "order" field.
func (cac *CommentAttachmentCreate) SetOrder(i int) *CommentAttachmentCreate {
	cac.mutation.SetOrder(i)
	return cac
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (cac *CommentAttachmentCreate) SetNillableOrder(i *int) *CommentAttachmentCreate {
	if i != nil {
		cac.SetOrder(*i)
	}
	return cac
}

// SetType sets the "type" field.
func (cac *CommentAttachmentCreate) SetType(c commentattachment.Type) *CommentAttachmentCreate {
	cac.mutation.SetType(c)
	return cac
}

// SetID sets the "id" field.
func (cac *CommentAttachmentCreate) SetID(i int) *CommentAttachmentCreate {
	cac.mutation.SetID(i)
	return cac
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (cac *CommentAttachmentCreate) SetCommentID(id int) *CommentAttachmentCreate {
	cac.mutation.SetCommentID(id)
	return cac
}

// SetComment sets the "comment" edge to the Comment entity.
func (cac *CommentAttachmentCreate) SetComment(c *Comment) *CommentAttachmentCreate {
	return cac.SetCommentID(c.ID)
}

// Mutation returns the CommentAttachmentMutation object of the builder.
func (cac *CommentAttachmentCreate) Mutation() *CommentAttachmentMutation {
	return cac.mutation
}

// Save creates the CommentAttachment in the database.
func (cac *CommentAttachmentCreate) Save(ctx context.Context) (*CommentAttachment, error) {
	cac.defaults()
	return withHooks(ctx, cac.sqlSave, cac.mutation, cac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cac *CommentAttachmentCreate) SaveX(ctx context.Context) *CommentAttachment {
	v, err := cac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cac *CommentAttachmentCreate) Exec(ctx context.Context) error {
	_, err := cac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cac *CommentAttachmentCreate) ExecX(ctx context.Context) {
	if err := cac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cac *CommentAttachmentCreate) defaults() {
	if _, ok := cac.mutation.Order(); !ok {
		v := commentattachment.DefaultOrder
		cac.mutation.SetOrder(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cac *CommentAttachmentCreate) check() error {
	if _, ok := cac.mutation.CommentId(); !ok {
		return &ValidationError{Name: "commentId", err: errors.New(`ent: missing required field "CommentAttachment.commentId"`)}
	}
	if _, ok := cac.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "CommentAttachment.url"`)}
	}
	if _, ok := cac.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "CommentAttachment.order"`)}
	}
	if _, ok := cac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "CommentAttachment.type"`)}
	}
	if v, ok := cac.mutation.GetType(); ok {
		if err := commentattachment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CommentAttachment.type": %w`, err)}
		}
	}
	if _, ok := cac.mutation.CommentID(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required edge "CommentAttachment.comment"`)}
	}
	return nil
}

func (cac *CommentAttachmentCreate) sqlSave(ctx context.Context) (*CommentAttachment, error) {
	if err := cac.check(); err != nil {
		return nil, err
	}
	_node, _spec := cac.createSpec()
	if err := sqlgraph.CreateNode(ctx, cac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cac.mutation.id = &_node.ID
	cac.mutation.done = true
	return _node, nil
}

func (cac *CommentAttachmentCreate) createSpec() (*CommentAttachment, *sqlgraph.CreateSpec) {
	var (
		_node = &CommentAttachment{config: cac.config}
		_spec = sqlgraph.NewCreateSpec(commentattachment.Table, sqlgraph.NewFieldSpec(commentattachment.FieldID, field.TypeInt))
	)
	if id, ok := cac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cac.mutation.URL(); ok {
		_spec.SetField(commentattachment.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := cac.mutation.Order(); ok {
		_spec.SetField(commentattachment.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := cac.mutation.GetType(); ok {
		_spec.SetField(commentattachment.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := cac.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentattachment.CommentTable,
			Columns: []string{commentattachment.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
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

// CommentAttachmentCreateBulk is the builder for creating many CommentAttachment entities in bulk.
type CommentAttachmentCreateBulk struct {
	config
	err      error
	builders []*CommentAttachmentCreate
}

// Save creates the CommentAttachment entities in the database.
func (cacb *CommentAttachmentCreateBulk) Save(ctx context.Context) ([]*CommentAttachment, error) {
	if cacb.err != nil {
		return nil, cacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cacb.builders))
	nodes := make([]*CommentAttachment, len(cacb.builders))
	mutators := make([]Mutator, len(cacb.builders))
	for i := range cacb.builders {
		func(i int, root context.Context) {
			builder := cacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentAttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, cacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cacb *CommentAttachmentCreateBulk) SaveX(ctx context.Context) []*CommentAttachment {
	v, err := cacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cacb *CommentAttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := cacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cacb *CommentAttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := cacb.Exec(ctx); err != nil {
		panic(err)
	}
}
