// Code generated by ent, DO NOT EDIT.

package adminuser

import (
	"server/infrastructure/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldID, id))
}

// UserName applies equality check predicate on the "userName" field. It's identical to UserNameEQ.
func UserName(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldUserName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldPassword, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserNameEQ applies the EQ predicate on the "userName" field.
func UserNameEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "userName" field.
func UserNameNEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "userName" field.
func UserNameIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "userName" field.
func UserNameNotIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "userName" field.
func UserNameGT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "userName" field.
func UserNameGTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "userName" field.
func UserNameLT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "userName" field.
func UserNameLTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "userName" field.
func UserNameContains(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "userName" field.
func UserNameHasPrefix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "userName" field.
func UserNameHasSuffix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "userName" field.
func UserNameEqualFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "userName" field.
func UserNameContainsFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContainsFold(FieldUserName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContainsFold(FieldPassword, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdminUser) predicate.AdminUser {
	return predicate.AdminUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdminUser) predicate.AdminUser {
	return predicate.AdminUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AdminUser) predicate.AdminUser {
	return predicate.AdminUser(sql.NotPredicates(p))
}
