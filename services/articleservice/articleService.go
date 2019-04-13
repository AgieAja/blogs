package articleservice

import (
	"html/template"
	"net/http"
)

//HandlerArticle - handler page article
func HandlerArticle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/article/article.html",
	))

	err := tmpl.ExecuteTemplate(w, "article", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
