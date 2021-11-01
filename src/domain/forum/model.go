package forum

import (
	"time"

	"github.com/google/uuid"
)

type Forum struct {
	ID          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Name        string
	Slug        string
	Description string
	AuthorID    uuid.UUID `bun:"type:uuid"`
	Topics      int
	Posts       int
	ParentID    uuid.UUID `bun:"type:uuid,default:uuid_nil()"`
	Children    []Forum   `bun:"rel:has-many,join:id=parent_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
