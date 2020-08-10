package controller

import (
	"fmt"
	"net/http"
)

type Home struct {
}

func (h *Home) Process(w http.ResponseWriter, req *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `home page`)
	return nil
}