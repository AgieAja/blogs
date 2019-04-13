package homeservice

import (
	articleBpc "blogs/businessprocess/articlebpc"

	"html/template"
	"log"
	"net/http"
)

//HandlerIndex - handler page home
func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/home/index.html",
	))

	myList := articleBpc.GetArticlePublish()
	if len(myList) == 0 {
		log.Println("data not found")
	}

	err := tmpl.ExecuteTemplate(w, "home", myList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
