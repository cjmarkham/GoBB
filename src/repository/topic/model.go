package topic

import (
	"github.com/cjmarkham/GoBB/src/domain/topic"
	"github.com/google/uuid"
)

type DBResult struct {
	ID uuid.UUID
}

func (res DBResult) ToTopic() topic.Topic {
	return topic.Topic{
		ID: res.ID,
	}
}
