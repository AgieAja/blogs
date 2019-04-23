package articleservice

import (
	articleBpc "blogs/businessprocess/articlebpc"
	"log"
	"strconv"

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

	myData := articleBpc.GetArticleAll()
	if len(myData) == 0 {
		log.Println("data not found")
	}

	err := tmpl.ExecuteTemplate(w, "article", myData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//HandlerDetailArticle - handler page detail article
func HandlerDetailArticle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/article/detailArticle.html",
	))

	data := r.URL.Query()
	myid := data.Get("id")
	id, _ := strconv.ParseInt(myid, 10, 64)
	myData := articleBpc.GetDetailArticle(id)

	err := tmpl.ExecuteTemplate(w, "detailArticle", myData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
