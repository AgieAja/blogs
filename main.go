package main

import (
	"blogs/config"

	aboutService "blogs/services/aboutservice"
	articleService "blogs/services/articleservice"
	contactService "blogs/services/contactservice"
	homeService "blogs/services/homeservice"
	loginService "blogs/services/loginservice"

	"log"

	"fmt"
	"net/http"
)

func init() {
	config.InitConnectDB()
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/assets/css"))))

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("views/assets/js"))))

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("views/assets/fonts"))))

	http.HandleFunc("/", homeService.HandlerIndex)
	http.HandleFunc("/about", aboutService.HandlerAbout)
	http.HandleFunc("/contact", contactService.HandlerContact)
	http.HandleFunc("/contact/add", contactService.HandlerAddMessage)
	http.HandleFunc("/article", articleService.HandlerArticle)
	http.HandleFunc("/article/add", articleService.HandlerCreateArticle)
	http.HandleFunc("/article/detail/", articleService.HandlerDetailArticle)
	http.HandleFunc("/article/delete/", articleService.HandlerDeleteArticle)
	http.HandleFunc("/article/publish/", articleService.HandlerPublishArticle)
	http.HandleFunc("/article/unpublish/", articleService.HandlerUnPublishArticle)
	http.HandleFunc("/article/prepareCreate", articleService.HandlerPrepareCreateArticle)
	http.HandleFunc("/login", loginService.HandlerLogin)

	// var address = "localhost:9001"
	port := "9002"
	fmt.Println("server jalan di port :", port)
	if errHTTP := http.ListenAndServe(":"+port, nil); errHTTP != nil {
		log.Println(errHTTP)
	}
}
