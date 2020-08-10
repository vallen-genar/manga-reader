package controller

import (
	"fmt"
	"net/http"
)

type Manga struct {
}

func (m *Manga) Process(w http.ResponseWriter, req *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `manga page`)
	return nil
}