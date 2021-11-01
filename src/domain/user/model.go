package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
