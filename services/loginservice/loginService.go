package loginservice

import (
	"html/template"
	"net/http"
)

//HandlerLogin - handler page login
func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/login/login.html",
	))

	err := tmpl.ExecuteTemplate(w, "login", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
