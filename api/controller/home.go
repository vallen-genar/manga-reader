package controller

import (
	"html/template"
	"net/http"
)

func NewHome(templates map[string]*template.Template) *Home {
	return &Home{
		htmlTemplate: templates["home.html"],
	}
}

type Home struct {
	htmlTemplate *template.Template
}

type Data struct {
	Title string
	Body string
}

func (h *Home) Process(w http.ResponseWriter, req *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "text/html")
	d := Data{
		Title: "Home test",
		Body: "home body",
	}
	h.htmlTemplate.Execute(w, d)
	return nil
}