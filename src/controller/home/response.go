package home

import "github.com/cjmarkham/GoBB/src/domain/forum"

type HomeIndexData struct {
	Title string
	Forums []forum.Forum
}
