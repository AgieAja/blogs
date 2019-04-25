package contactservice

import (
	contactBpc "blogs/businessprocess/contactbpc"
	"log"

	contactModel "blogs/models/contactmodel"

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

//HandlerAddMessage - handler for add messages data
func HandlerAddMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var email = r.FormValue("email")
		var nama = r.FormValue("nama")
		var pesan = r.Form.Get("pesan")
		myData := contactModel.AddMessages{
			FromName:  nama,
			FromEmail: email,
			Message:   pesan,
			CreatedBy: 1,
			UpdatedBy: 1,
		}

		res := contactBpc.AddMessages(myData)
		if !res {
			http.Error(w, "Failed Saved", http.StatusBadRequest)
			return
		}
		log.Println("Data has been saved")

		tmpl := template.Must(template.ParseFiles(
			"views/templates/header.html",
			"views/templates/navbar.html",
			"views/contact/contact.html",
		))

		err := tmpl.ExecuteTemplate(w, "contact", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}
