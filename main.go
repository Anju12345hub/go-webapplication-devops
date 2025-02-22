package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func contentsPage(w http.ResponseWriter, r *http.Request) {
	// Render the contents html page
	http.ServeFile(w, r, "static/contents.html")
}


func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contacts.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/contents", contentsPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
