package api

import (
	"fmt"
	"github.com/vallen-genar/manga-reader/api/controller"
	"github.com/vallen-genar/router"
	"net/http"
)

func Init() {
	cf := &router.Config{
		Prefix:"",
	}
	r := router.New(cf)

	// populate templates
	templates := populateTemplates()

	// home
	r.Get("", "/", "/home").Controller(controller.NewHome(templates))

	// manga
	r.Get("/manga/{manga}").Controller(controller.NewManga(templates))

	// chapter
	r.Get("/manga/{manga}/chapter/{chapter}").Controller(controller.NewChapter(templates))

	port := 8000
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}