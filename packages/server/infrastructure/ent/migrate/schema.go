// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminUsersColumns holds the columns for the "admin_users" table.
	AdminUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 254},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AdminUsersTable holds the schema information for the "admin_users" table.
	AdminUsersTable = &schema.Table{
		Name:       "admin_users",
		Columns:    AdminUsersColumns,
		PrimaryKey: []*schema.Column{AdminUsersColumns[0]},
	}
	// BoardsColumns holds the columns for the "boards" table.
	BoardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "thumbnail_url", Type: field.TypeString, Nullable: true},
		{Name: "order", Type: field.TypeInt, Default: 0},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Public", "Private", "Archived", "Deleted"}, Default: "Public"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_boards", Type: field.TypeInt, Nullable: true},
	}
	// BoardsTable holds the schema information for the "boards" table.
	BoardsTable = &schema.Table{
		Name:       "boards",
		Columns:    BoardsColumns,
		PrimaryKey: []*schema.Column{BoardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "boards_users_boards",
				Columns:    []*schema.Column{BoardsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "guest_name", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "message", Type: field.TypeString, Size: 2147483647},
		{Name: "ip_address", Type: field.TypeString, Size: 64},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Visible", "Deleted"}, Default: "Visible"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "parent_comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "thread_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_comments_replies",
				Columns:    []*schema.Column{CommentsColumns[7]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_threads_comments",
				Columns:    []*schema.Column{CommentsColumns[8]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentAttachmentsColumns holds the columns for the "comment_attachments" table.
	CommentAttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt, Default: 0},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"image", "video"}},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// CommentAttachmentsTable holds the schema information for the "comment_attachments" table.
	CommentAttachmentsTable = &schema.Table{
		Name:       "comment_attachments",
		Columns:    CommentAttachmentsColumns,
		PrimaryKey: []*schema.Column{CommentAttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_attachments_comments_comment_attachments",
				Columns:    []*schema.Column{CommentAttachmentsColumns[4]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ThreadsColumns holds the columns for the "threads" table.
	ThreadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "thumbnail_url", Type: field.TypeString, Nullable: true},
		{Name: "is_auto_generated", Type: field.TypeBool, Default: false},
		{Name: "is_notify_on_comment", Type: field.TypeBool, Default: true},
		{Name: "ip_address", Type: field.TypeString, Size: 64},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Open", "Archived", "Deleted"}, Default: "Open"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "board_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ThreadsTable holds the schema information for the "threads" table.
	ThreadsTable = &schema.Table{
		Name:       "threads",
		Columns:    ThreadsColumns,
		PrimaryKey: []*schema.Column{ThreadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "threads_boards_threads",
				Columns:    []*schema.Column{ThreadsColumns[10]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "threads_users_threads",
				Columns:    []*schema.Column{ThreadsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ThreadTagsColumns holds the columns for the "thread_tags" table.
	ThreadTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ThreadTagsTable holds the schema information for the "thread_tags" table.
	ThreadTagsTable = &schema.Table{
		Name:       "thread_tags",
		Columns:    ThreadTagsColumns,
		PrimaryKey: []*schema.Column{ThreadTagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "threadtag_name",
				Unique:  true,
				Columns: []*schema.Column{ThreadTagsColumns[1]},
			},
		},
	}
	// ThreadTaggingsColumns holds the columns for the "thread_taggings" table.
	ThreadTaggingsColumns = []*schema.Column{
		{Name: "thread_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// ThreadTaggingsTable holds the schema information for the "thread_taggings" table.
	ThreadTaggingsTable = &schema.Table{
		Name:       "thread_taggings",
		Columns:    ThreadTaggingsColumns,
		PrimaryKey: []*schema.Column{ThreadTaggingsColumns[0], ThreadTaggingsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_taggings_threads_thread",
				Columns:    []*schema.Column{ThreadTaggingsColumns[0]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "thread_taggings_thread_tags_tag",
				Columns:    []*schema.Column{ThreadTaggingsColumns[1]},
				RefColumns: []*schema.Column{ThreadTagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 254},
		{Name: "password", Type: field.TypeString},
		{Name: "display_name", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Active", "Withdrawn", "Suspended", "Inactive"}, Default: "Active"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserBoardLikesColumns holds the columns for the "user_board_likes" table.
	UserBoardLikesColumns = []*schema.Column{
		{Name: "liked_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "board_id", Type: field.TypeInt},
	}
	// UserBoardLikesTable holds the schema information for the "user_board_likes" table.
	UserBoardLikesTable = &schema.Table{
		Name:       "user_board_likes",
		Columns:    UserBoardLikesColumns,
		PrimaryKey: []*schema.Column{UserBoardLikesColumns[1], UserBoardLikesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_board_likes_users_user",
				Columns:    []*schema.Column{UserBoardLikesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_board_likes_boards_board",
				Columns:    []*schema.Column{UserBoardLikesColumns[2]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserBoardSubscriptionsColumns holds the columns for the "user_board_subscriptions" table.
	UserBoardSubscriptionsColumns = []*schema.Column{
		{Name: "is_notified", Type: field.TypeBool, Default: true},
		{Name: "is_checked", Type: field.TypeBool, Default: false},
		{Name: "subscribed_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "board_id", Type: field.TypeInt},
	}
	// UserBoardSubscriptionsTable holds the schema information for the "user_board_subscriptions" table.
	UserBoardSubscriptionsTable = &schema.Table{
		Name:       "user_board_subscriptions",
		Columns:    UserBoardSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserBoardSubscriptionsColumns[3], UserBoardSubscriptionsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_board_subscriptions_users_user",
				Columns:    []*schema.Column{UserBoardSubscriptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_board_subscriptions_boards_board",
				Columns:    []*schema.Column{UserBoardSubscriptionsColumns[4]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserCommentLikesColumns holds the columns for the "user_comment_likes" table.
	UserCommentLikesColumns = []*schema.Column{
		{Name: "liked_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// UserCommentLikesTable holds the schema information for the "user_comment_likes" table.
	UserCommentLikesTable = &schema.Table{
		Name:       "user_comment_likes",
		Columns:    UserCommentLikesColumns,
		PrimaryKey: []*schema.Column{UserCommentLikesColumns[1], UserCommentLikesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_comment_likes_users_user",
				Columns:    []*schema.Column{UserCommentLikesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_comment_likes_comments_comment",
				Columns:    []*schema.Column{UserCommentLikesColumns[2]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserCommentSubscriptionsColumns holds the columns for the "user_comment_subscriptions" table.
	UserCommentSubscriptionsColumns = []*schema.Column{
		{Name: "is_notified", Type: field.TypeBool, Default: true},
		{Name: "is_checked", Type: field.TypeBool, Default: false},
		{Name: "subscribed_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// UserCommentSubscriptionsTable holds the schema information for the "user_comment_subscriptions" table.
	UserCommentSubscriptionsTable = &schema.Table{
		Name:       "user_comment_subscriptions",
		Columns:    UserCommentSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserCommentSubscriptionsColumns[3], UserCommentSubscriptionsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_comment_subscriptions_users_user",
				Columns:    []*schema.Column{UserCommentSubscriptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_comment_subscriptions_comments_comment",
				Columns:    []*schema.Column{UserCommentSubscriptionsColumns[4]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserThreadLikesColumns holds the columns for the "user_thread_likes" table.
	UserThreadLikesColumns = []*schema.Column{
		{Name: "liked_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "thread_id", Type: field.TypeInt},
	}
	// UserThreadLikesTable holds the schema information for the "user_thread_likes" table.
	UserThreadLikesTable = &schema.Table{
		Name:       "user_thread_likes",
		Columns:    UserThreadLikesColumns,
		PrimaryKey: []*schema.Column{UserThreadLikesColumns[1], UserThreadLikesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_thread_likes_users_user",
				Columns:    []*schema.Column{UserThreadLikesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_thread_likes_threads_thread",
				Columns:    []*schema.Column{UserThreadLikesColumns[2]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserThreadSubscriptionsColumns holds the columns for the "user_thread_subscriptions" table.
	UserThreadSubscriptionsColumns = []*schema.Column{
		{Name: "is_notified", Type: field.TypeBool, Default: true},
		{Name: "is_checked", Type: field.TypeBool, Default: false},
		{Name: "subscribed_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "thread_id", Type: field.TypeInt},
	}
	// UserThreadSubscriptionsTable holds the schema information for the "user_thread_subscriptions" table.
	UserThreadSubscriptionsTable = &schema.Table{
		Name:       "user_thread_subscriptions",
		Columns:    UserThreadSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserThreadSubscriptionsColumns[3], UserThreadSubscriptionsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_thread_subscriptions_users_user",
				Columns:    []*schema.Column{UserThreadSubscriptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_thread_subscriptions_threads_thread",
				Columns:    []*schema.Column{UserThreadSubscriptionsColumns[4]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminUsersTable,
		BoardsTable,
		CommentsTable,
		CommentAttachmentsTable,
		ThreadsTable,
		ThreadTagsTable,
		ThreadTaggingsTable,
		UsersTable,
		UserBoardLikesTable,
		UserBoardSubscriptionsTable,
		UserCommentLikesTable,
		UserCommentSubscriptionsTable,
		UserThreadLikesTable,
		UserThreadSubscriptionsTable,
	}
)

func init() {
	BoardsTable.ForeignKeys[0].RefTable = UsersTable
	CommentsTable.ForeignKeys[0].RefTable = CommentsTable
	CommentsTable.ForeignKeys[1].RefTable = ThreadsTable
	CommentsTable.ForeignKeys[2].RefTable = UsersTable
	CommentAttachmentsTable.ForeignKeys[0].RefTable = CommentsTable
	ThreadsTable.ForeignKeys[0].RefTable = BoardsTable
	ThreadsTable.ForeignKeys[1].RefTable = UsersTable
	ThreadTaggingsTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadTaggingsTable.ForeignKeys[1].RefTable = ThreadTagsTable
	UserBoardLikesTable.ForeignKeys[0].RefTable = UsersTable
	UserBoardLikesTable.ForeignKeys[1].RefTable = BoardsTable
	UserBoardSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	UserBoardSubscriptionsTable.ForeignKeys[1].RefTable = BoardsTable
	UserCommentLikesTable.ForeignKeys[0].RefTable = UsersTable
	UserCommentLikesTable.ForeignKeys[1].RefTable = CommentsTable
	UserCommentSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	UserCommentSubscriptionsTable.ForeignKeys[1].RefTable = CommentsTable
	UserThreadLikesTable.ForeignKeys[0].RefTable = UsersTable
	UserThreadLikesTable.ForeignKeys[1].RefTable = ThreadsTable
	UserThreadSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	UserThreadSubscriptionsTable.ForeignKeys[1].RefTable = ThreadsTable
}
