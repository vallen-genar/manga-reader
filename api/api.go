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

	// home
	r.Get("", "/", "/home").Controller(&controller.Home{})

	// manga
	r.Get("/manga/{manga}").Controller(&controller.Manga{})

	// manga
	r.Get("/manga/{manga}/chapter/{chapter}").Controller(&controller.Chapter{})

	port := 8000
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}