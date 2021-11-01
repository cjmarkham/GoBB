package helpers

type Crumb struct {
	Name string
	Link string
}

func GlobalTemplates() []string {
	return []string{
		"src/views/layout.html",
		"src/views/navigation.html",
		"src/views/header.html",
	}
}
