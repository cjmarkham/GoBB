package home

import (
	"github.com/cjmarkham/GoBB/src/controller/helpers"
	"github.com/cjmarkham/GoBB/src/domain/forum"
)

// IndexData is the template data for the home index page
type IndexData struct {
	Title string
	Forums []forum.Forum
	Crumbs []helpers.Crumb
}
