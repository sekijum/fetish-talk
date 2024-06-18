// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Status holds the value of the "status" field.
	Status user.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Forums holds the value of the forums edge.
	Forums []*Forum `json:"forums,omitempty"`
	// Topics holds the value of the topics edge.
	Topics []*Topic `json:"topics,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// UserTopicNotifications holds the value of the user_topic_notifications edge.
	UserTopicNotifications []*UserTopicNotification `json:"user_topic_notifications,omitempty"`
	// UserCommentNotifications holds the value of the user_comment_notifications edge.
	UserCommentNotifications []*UserCommentNotification `json:"user_comment_notifications,omitempty"`
	// ForumLikes holds the value of the forum_likes edge.
	ForumLikes []*ForumLike `json:"forum_likes,omitempty"`
	// TopicLikes holds the value of the topic_likes edge.
	TopicLikes []*TopicLike `json:"topic_likes,omitempty"`
	// CommentLikes holds the value of the comment_likes edge.
	CommentLikes []*CommentLike `json:"comment_likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// ForumsOrErr returns the Forums value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ForumsOrErr() ([]*Forum, error) {
	if e.loadedTypes[0] {
		return e.Forums, nil
	}
	return nil, &NotLoadedError{edge: "forums"}
}

// TopicsOrErr returns the Topics value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TopicsOrErr() ([]*Topic, error) {
	if e.loadedTypes[1] {
		return e.Topics, nil
	}
	return nil, &NotLoadedError{edge: "topics"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// UserTopicNotificationsOrErr returns the UserTopicNotifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserTopicNotificationsOrErr() ([]*UserTopicNotification, error) {
	if e.loadedTypes[3] {
		return e.UserTopicNotifications, nil
	}
	return nil, &NotLoadedError{edge: "user_topic_notifications"}
}

// UserCommentNotificationsOrErr returns the UserCommentNotifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserCommentNotificationsOrErr() ([]*UserCommentNotification, error) {
	if e.loadedTypes[4] {
		return e.UserCommentNotifications, nil
	}
	return nil, &NotLoadedError{edge: "user_comment_notifications"}
}

// ForumLikesOrErr returns the ForumLikes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ForumLikesOrErr() ([]*ForumLike, error) {
	if e.loadedTypes[5] {
		return e.ForumLikes, nil
	}
	return nil, &NotLoadedError{edge: "forum_likes"}
}

// TopicLikesOrErr returns the TopicLikes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TopicLikesOrErr() ([]*TopicLike, error) {
	if e.loadedTypes[6] {
		return e.TopicLikes, nil
	}
	return nil, &NotLoadedError{edge: "topic_likes"}
}

// CommentLikesOrErr returns the CommentLikes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentLikesOrErr() ([]*CommentLike, error) {
	if e.loadedTypes[7] {
		return e.CommentLikes, nil
	}
	return nil, &NotLoadedError{edge: "comment_likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldEmail, user.FieldPassword, user.FieldDisplayName, user.FieldAvatar, user.FieldStatus:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryForums queries the "forums" edge of the User entity.
func (u *User) QueryForums() *ForumQuery {
	return NewUserClient(u.config).QueryForums(u)
}

// QueryTopics queries the "topics" edge of the User entity.
func (u *User) QueryTopics() *TopicQuery {
	return NewUserClient(u.config).QueryTopics(u)
}

// QueryComments queries the "comments" edge of the User entity.
func (u *User) QueryComments() *CommentQuery {
	return NewUserClient(u.config).QueryComments(u)
}

// QueryUserTopicNotifications queries the "user_topic_notifications" edge of the User entity.
func (u *User) QueryUserTopicNotifications() *UserTopicNotificationQuery {
	return NewUserClient(u.config).QueryUserTopicNotifications(u)
}

// QueryUserCommentNotifications queries the "user_comment_notifications" edge of the User entity.
func (u *User) QueryUserCommentNotifications() *UserCommentNotificationQuery {
	return NewUserClient(u.config).QueryUserCommentNotifications(u)
}

// QueryForumLikes queries the "forum_likes" edge of the User entity.
func (u *User) QueryForumLikes() *ForumLikeQuery {
	return NewUserClient(u.config).QueryForumLikes(u)
}

// QueryTopicLikes queries the "topic_likes" edge of the User entity.
func (u *User) QueryTopicLikes() *TopicLikeQuery {
	return NewUserClient(u.config).QueryTopicLikes(u)
}

// QueryCommentLikes queries the "comment_likes" edge of the User entity.
func (u *User) QueryCommentLikes() *CommentLikeQuery {
	return NewUserClient(u.config).QueryCommentLikes(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
