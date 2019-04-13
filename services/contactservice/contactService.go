package contactservice

import (
	"html/template"
	"net/http"
)

//HandlerContact - handler for page contact
func HandlerContact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/contact/contact.html",
	))

	err := tmpl.ExecuteTemplate(w, "contact", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
