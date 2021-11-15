package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("Website/homepage.html"))

	tmpl.Execute(w, "data goes here")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)

	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
