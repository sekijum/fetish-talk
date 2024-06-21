// Code generated by ent, DO NOT EDIT.

package usercommentlike

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usercommentlike type in the database.
	Label = "user_comment_like"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldCommentId holds the string denoting the commentid field in the database.
	FieldCommentId = "comment_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldLikedAt holds the string denoting the likedat field in the database.
	FieldLikedAt = "liked_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// CommentFieldID holds the string denoting the ID field of the Comment.
	CommentFieldID = "id"
	// Table holds the table name of the usercommentlike in the database.
	Table = "user_comment_likes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_comment_likes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CommentTable is the table that holds the comment relation/edge.
	CommentTable = "user_comment_likes"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
	// CommentColumn is the table column denoting the comment relation/edge.
	CommentColumn = "comment_id"
)

// Columns holds all SQL columns for usercommentlike fields.
var Columns = []string{
	FieldUserId,
	FieldCommentId,
	FieldType,
	FieldLikedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLikedAt holds the default value on creation for the "likedAt" field.
	DefaultLikedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeLike    Type = "like"
	TypeDislike Type = "dislike"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeLike, TypeDislike:
		return nil
	default:
		return fmt.Errorf("usercommentlike: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the UserCommentLike queries.
type OrderOption func(*sql.Selector)

// ByUserId orders the results by the userId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByCommentId orders the results by the commentId field.
func ByCommentId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentId, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByLikedAt orders the results by the likedAt field.
func ByLikedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentField orders the results by comment field.
func ByCommentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, UserColumn),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
func newCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, CommentColumn),
		sqlgraph.To(CommentInverseTable, CommentFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CommentTable, CommentColumn),
	)
}
