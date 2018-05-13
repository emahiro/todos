package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func HTMLRender(file string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(fmt.Sprintf("tmpl/layout/%s", file))
	if err != nil {
		return fmt.Errorf("parse html error. err: %v", err)
	}

	return t.Execute(w, data)
}
