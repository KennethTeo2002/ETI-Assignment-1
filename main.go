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

var passenger passengerInfo

// Webpages
func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("Website/homepage.html"))

	tmpl.Execute(w, nil)
}

func passengerHome(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := PassengerAPIbaseURL
	passengerID := params["passengerID"]
	if passengerID != "" {
		url = PassengerAPIbaseURL + "/" + passengerID
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &passenger)
	}
	tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerHome.html"))

	tmpl.Execute(w, passenger)

}
func passengerEditDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerEdit.html"))
		tmpl.Execute(w, passenger)
	} else {
		r.ParseForm()
		url := PassengerAPIbaseURL
		passengerData := new(passengerInfo)
		passengerData.Id = passenger.Id
		passengerData.Firstname = r.FormValue("firstname")
		passengerData.Lastname = r.FormValue("lastname")
		passengerData.Mobilenumber = r.FormValue("mobileNo")
		passengerData.Email = r.FormValue("email")

		passengerToUpdate, _ := json.Marshal(passengerData)

		request, _ := http.NewRequest(http.MethodPut,
			url+"/"+passenger.Id,
			bytes.NewBuffer(passengerToUpdate))

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {

			redirectURL := fmt.Sprintf("/passenger/%s", passenger.Id)

			http.Redirect(w, r, redirectURL, http.StatusFound)
		}

	}
}
func passengerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerLogin.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		id := r.FormValue("id")
		redirectURL := fmt.Sprintf("/passenger/%s", id)

		http.Redirect(w, r, redirectURL, http.StatusFound)

	}
}

func passengerSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerSignup.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		url := PassengerAPIbaseURL

		passengerData := new(passengerInfo)
		passengerData.Id = r.FormValue("id")
		passengerData.Firstname = r.FormValue("firstname")
		passengerData.Lastname = r.FormValue("lastname")
		passengerData.Mobilenumber = r.FormValue("mobileNo")
		passengerData.Email = r.FormValue("email")

		passengerToAdd, _ := json.Marshal(passengerData)

		_, err := http.Post(url+"/"+passengerData.Id,
			"application/json", bytes.NewBuffer(passengerToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			redirectURL := fmt.Sprintf("/passenger/%s", passengerData.Id)

			http.Redirect(w, r, redirectURL, http.StatusFound)

		}
	}
}

// main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)

	//routes for passenger
	router.HandleFunc("/passengerLogin", passengerLogin)
	router.HandleFunc("/passengerSignup", passengerSignup)
	router.HandleFunc("/passenger/{passengerID}", passengerHome)
	router.HandleFunc("/passenger/{passengerID}/editPDetails", passengerEditDetails)

	//routes for driver
	router.HandleFunc("/driverLogin", passengerLogin)
	router.HandleFunc("/driverSignup", passengerSignup)
	router.HandleFunc("/driver/{driverID}", passengerHome)
	router.HandleFunc("/driver/{driverID}/editDDetails", passengerEditDetails)

	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
