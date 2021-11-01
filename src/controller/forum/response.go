package forum

import (
	"github.com/cjmarkham/GoBB/src/controller/helpers"
	"github.com/cjmarkham/GoBB/src/domain/topic"
)

// IndexData is the template data for the forum index page
type IndexData struct {
	Title string
	Topics []topic.Topic
	Crumbs []helpers.Crumb
}
