package aboutservice

import (
	"html/template"
	"net/http"
	"path"
)

//HandlerAbout - handler for page about
func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/about", "about.html")

	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
