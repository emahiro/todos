package handler

import (
	"net/http"

	"todos/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.RenderHTML("index.tmpl", w, nil)
}
