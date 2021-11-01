package post

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        string    `bun:"type:uuid,default:uuid_generate_v4()"`
	TopicID   uuid.UUID `bun:"type:uuid"`
	Topic     Topic     `bun:"rel:belongs-to,join:topic_id=id"`
	Name      string
	Slug      string
	AuthorID  uuid.UUID `bun:"type:uuid"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Topic is the topic representation of this posts topic
// This avoids import cycle issues, plus we don't really need everything
// from topic.Topic anyway
type Topic struct {
	ID   uuid.UUID
	Name string
	Slug string
}
