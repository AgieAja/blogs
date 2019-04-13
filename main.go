package main

import (
	aboutService "blogs/services/aboutservice"
	articleService "blogs/services/articleservice"
	contactService "blogs/services/contactservice"
	loginService "blogs/services/loginservice"

	"fmt"
	"html/template"
	"net/http"
	"path"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/home", "index.html")

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

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/assets/css"))))

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("views/assets/js"))))

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("views/assets/fonts"))))

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/about", aboutService.HandlerAbout)
	http.HandleFunc("/contact", contactService.HandlerContact)
	http.HandleFunc("/article", articleService.HandlerArticle)
	http.HandleFunc("/login", loginService.HandlerLogin)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
