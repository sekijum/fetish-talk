// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/comment"
	"server/infrastructure/ent/commentattachment"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentAttachmentCreate is the builder for creating a CommentAttachment entity.
type CommentAttachmentCreate struct {
	config
	mutation *CommentAttachmentMutation
	hooks    []Hook
}

// SetCommentID sets the "comment_id" field.
func (cac *CommentAttachmentCreate) SetCommentID(i int) *CommentAttachmentCreate {
	cac.mutation.SetCommentID(i)
	return cac
}

// SetPath sets the "path" field.
func (cac *CommentAttachmentCreate) SetPath(s string) *CommentAttachmentCreate {
	cac.mutation.SetPath(s)
	return cac
}

// SetType sets the "type" field.
func (cac *CommentAttachmentCreate) SetType(c commentattachment.Type) *CommentAttachmentCreate {
	cac.mutation.SetType(c)
	return cac
}

// SetCreatedAt sets the "created_at" field.
func (cac *CommentAttachmentCreate) SetCreatedAt(t time.Time) *CommentAttachmentCreate {
	cac.mutation.SetCreatedAt(t)
	return cac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cac *CommentAttachmentCreate) SetNillableCreatedAt(t *time.Time) *CommentAttachmentCreate {
	if t != nil {
		cac.SetCreatedAt(*t)
	}
	return cac
}

// SetID sets the "id" field.
func (cac *CommentAttachmentCreate) SetID(i int) *CommentAttachmentCreate {
	cac.mutation.SetID(i)
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
	if _, ok := cac.mutation.CreatedAt(); !ok {
		v := commentattachment.DefaultCreatedAt()
		cac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cac *CommentAttachmentCreate) check() error {
	if _, ok := cac.mutation.CommentID(); !ok {
		return &ValidationError{Name: "comment_id", err: errors.New(`ent: missing required field "CommentAttachment.comment_id"`)}
	}
	if _, ok := cac.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "CommentAttachment.path"`)}
	}
	if _, ok := cac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "CommentAttachment.type"`)}
	}
	if v, ok := cac.mutation.GetType(); ok {
		if err := commentattachment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CommentAttachment.type": %w`, err)}
		}
	}
	if _, ok := cac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CommentAttachment.created_at"`)}
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
	if value, ok := cac.mutation.Path(); ok {
		_spec.SetField(commentattachment.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := cac.mutation.GetType(); ok {
		_spec.SetField(commentattachment.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := cac.mutation.CreatedAt(); ok {
		_spec.SetField(commentattachment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
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
		_node.CommentID = nodes[0]
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
