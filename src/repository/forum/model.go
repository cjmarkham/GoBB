package forum

import (
	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/google/uuid"
)

type DBResult struct {
	ID uuid.UUID
}

func (res DBResult) ToForum() forum.Forum {
	return forum.Forum{
		ID: res.ID,
	}
}
