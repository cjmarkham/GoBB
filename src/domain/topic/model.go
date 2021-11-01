package topic

import (
	"time"

	"github.com/cjmarkham/GoBB/src/domain/user"
	"github.com/google/uuid"
)

type Topic struct {
	ID           uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	ForumID      uuid.UUID `bun:"type:uuid"`
	Forum        Forum     `bun:"rel:belongs-to,join:forum_id=id"`
	Name         string
	AuthorID     uuid.UUID `bun:"type:uuid"`
	Author       user.User `bun:"rel:belongs-to,join:author_id=id"`
	LastPosterID uuid.UUID `bun:"type:uuid,default:uuid_nil()"`
	LastPoster   user.User `bun:"rel:belongs-to,join:last_poster_id=id"`
	Posts        int
	Views        int
	Slug         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Forum is the forum representation of this topics parent forum
// This avoids import cycle issues, plus we don't really need everything
// from forum.Forum anyway
type Forum struct {
	ID   uuid.UUID
	Name string
	Slug string
}
