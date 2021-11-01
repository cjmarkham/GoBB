package post

import (
	"github.com/cjmarkham/GoBB/src/domain/post"
	"github.com/google/uuid"
)

type DBResult struct {
	ID uuid.UUID
}

func (res DBResult) ToPost() post.Post {
	return post.Post{
		ID: res.ID,
	}
}
