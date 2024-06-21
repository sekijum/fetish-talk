// Code generated by ent, DO NOT EDIT.

package user

import (
	"server/infrastructure/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// UserName applies equality check predicate on the "userName" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// DisplayName applies equality check predicate on the "displayName" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// AvatarUrl applies equality check predicate on the "avatarUrl" field. It's identical to AvatarUrlEQ.
func AvatarUrl(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarUrl, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserNameEQ applies the EQ predicate on the "userName" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "userName" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "userName" field.
func UserNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "userName" field.
func UserNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "userName" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "userName" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "userName" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "userName" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "userName" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "userName" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "userName" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "userName" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "userName" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// DisplayNameEQ applies the EQ predicate on the "displayName" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "displayName" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "displayName" field.
func DisplayNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "displayName" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "displayName" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "displayName" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "displayName" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "displayName" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "displayName" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "displayName" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "displayName" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "displayName" field.
func DisplayNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "displayName" field.
func DisplayNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "displayName" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "displayName" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisplayName, v))
}

// AvatarUrlEQ applies the EQ predicate on the "avatarUrl" field.
func AvatarUrlEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarUrl, v))
}

// AvatarUrlNEQ applies the NEQ predicate on the "avatarUrl" field.
func AvatarUrlNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatarUrl, v))
}

// AvatarUrlIn applies the In predicate on the "avatarUrl" field.
func AvatarUrlIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatarUrl, vs...))
}

// AvatarUrlNotIn applies the NotIn predicate on the "avatarUrl" field.
func AvatarUrlNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatarUrl, vs...))
}

// AvatarUrlGT applies the GT predicate on the "avatarUrl" field.
func AvatarUrlGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatarUrl, v))
}

// AvatarUrlGTE applies the GTE predicate on the "avatarUrl" field.
func AvatarUrlGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatarUrl, v))
}

// AvatarUrlLT applies the LT predicate on the "avatarUrl" field.
func AvatarUrlLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatarUrl, v))
}

// AvatarUrlLTE applies the LTE predicate on the "avatarUrl" field.
func AvatarUrlLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatarUrl, v))
}

// AvatarUrlContains applies the Contains predicate on the "avatarUrl" field.
func AvatarUrlContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatarUrl, v))
}

// AvatarUrlHasPrefix applies the HasPrefix predicate on the "avatarUrl" field.
func AvatarUrlHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatarUrl, v))
}

// AvatarUrlHasSuffix applies the HasSuffix predicate on the "avatarUrl" field.
func AvatarUrlHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatarUrl, v))
}

// AvatarUrlIsNil applies the IsNil predicate on the "avatarUrl" field.
func AvatarUrlIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatarUrl))
}

// AvatarUrlNotNil applies the NotNil predicate on the "avatarUrl" field.
func AvatarUrlNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatarUrl))
}

// AvatarUrlEqualFold applies the EqualFold predicate on the "avatarUrl" field.
func AvatarUrlEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatarUrl, v))
}

// AvatarUrlContainsFold applies the ContainsFold predicate on the "avatarUrl" field.
func AvatarUrlContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatarUrl, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasForums applies the HasEdge predicate on the "forums" edge.
func HasForums() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ForumsTable, ForumsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasForumsWith applies the HasEdge predicate on the "forums" edge with a given conditions (other predicates).
func HasForumsWith(preds ...predicate.Forum) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newForumsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTopics applies the HasEdge predicate on the "topics" edge.
func HasTopics() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TopicsTable, TopicsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTopicsWith applies the HasEdge predicate on the "topics" edge with a given conditions (other predicates).
func HasTopicsWith(preds ...predicate.Topic) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTopicsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikedForums applies the HasEdge predicate on the "liked_forums" edge.
func HasLikedForums() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LikedForumsTable, LikedForumsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikedForumsWith applies the HasEdge predicate on the "liked_forums" edge with a given conditions (other predicates).
func HasLikedForumsWith(preds ...predicate.Forum) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newLikedForumsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikedTopics applies the HasEdge predicate on the "liked_topics" edge.
func HasLikedTopics() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LikedTopicsTable, LikedTopicsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikedTopicsWith applies the HasEdge predicate on the "liked_topics" edge with a given conditions (other predicates).
func HasLikedTopicsWith(preds ...predicate.Topic) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newLikedTopicsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikedComments applies the HasEdge predicate on the "liked_comments" edge.
func HasLikedComments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LikedCommentsTable, LikedCommentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikedCommentsWith applies the HasEdge predicate on the "liked_comments" edge with a given conditions (other predicates).
func HasLikedCommentsWith(preds ...predicate.Comment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newLikedCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscribedForums applies the HasEdge predicate on the "subscribed_forums" edge.
func HasSubscribedForums() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubscribedForumsTable, SubscribedForumsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscribedForumsWith applies the HasEdge predicate on the "subscribed_forums" edge with a given conditions (other predicates).
func HasSubscribedForumsWith(preds ...predicate.Forum) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSubscribedForumsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscribedTopics applies the HasEdge predicate on the "subscribed_topics" edge.
func HasSubscribedTopics() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubscribedTopicsTable, SubscribedTopicsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscribedTopicsWith applies the HasEdge predicate on the "subscribed_topics" edge with a given conditions (other predicates).
func HasSubscribedTopicsWith(preds ...predicate.Topic) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSubscribedTopicsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscribedComments applies the HasEdge predicate on the "subscribed_comments" edge.
func HasSubscribedComments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubscribedCommentsTable, SubscribedCommentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscribedCommentsWith applies the HasEdge predicate on the "subscribed_comments" edge with a given conditions (other predicates).
func HasSubscribedCommentsWith(preds ...predicate.Comment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSubscribedCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserForumLike applies the HasEdge predicate on the "user_forum_like" edge.
func HasUserForumLike() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserForumLikeTable, UserForumLikeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserForumLikeWith applies the HasEdge predicate on the "user_forum_like" edge with a given conditions (other predicates).
func HasUserForumLikeWith(preds ...predicate.UserForumLike) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserForumLikeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserTopicLike applies the HasEdge predicate on the "user_topic_like" edge.
func HasUserTopicLike() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserTopicLikeTable, UserTopicLikeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserTopicLikeWith applies the HasEdge predicate on the "user_topic_like" edge with a given conditions (other predicates).
func HasUserTopicLikeWith(preds ...predicate.UserTopicLike) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserTopicLikeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserCommentLike applies the HasEdge predicate on the "user_comment_like" edge.
func HasUserCommentLike() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserCommentLikeTable, UserCommentLikeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserCommentLikeWith applies the HasEdge predicate on the "user_comment_like" edge with a given conditions (other predicates).
func HasUserCommentLikeWith(preds ...predicate.UserCommentLike) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserCommentLikeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserForumSubscription applies the HasEdge predicate on the "user_forum_subscription" edge.
func HasUserForumSubscription() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserForumSubscriptionTable, UserForumSubscriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserForumSubscriptionWith applies the HasEdge predicate on the "user_forum_subscription" edge with a given conditions (other predicates).
func HasUserForumSubscriptionWith(preds ...predicate.UserForumSubscription) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserForumSubscriptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserTopicSubscription applies the HasEdge predicate on the "user_topic_subscription" edge.
func HasUserTopicSubscription() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserTopicSubscriptionTable, UserTopicSubscriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserTopicSubscriptionWith applies the HasEdge predicate on the "user_topic_subscription" edge with a given conditions (other predicates).
func HasUserTopicSubscriptionWith(preds ...predicate.UserTopicSubscription) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserTopicSubscriptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserCommentSubscription applies the HasEdge predicate on the "user_comment_subscription" edge.
func HasUserCommentSubscription() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserCommentSubscriptionTable, UserCommentSubscriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserCommentSubscriptionWith applies the HasEdge predicate on the "user_comment_subscription" edge with a given conditions (other predicates).
func HasUserCommentSubscriptionWith(preds ...predicate.UserCommentSubscription) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserCommentSubscriptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
