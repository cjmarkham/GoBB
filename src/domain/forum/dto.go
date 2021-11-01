package forum

import (
	"time"

	"github.com/google/uuid"
)

// DTO is a representation of a forum
type DTO struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description string
	AuthorID    uuid.UUID
	Topics      int
	Posts       int
	ParentID    *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
