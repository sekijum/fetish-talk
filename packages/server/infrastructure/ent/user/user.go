// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAvatarUrl holds the string denoting the avatarurl field in the database.
	FieldAvatarUrl = "avatar_url"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBoards holds the string denoting the boards edge name in mutations.
	EdgeBoards = "boards"
	// EdgeThreads holds the string denoting the threads edge name in mutations.
	EdgeThreads = "threads"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeLikedBoards holds the string denoting the liked_boards edge name in mutations.
	EdgeLikedBoards = "liked_boards"
	// EdgeLikedThreads holds the string denoting the liked_threads edge name in mutations.
	EdgeLikedThreads = "liked_threads"
	// EdgeLikedComments holds the string denoting the liked_comments edge name in mutations.
	EdgeLikedComments = "liked_comments"
	// EdgeSubscribedBoards holds the string denoting the subscribed_boards edge name in mutations.
	EdgeSubscribedBoards = "subscribed_boards"
	// EdgeSubscribedThreads holds the string denoting the subscribed_threads edge name in mutations.
	EdgeSubscribedThreads = "subscribed_threads"
	// EdgeSubscribedComments holds the string denoting the subscribed_comments edge name in mutations.
	EdgeSubscribedComments = "subscribed_comments"
	// EdgeUserBoardLike holds the string denoting the user_board_like edge name in mutations.
	EdgeUserBoardLike = "user_board_like"
	// EdgeUserThreadLike holds the string denoting the user_thread_like edge name in mutations.
	EdgeUserThreadLike = "user_thread_like"
	// EdgeUserCommentLike holds the string denoting the user_comment_like edge name in mutations.
	EdgeUserCommentLike = "user_comment_like"
	// EdgeUserBoardSubscription holds the string denoting the user_board_subscription edge name in mutations.
	EdgeUserBoardSubscription = "user_board_subscription"
	// EdgeUserThreadSubscription holds the string denoting the user_thread_subscription edge name in mutations.
	EdgeUserThreadSubscription = "user_thread_subscription"
	// EdgeUserCommentSubscription holds the string denoting the user_comment_subscription edge name in mutations.
	EdgeUserCommentSubscription = "user_comment_subscription"
	// Table holds the table name of the user in the database.
	Table = "users"
	// BoardsTable is the table that holds the boards relation/edge.
	BoardsTable = "boards"
	// BoardsInverseTable is the table name for the Board entity.
	// It exists in this package in order to avoid circular dependency with the "board" package.
	BoardsInverseTable = "boards"
	// BoardsColumn is the table column denoting the boards relation/edge.
	BoardsColumn = "user_id"
	// ThreadsTable is the table that holds the threads relation/edge.
	ThreadsTable = "threads"
	// ThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadsInverseTable = "threads"
	// ThreadsColumn is the table column denoting the threads relation/edge.
	ThreadsColumn = "user_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "thread_comments"
	// CommentsInverseTable is the table name for the ThreadComment entity.
	// It exists in this package in order to avoid circular dependency with the "threadcomment" package.
	CommentsInverseTable = "thread_comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "user_id"
	// LikedBoardsTable is the table that holds the liked_boards relation/edge. The primary key declared below.
	LikedBoardsTable = "user_board_likes"
	// LikedBoardsInverseTable is the table name for the Board entity.
	// It exists in this package in order to avoid circular dependency with the "board" package.
	LikedBoardsInverseTable = "boards"
	// LikedThreadsTable is the table that holds the liked_threads relation/edge. The primary key declared below.
	LikedThreadsTable = "user_thread_likes"
	// LikedThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	LikedThreadsInverseTable = "threads"
	// LikedCommentsTable is the table that holds the liked_comments relation/edge. The primary key declared below.
	LikedCommentsTable = "user_comment_likes"
	// LikedCommentsInverseTable is the table name for the ThreadComment entity.
	// It exists in this package in order to avoid circular dependency with the "threadcomment" package.
	LikedCommentsInverseTable = "thread_comments"
	// SubscribedBoardsTable is the table that holds the subscribed_boards relation/edge. The primary key declared below.
	SubscribedBoardsTable = "user_board_subscriptions"
	// SubscribedBoardsInverseTable is the table name for the Board entity.
	// It exists in this package in order to avoid circular dependency with the "board" package.
	SubscribedBoardsInverseTable = "boards"
	// SubscribedThreadsTable is the table that holds the subscribed_threads relation/edge. The primary key declared below.
	SubscribedThreadsTable = "user_thread_subscriptions"
	// SubscribedThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	SubscribedThreadsInverseTable = "threads"
	// SubscribedCommentsTable is the table that holds the subscribed_comments relation/edge. The primary key declared below.
	SubscribedCommentsTable = "user_comment_subscriptions"
	// SubscribedCommentsInverseTable is the table name for the ThreadComment entity.
	// It exists in this package in order to avoid circular dependency with the "threadcomment" package.
	SubscribedCommentsInverseTable = "thread_comments"
	// UserBoardLikeTable is the table that holds the user_board_like relation/edge.
	UserBoardLikeTable = "user_board_likes"
	// UserBoardLikeInverseTable is the table name for the UserBoardLike entity.
	// It exists in this package in order to avoid circular dependency with the "userboardlike" package.
	UserBoardLikeInverseTable = "user_board_likes"
	// UserBoardLikeColumn is the table column denoting the user_board_like relation/edge.
	UserBoardLikeColumn = "user_id"
	// UserThreadLikeTable is the table that holds the user_thread_like relation/edge.
	UserThreadLikeTable = "user_thread_likes"
	// UserThreadLikeInverseTable is the table name for the UserThreadLike entity.
	// It exists in this package in order to avoid circular dependency with the "userthreadlike" package.
	UserThreadLikeInverseTable = "user_thread_likes"
	// UserThreadLikeColumn is the table column denoting the user_thread_like relation/edge.
	UserThreadLikeColumn = "user_id"
	// UserCommentLikeTable is the table that holds the user_comment_like relation/edge.
	UserCommentLikeTable = "user_comment_likes"
	// UserCommentLikeInverseTable is the table name for the UserCommentLike entity.
	// It exists in this package in order to avoid circular dependency with the "usercommentlike" package.
	UserCommentLikeInverseTable = "user_comment_likes"
	// UserCommentLikeColumn is the table column denoting the user_comment_like relation/edge.
	UserCommentLikeColumn = "user_id"
	// UserBoardSubscriptionTable is the table that holds the user_board_subscription relation/edge.
	UserBoardSubscriptionTable = "user_board_subscriptions"
	// UserBoardSubscriptionInverseTable is the table name for the UserBoardSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "userboardsubscription" package.
	UserBoardSubscriptionInverseTable = "user_board_subscriptions"
	// UserBoardSubscriptionColumn is the table column denoting the user_board_subscription relation/edge.
	UserBoardSubscriptionColumn = "user_id"
	// UserThreadSubscriptionTable is the table that holds the user_thread_subscription relation/edge.
	UserThreadSubscriptionTable = "user_thread_subscriptions"
	// UserThreadSubscriptionInverseTable is the table name for the UserThreadSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "userthreadsubscription" package.
	UserThreadSubscriptionInverseTable = "user_thread_subscriptions"
	// UserThreadSubscriptionColumn is the table column denoting the user_thread_subscription relation/edge.
	UserThreadSubscriptionColumn = "user_id"
	// UserCommentSubscriptionTable is the table that holds the user_comment_subscription relation/edge.
	UserCommentSubscriptionTable = "user_comment_subscriptions"
	// UserCommentSubscriptionInverseTable is the table name for the UserCommentSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "usercommentsubscription" package.
	UserCommentSubscriptionInverseTable = "user_comment_subscriptions"
	// UserCommentSubscriptionColumn is the table column denoting the user_comment_subscription relation/edge.
	UserCommentSubscriptionColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldAvatarUrl,
	FieldStatus,
	FieldRole,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// LikedBoardsPrimaryKey and LikedBoardsColumn2 are the table columns denoting the
	// primary key for the liked_boards relation (M2M).
	LikedBoardsPrimaryKey = []string{"userId", "boardId"}
	// LikedThreadsPrimaryKey and LikedThreadsColumn2 are the table columns denoting the
	// primary key for the liked_threads relation (M2M).
	LikedThreadsPrimaryKey = []string{"userId", "threadId"}
	// LikedCommentsPrimaryKey and LikedCommentsColumn2 are the table columns denoting the
	// primary key for the liked_comments relation (M2M).
	LikedCommentsPrimaryKey = []string{"userId", "commentId"}
	// SubscribedBoardsPrimaryKey and SubscribedBoardsColumn2 are the table columns denoting the
	// primary key for the subscribed_boards relation (M2M).
	SubscribedBoardsPrimaryKey = []string{"userId", "boardId"}
	// SubscribedThreadsPrimaryKey and SubscribedThreadsColumn2 are the table columns denoting the
	// primary key for the subscribed_threads relation (M2M).
	SubscribedThreadsPrimaryKey = []string{"userId", "threadId"}
	// SubscribedCommentsPrimaryKey and SubscribedCommentsColumn2 are the table columns denoting the
	// primary key for the subscribed_comments relation (M2M).
	SubscribedCommentsPrimaryKey = []string{"userId", "commentId"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole int
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByAvatarUrl orders the results by the avatarUrl field.
func ByAvatarUrl(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarUrl, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBoardsCount orders the results by boards count.
func ByBoardsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBoardsStep(), opts...)
	}
}

// ByBoards orders the results by boards terms.
func ByBoards(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBoardsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByThreadsCount orders the results by threads count.
func ByThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadsStep(), opts...)
	}
}

// ByThreads orders the results by threads terms.
func ByThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByLikedBoardsCount orders the results by liked_boards count.
func ByLikedBoardsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikedBoardsStep(), opts...)
	}
}

// ByLikedBoards orders the results by liked_boards terms.
func ByLikedBoards(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikedBoardsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikedThreadsCount orders the results by liked_threads count.
func ByLikedThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikedThreadsStep(), opts...)
	}
}

// ByLikedThreads orders the results by liked_threads terms.
func ByLikedThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikedThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikedCommentsCount orders the results by liked_comments count.
func ByLikedCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikedCommentsStep(), opts...)
	}
}

// ByLikedComments orders the results by liked_comments terms.
func ByLikedComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikedCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubscribedBoardsCount orders the results by subscribed_boards count.
func BySubscribedBoardsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscribedBoardsStep(), opts...)
	}
}

// BySubscribedBoards orders the results by subscribed_boards terms.
func BySubscribedBoards(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscribedBoardsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubscribedThreadsCount orders the results by subscribed_threads count.
func BySubscribedThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscribedThreadsStep(), opts...)
	}
}

// BySubscribedThreads orders the results by subscribed_threads terms.
func BySubscribedThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscribedThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubscribedCommentsCount orders the results by subscribed_comments count.
func BySubscribedCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscribedCommentsStep(), opts...)
	}
}

// BySubscribedComments orders the results by subscribed_comments terms.
func BySubscribedComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscribedCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserBoardLikeCount orders the results by user_board_like count.
func ByUserBoardLikeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserBoardLikeStep(), opts...)
	}
}

// ByUserBoardLike orders the results by user_board_like terms.
func ByUserBoardLike(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserBoardLikeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserThreadLikeCount orders the results by user_thread_like count.
func ByUserThreadLikeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserThreadLikeStep(), opts...)
	}
}

// ByUserThreadLike orders the results by user_thread_like terms.
func ByUserThreadLike(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserThreadLikeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCommentLikeCount orders the results by user_comment_like count.
func ByUserCommentLikeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserCommentLikeStep(), opts...)
	}
}

// ByUserCommentLike orders the results by user_comment_like terms.
func ByUserCommentLike(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserCommentLikeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserBoardSubscriptionCount orders the results by user_board_subscription count.
func ByUserBoardSubscriptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserBoardSubscriptionStep(), opts...)
	}
}

// ByUserBoardSubscription orders the results by user_board_subscription terms.
func ByUserBoardSubscription(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserBoardSubscriptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserThreadSubscriptionCount orders the results by user_thread_subscription count.
func ByUserThreadSubscriptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserThreadSubscriptionStep(), opts...)
	}
}

// ByUserThreadSubscription orders the results by user_thread_subscription terms.
func ByUserThreadSubscription(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserThreadSubscriptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCommentSubscriptionCount orders the results by user_comment_subscription count.
func ByUserCommentSubscriptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserCommentSubscriptionStep(), opts...)
	}
}

// ByUserCommentSubscription orders the results by user_comment_subscription terms.
func ByUserCommentSubscription(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserCommentSubscriptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBoardsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BoardsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BoardsTable, BoardsColumn),
	)
}
func newThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ThreadsTable, ThreadsColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newLikedBoardsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikedBoardsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LikedBoardsTable, LikedBoardsPrimaryKey...),
	)
}
func newLikedThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikedThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LikedThreadsTable, LikedThreadsPrimaryKey...),
	)
}
func newLikedCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikedCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LikedCommentsTable, LikedCommentsPrimaryKey...),
	)
}
func newSubscribedBoardsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscribedBoardsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubscribedBoardsTable, SubscribedBoardsPrimaryKey...),
	)
}
func newSubscribedThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscribedThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubscribedThreadsTable, SubscribedThreadsPrimaryKey...),
	)
}
func newSubscribedCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscribedCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubscribedCommentsTable, SubscribedCommentsPrimaryKey...),
	)
}
func newUserBoardLikeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserBoardLikeInverseTable, UserBoardLikeColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserBoardLikeTable, UserBoardLikeColumn),
	)
}
func newUserThreadLikeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserThreadLikeInverseTable, UserThreadLikeColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserThreadLikeTable, UserThreadLikeColumn),
	)
}
func newUserCommentLikeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserCommentLikeInverseTable, UserCommentLikeColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserCommentLikeTable, UserCommentLikeColumn),
	)
}
func newUserBoardSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserBoardSubscriptionInverseTable, UserBoardSubscriptionColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserBoardSubscriptionTable, UserBoardSubscriptionColumn),
	)
}
func newUserThreadSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserThreadSubscriptionInverseTable, UserThreadSubscriptionColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserThreadSubscriptionTable, UserThreadSubscriptionColumn),
	)
}
func newUserCommentSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserCommentSubscriptionInverseTable, UserCommentSubscriptionColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserCommentSubscriptionTable, UserCommentSubscriptionColumn),
	)
}
