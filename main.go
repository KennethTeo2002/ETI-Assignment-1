package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Variables
const PassengerAPIbaseURL = "http://localhost:5000/api/v1/passenger"

// HTML display functions
func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("Website/homepage.html"))

	tmpl.Execute(w, "data goes here")
}

func passengerPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("Website/homepage.html"))

	tmpl.Execute(w, "data goes here")
}

// API call functions
func passengerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/passengerLogin.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		url := PassengerAPIbaseURL
		id := r.FormValue("id")

		if id != "" {
			url = PassengerAPIbaseURL + "/" + id
		}
		fmt.Println("url:", url)
		response, err := http.Get(url)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			fmt.Printf("working")
			// data, _ := ioutil.ReadAll(response.Body)
			// fmt.Println(response.StatusCode)
			// fmt.Println(string(data))
			response.Body.Close()

		}
	}

}

// main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/passenger", passengerPage)
	router.HandleFunc("/passengerLogin", passengerLogin)

	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
