package handler

import (
	"net/http"

	"vueSampleApp/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.HTMLRender("index.tmpl", w, nil)
}
