package articleservice

import (
	articleBpc "blogs/businessprocess/articlebpc"
	articleModel "blogs/models/articlemodel"
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

	err := tmpl.ExecuteTemplate(w, "article", myData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//HandlerPrepareCreateArticle - handler for show form create article
func HandlerPrepareCreateArticle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/article/add_article.html",
	))

	err := tmpl.ExecuteTemplate(w, "add_article", nil)
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

//HandlerCreateArticle - handler create new article
func HandlerCreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var title = r.FormValue("title")
		var shortdesc = r.FormValue("shortdesc")
		var desc = r.Form.Get("desc")
		myData := articleModel.AddArticle{
			Title:     title,
			ShortDesc: shortdesc,
			Desc:      desc,
			CreatedBy: 1,
			UpdatedBy: 1,
		}

		res := articleBpc.AddArticle(myData)
		if !res {
			http.Error(w, "Failed Saved", http.StatusBadRequest)
			return
		}

		log.Println("Data has been saved")

		tmpl := template.Must(template.ParseFiles(
			"views/templates/header.html",
			"views/templates/navbar.html",
			"views/article/article.html",
		))

		myList := articleBpc.GetArticleAll()
		err := tmpl.ExecuteTemplate(w, "article", myList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}

//HandlerDeleteArticle - handler delete article
func HandlerDeleteArticle(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	myid := data.Get("id")
	id, _ := strconv.ParseInt(myid, 10, 64)

	res := articleBpc.DeleteArticle(id)

	if !res {
		http.Error(w, "Failed Deleted", http.StatusBadRequest)
		return
	}

	log.Println("Data has been deleted")

	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/article/article.html",
	))

	myList := articleBpc.GetArticleAll()
	err := tmpl.ExecuteTemplate(w, "article", myList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//HandlerPublishArticle - handler publish article
func HandlerPublishArticle(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	myid := data.Get("id")
	id, _ := strconv.ParseInt(myid, 10, 64)

	res := articleBpc.PublishArticle(id)

	if !res {
		http.Error(w, "Failed Publish", http.StatusBadRequest)
		return
	}

	log.Println("Data has been publish")

	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/article/article.html",
	))

	myList := articleBpc.GetArticleAll()
	err := tmpl.ExecuteTemplate(w, "article", myList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//HandlerUnPublishArticle - handler unpublish article
func HandlerUnPublishArticle(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	myid := data.Get("id")
	id, _ := strconv.ParseInt(myid, 10, 64)

	res := articleBpc.UnPublishArticle(id)

	if !res {
		http.Error(w, "Failed UnPublish", http.StatusBadRequest)
		return
	}

	log.Println("Data has been unpublish")

	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/article/article.html",
	))

	myList := articleBpc.GetArticleAll()
	err := tmpl.ExecuteTemplate(w, "article", myList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
