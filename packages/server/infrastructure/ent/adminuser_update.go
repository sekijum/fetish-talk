// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/adminuser"
	"server/infrastructure/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminUserUpdate is the builder for updating AdminUser entities.
type AdminUserUpdate struct {
	config
	hooks    []Hook
	mutation *AdminUserMutation
}

// Where appends a list predicates to the AdminUserUpdate builder.
func (auu *AdminUserUpdate) Where(ps ...predicate.AdminUser) *AdminUserUpdate {
	auu.mutation.Where(ps...)
	return auu
}

// SetName sets the "name" field.
func (auu *AdminUserUpdate) SetName(s string) *AdminUserUpdate {
	auu.mutation.SetName(s)
	return auu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableName(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetName(*s)
	}
	return auu
}

// SetEmail sets the "email" field.
func (auu *AdminUserUpdate) SetEmail(s string) *AdminUserUpdate {
	auu.mutation.SetEmail(s)
	return auu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableEmail(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetEmail(*s)
	}
	return auu
}

// SetPassword sets the "password" field.
func (auu *AdminUserUpdate) SetPassword(s string) *AdminUserUpdate {
	auu.mutation.SetPassword(s)
	return auu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillablePassword(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetPassword(*s)
	}
	return auu
}

// SetCreatedAt sets the "createdAt" field.
func (auu *AdminUserUpdate) SetCreatedAt(t time.Time) *AdminUserUpdate {
	auu.mutation.SetCreatedAt(t)
	return auu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableCreatedAt(t *time.Time) *AdminUserUpdate {
	if t != nil {
		auu.SetCreatedAt(*t)
	}
	return auu
}

// SetUpdatedAt sets the "updatedAt" field.
func (auu *AdminUserUpdate) SetUpdatedAt(t time.Time) *AdminUserUpdate {
	auu.mutation.SetUpdatedAt(t)
	return auu
}

// Mutation returns the AdminUserMutation object of the builder.
func (auu *AdminUserUpdate) Mutation() *AdminUserMutation {
	return auu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auu *AdminUserUpdate) Save(ctx context.Context) (int, error) {
	auu.defaults()
	return withHooks(ctx, auu.sqlSave, auu.mutation, auu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auu *AdminUserUpdate) SaveX(ctx context.Context) int {
	affected, err := auu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auu *AdminUserUpdate) Exec(ctx context.Context) error {
	_, err := auu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auu *AdminUserUpdate) ExecX(ctx context.Context) {
	if err := auu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auu *AdminUserUpdate) defaults() {
	if _, ok := auu.mutation.UpdatedAt(); !ok {
		v := adminuser.UpdateDefaultUpdatedAt()
		auu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auu *AdminUserUpdate) check() error {
	if v, ok := auu.mutation.Name(); ok {
		if err := adminuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AdminUser.name": %w`, err)}
		}
	}
	if v, ok := auu.mutation.Email(); ok {
		if err := adminuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "AdminUser.email": %w`, err)}
		}
	}
	return nil
}

func (auu *AdminUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := auu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(adminuser.Table, adminuser.Columns, sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeInt))
	if ps := auu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auu.mutation.Name(); ok {
		_spec.SetField(adminuser.FieldName, field.TypeString, value)
	}
	if value, ok := auu.mutation.Email(); ok {
		_spec.SetField(adminuser.FieldEmail, field.TypeString, value)
	}
	if value, ok := auu.mutation.Password(); ok {
		_spec.SetField(adminuser.FieldPassword, field.TypeString, value)
	}
	if value, ok := auu.mutation.CreatedAt(); ok {
		_spec.SetField(adminuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auu.mutation.UpdatedAt(); ok {
		_spec.SetField(adminuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	auu.mutation.done = true
	return n, nil
}

// AdminUserUpdateOne is the builder for updating a single AdminUser entity.
type AdminUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminUserMutation
}

// SetName sets the "name" field.
func (auuo *AdminUserUpdateOne) SetName(s string) *AdminUserUpdateOne {
	auuo.mutation.SetName(s)
	return auuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableName(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetName(*s)
	}
	return auuo
}

// SetEmail sets the "email" field.
func (auuo *AdminUserUpdateOne) SetEmail(s string) *AdminUserUpdateOne {
	auuo.mutation.SetEmail(s)
	return auuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableEmail(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetEmail(*s)
	}
	return auuo
}

// SetPassword sets the "password" field.
func (auuo *AdminUserUpdateOne) SetPassword(s string) *AdminUserUpdateOne {
	auuo.mutation.SetPassword(s)
	return auuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillablePassword(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetPassword(*s)
	}
	return auuo
}

// SetCreatedAt sets the "createdAt" field.
func (auuo *AdminUserUpdateOne) SetCreatedAt(t time.Time) *AdminUserUpdateOne {
	auuo.mutation.SetCreatedAt(t)
	return auuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableCreatedAt(t *time.Time) *AdminUserUpdateOne {
	if t != nil {
		auuo.SetCreatedAt(*t)
	}
	return auuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (auuo *AdminUserUpdateOne) SetUpdatedAt(t time.Time) *AdminUserUpdateOne {
	auuo.mutation.SetUpdatedAt(t)
	return auuo
}

// Mutation returns the AdminUserMutation object of the builder.
func (auuo *AdminUserUpdateOne) Mutation() *AdminUserMutation {
	return auuo.mutation
}

// Where appends a list predicates to the AdminUserUpdate builder.
func (auuo *AdminUserUpdateOne) Where(ps ...predicate.AdminUser) *AdminUserUpdateOne {
	auuo.mutation.Where(ps...)
	return auuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auuo *AdminUserUpdateOne) Select(field string, fields ...string) *AdminUserUpdateOne {
	auuo.fields = append([]string{field}, fields...)
	return auuo
}

// Save executes the query and returns the updated AdminUser entity.
func (auuo *AdminUserUpdateOne) Save(ctx context.Context) (*AdminUser, error) {
	auuo.defaults()
	return withHooks(ctx, auuo.sqlSave, auuo.mutation, auuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auuo *AdminUserUpdateOne) SaveX(ctx context.Context) *AdminUser {
	node, err := auuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auuo *AdminUserUpdateOne) Exec(ctx context.Context) error {
	_, err := auuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auuo *AdminUserUpdateOne) ExecX(ctx context.Context) {
	if err := auuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auuo *AdminUserUpdateOne) defaults() {
	if _, ok := auuo.mutation.UpdatedAt(); !ok {
		v := adminuser.UpdateDefaultUpdatedAt()
		auuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auuo *AdminUserUpdateOne) check() error {
	if v, ok := auuo.mutation.Name(); ok {
		if err := adminuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AdminUser.name": %w`, err)}
		}
	}
	if v, ok := auuo.mutation.Email(); ok {
		if err := adminuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "AdminUser.email": %w`, err)}
		}
	}
	return nil
}

func (auuo *AdminUserUpdateOne) sqlSave(ctx context.Context) (_node *AdminUser, err error) {
	if err := auuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(adminuser.Table, adminuser.Columns, sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeInt))
	id, ok := auuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdminUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminuser.FieldID)
		for _, f := range fields {
			if !adminuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adminuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auuo.mutation.Name(); ok {
		_spec.SetField(adminuser.FieldName, field.TypeString, value)
	}
	if value, ok := auuo.mutation.Email(); ok {
		_spec.SetField(adminuser.FieldEmail, field.TypeString, value)
	}
	if value, ok := auuo.mutation.Password(); ok {
		_spec.SetField(adminuser.FieldPassword, field.TypeString, value)
	}
	if value, ok := auuo.mutation.CreatedAt(); ok {
		_spec.SetField(adminuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auuo.mutation.UpdatedAt(); ok {
		_spec.SetField(adminuser.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &AdminUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auuo.mutation.done = true
	return _node, nil
}