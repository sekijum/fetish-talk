// Code generated by ent, DO NOT EDIT.

package topic

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the topic type in the database.
	Label = "topic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldForumID holds the string denoting the forum_id field in the database.
	FieldForumID = "forum_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldIsDefault holds the string denoting the is_default field in the database.
	FieldIsDefault = "is_default"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeForum holds the string denoting the forum edge name in mutations.
	EdgeForum = "forum"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeTopicLikes holds the string denoting the topic_likes edge name in mutations.
	EdgeTopicLikes = "topic_likes"
	// EdgeUserTopicNotifications holds the string denoting the user_topic_notifications edge name in mutations.
	EdgeUserTopicNotifications = "user_topic_notifications"
	// Table holds the table name of the topic in the database.
	Table = "topics"
	// ForumTable is the table that holds the forum relation/edge.
	ForumTable = "topics"
	// ForumInverseTable is the table name for the Forum entity.
	// It exists in this package in order to avoid circular dependency with the "forum" package.
	ForumInverseTable = "forums"
	// ForumColumn is the table column denoting the forum relation/edge.
	ForumColumn = "forum_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "topics"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "topic_id"
	// TopicLikesTable is the table that holds the topic_likes relation/edge.
	TopicLikesTable = "topic_likes"
	// TopicLikesInverseTable is the table name for the TopicLike entity.
	// It exists in this package in order to avoid circular dependency with the "topiclike" package.
	TopicLikesInverseTable = "topic_likes"
	// TopicLikesColumn is the table column denoting the topic_likes relation/edge.
	TopicLikesColumn = "topic_id"
	// UserTopicNotificationsTable is the table that holds the user_topic_notifications relation/edge.
	UserTopicNotificationsTable = "user_topic_notifications"
	// UserTopicNotificationsInverseTable is the table name for the UserTopicNotification entity.
	// It exists in this package in order to avoid circular dependency with the "usertopicnotification" package.
	UserTopicNotificationsInverseTable = "user_topic_notifications"
	// UserTopicNotificationsColumn is the table column denoting the user_topic_notifications relation/edge.
	UserTopicNotificationsColumn = "topic_id"
)

// Columns holds all SQL columns for topic fields.
var Columns = []string{
	FieldID,
	FieldForumID,
	FieldUserID,
	FieldTitle,
	FieldContent,
	FieldIsDefault,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultIsDefault holds the default value on creation for the "is_default" field.
	DefaultIsDefault bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusOpen is the default value of the Status enum.
const DefaultStatus = StatusOpen

// Status values.
const (
	StatusOpen        Status = "Open"
	StatusClosed      Status = "Closed"
	StatusArchived    Status = "Archived"
	StatusDisapproved Status = "Disapproved"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOpen, StatusClosed, StatusArchived, StatusDisapproved:
		return nil
	default:
		return fmt.Errorf("topic: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Topic queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByForumID orders the results by the forum_id field.
func ByForumID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForumID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByIsDefault orders the results by the is_default field.
func ByIsDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDefault, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByForumField orders the results by forum field.
func ByForumField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newForumStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTopicLikesCount orders the results by topic_likes count.
func ByTopicLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTopicLikesStep(), opts...)
	}
}

// ByTopicLikes orders the results by topic_likes terms.
func ByTopicLikes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTopicLikesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserTopicNotificationsCount orders the results by user_topic_notifications count.
func ByUserTopicNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserTopicNotificationsStep(), opts...)
	}
}

// ByUserTopicNotifications orders the results by user_topic_notifications terms.
func ByUserTopicNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserTopicNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newForumStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ForumInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ForumTable, ForumColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newTopicLikesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TopicLikesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TopicLikesTable, TopicLikesColumn),
	)
}
func newUserTopicNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserTopicNotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserTopicNotificationsTable, UserTopicNotificationsColumn),
	)
}
