package aboutservice

import (
	"html/template"
	"net/http"
)

//HandlerAbout - handler for page about
func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/about/about.html",
	))

	err := tmpl.ExecuteTemplate(w, "about", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
