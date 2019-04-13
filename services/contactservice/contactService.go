package contactservice

import (
	"html/template"
	"net/http"
	"path"
)

//HandlerContact - handler for page contact
func HandlerContact(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/contact", "contact.html")

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
