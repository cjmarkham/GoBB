package forum

import "github.com/cjmarkham/GoBB/src/domain/forum"

type DBResult struct {
	ID string
}

func (fr DBResult) ToForum() forum.Forum {
	return forum.Forum{
		ID: fr.ID,
	}
}
