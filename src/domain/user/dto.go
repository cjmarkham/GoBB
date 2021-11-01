package user

import (
	"time"

	"github.com/google/uuid"
)

// DTO is a representation of a user
type DTO struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
