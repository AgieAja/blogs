package main

import (
	aboutService "blogs/services/aboutservice"
	articleService "blogs/services/articleservice"
	contactService "blogs/services/contactservice"
	loginService "blogs/services/loginservice"
	"log"

	"fmt"
	"html/template"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	// var filepath = path.Join("views/home", "index.html")

	// tmpl, err := template.ParseFiles(filepath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	tmpl := template.Must(template.ParseFiles(
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/home/index.html",
	))

	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/assets/css"))))

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("views/assets/js"))))

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("views/assets/fonts"))))

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/about", aboutService.HandlerAbout)
	http.HandleFunc("/contact", contactService.HandlerContact)
	http.HandleFunc("/article", articleService.HandlerArticle)
	http.HandleFunc("/login", loginService.HandlerLogin)

	// var address = "localhost:9001"
	port := "9090"
	fmt.Println("server jalan di port :", port)
	if errHTTP := http.ListenAndServe(":"+port, nil); errHTTP != nil {
		log.Println(errHTTP)
	}
}
