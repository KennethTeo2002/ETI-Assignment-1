package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Variables
const PassengerAPIbaseURL = "http://localhost:5000/api/v1/passenger"

type passengerInfo struct {
	Id           string
	Firstname    string
	Lastname     string
	Mobilenumber string
	Email        string
}

// HTML display functions
func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("Website/homepage.html"))

	tmpl.Execute(w, "data goes here")
}

func passengerPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("Website/passengerpage.html"))

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

			response.Body.Close()
			// reroute to passenger home

		}
	}
}
func passengerSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/passengerSignup.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		url := PassengerAPIbaseURL
		passengerDetails := &passengerInfo{
			Id:           r.FormValue("id"),
			Firstname:    r.FormValue("firstname"),
			Lastname:     r.FormValue("lastname"),
			Mobilenumber: r.FormValue("mobileNo"),
			Email:        r.FormValue("email"),
		}
		passengerData := new(passengerInfo)
		passengerData.Id = r.FormValue("id")
		passengerData.Firstname = r.FormValue("firstname")
		passengerData.Lastname = r.FormValue("lastname")
		passengerData.Mobilenumber = r.FormValue("mobileNo")
		passengerData.Email = r.FormValue("email")

		passengerToAdd, _ := json.Marshal(passengerData)

		response, err := http.Post(url+"/"+passengerDetails.Id,
			"application/json", bytes.NewBuffer(passengerToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
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
	router.HandleFunc("/passengerSignup", passengerSignup)

	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
