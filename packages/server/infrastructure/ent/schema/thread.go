package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Thread struct {
	ent.Schema
}

func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("boardId").StorageKey("board_id"),
		field.Int("userId").StorageKey("user_id"),
		field.String("title").Unique().MaxLen(50),
		field.String("description").Optional().MaxLen(255),
		field.String("thumbnailUrl").Optional().StorageKey("thumbnail_url"),
		field.String("ip_address").MaxLen(64).Comment("スレッド作成者のIPアドレス"),
		field.Int("status").Default(0).Comment("0: Open, 1: Pending, 2: Archived"),
		field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
	}
}

func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("board", Board.Type).Ref("threads").Unique().Field("boardId").Required(),
		edge.From("owner", User.Type).Ref("threads").Unique().Field("userId").Required(),
		edge.To("comments", ThreadComment.Type),
		edge.To("tags", Tag.Type).Through("thread_tags", ThreadTag.Type),
		edge.From("liked_users", User.Type).Ref("liked_threads").Through("user_thread_like", UserThreadLike.Type),
		edge.From("subscribed_users", User.Type).Ref("subscribed_threads").Through("user_thread_subscription", UserThreadSubscription.Type),
	}
}

func (Thread) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique(),
	}
}
