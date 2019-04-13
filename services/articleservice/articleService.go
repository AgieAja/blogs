package articleservice

import (
	"html/template"
	"net/http"
	"path"
)

//HandlerArticle - handler page article
func HandlerArticle(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/article", "article.html")

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
