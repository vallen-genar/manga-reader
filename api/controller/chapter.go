package controller

import (
	"html/template"
	"net/http"
)

func NewChapter(templates map[string]*template.Template) *Chapter {
	return &Chapter{
		htmlTemplate: templates["chapter.html"],
	}
}

type Chapter struct {
	htmlTemplate *template.Template
}

type dd struct {
	Title string
	Manga string
}

func (c *Chapter) Process(w http.ResponseWriter, req *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "text/html")
	d := dd{
		Title: params["manga"] + "_" + params["chapter"],
		Manga: params["manga"],
	}

	c.htmlTemplate.Execute(w, d)
	return nil
}