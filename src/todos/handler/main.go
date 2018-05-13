package handler

import (
	"net/http"

	"todos/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.HTMLRender("index.tmpl", w, nil)
}
