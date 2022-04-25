package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
