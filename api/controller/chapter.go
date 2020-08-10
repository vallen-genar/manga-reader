package controller

import (
	"fmt"
	"net/http"
)

type Chapter struct {
}

func (c *Chapter) Process(w http.ResponseWriter, req *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `chapter page`)
	return nil
}