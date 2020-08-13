package controller

import (
	"html/template"
	"net/http"
)

func NewManga(templates map[string]*template.Template) *Manga {
	return &Manga{
		htmlTemplate: templates["manga.html"],
	}
}

type Manga struct {
	htmlTemplate *template.Template
}

func (m *Manga) Process(w http.ResponseWriter, req *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "text/html")
	m.htmlTemplate.Execute(w, nil)
	return nil
}