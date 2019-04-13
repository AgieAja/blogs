package loginservice

import (
	"html/template"
	"net/http"
	"path"
)

//HandlerLogin - handler page login
func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/login", "login.html")

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
