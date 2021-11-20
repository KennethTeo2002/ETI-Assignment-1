package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Variables
const PassengerAPIbaseURL = "http://localhost:5000/api/v1/passenger"
const DriverAPIbaseURL = "http://localhost:5001/api/v1/driver"
const TripAPIbaseURL = "http://localhost:5002/api/v1/trip"

type passengerInfo struct {
	Id           string
	Firstname    string
	Lastname     string
	Mobilenumber string
	Email        string
}
type driverInfo struct {
	Id             string
	Firstname      string
	Lastname       string
	Mobilenumber   string
	Email          string
	Identification string
	CarLicense     string
	Driving        bool
}
type tripInfo struct {
	Id        string
	CustID    string
	DriverID  string
	PickUp    string
	DropOff   string
	StartTime time.Time
	EndTime   time.Time
}
type DriverHomeData struct {
	Driver     driverInfo
	ActiveTrip tripInfo
}

var passenger passengerInfo
var driver driverInfo
var activeTrip tripInfo

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

func driverHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		params := mux.Vars(r)

		url := DriverAPIbaseURL
		driverID := params["driverID"]
		if driverID != "" {
			url = DriverAPIbaseURL + "/" + driverID
		}
		response, err := http.Get(url)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal([]byte(data), &driver)

		}

		if driverID != "" {
			url = TripAPIbaseURL + "/driver/" + driverID
		}
		responsetrip, errtrip := http.Get(url)

		if errtrip != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(responsetrip.Body)
			json.Unmarshal([]byte(data), &activeTrip)
		}

		pageData := DriverHomeData{
			Driver:     driver,
			ActiveTrip: activeTrip,
		}

		tmpl := template.Must(template.ParseFiles("Website/Driver/driverHome.html"))

		tmpl.Execute(w, pageData)
	} else {
		r.ParseForm()
		tripUpdate := new(tripInfo)
		tripUpdate.Id = activeTrip.Id
		tripUpdate.DriverID = activeTrip.DriverID

		if r.FormValue("start") == "Start Trip" {
			tripUpdate.StartTime = time.Now()
		} else {
			tripUpdate.EndTime = time.Now()
		}

		tripToUpdate, _ := json.Marshal(tripUpdate)

		request, _ := http.NewRequest(http.MethodPut,
			TripAPIbaseURL+"/driver/"+tripUpdate.DriverID,
			bytes.NewBuffer(tripToUpdate))

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			redirectURL := fmt.Sprintf("/driver/%s", tripUpdate.DriverID)

			http.Redirect(w, r, redirectURL, http.StatusFound)

		}

	}

}

func driverEditDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Driver/driverEdit.html"))
		tmpl.Execute(w, driver)
	} else {
		r.ParseForm()
		url := DriverAPIbaseURL
		driverData := new(driverInfo)
		driverData.Id = driver.Id
		driverData.Firstname = r.FormValue("firstname")
		driverData.Lastname = r.FormValue("lastname")
		driverData.Mobilenumber = r.FormValue("mobileNo")
		driverData.Email = r.FormValue("email")
		driverData.CarLicense = r.FormValue("carLicense")
		driverToUpdate, _ := json.Marshal(driverData)

		request, _ := http.NewRequest(http.MethodPut,
			url+"/"+driver.Id,
			bytes.NewBuffer(driverToUpdate))

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {

			redirectURL := fmt.Sprintf("/driver/%s", driver.Id)

			http.Redirect(w, r, redirectURL, http.StatusFound)
		}

	}
}
func driverLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Driver/driverLogin.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		id := r.FormValue("id")
		redirectURL := fmt.Sprintf("/driver/%s", id)

		http.Redirect(w, r, redirectURL, http.StatusFound)

	}
}

func driverSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Driver/driverSignup.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		url := DriverAPIbaseURL

		driverData := new(driverInfo)
		driverData.Id = r.FormValue("id")
		driverData.Firstname = r.FormValue("firstname")
		driverData.Lastname = r.FormValue("lastname")
		driverData.Mobilenumber = r.FormValue("mobileNo")
		driverData.Email = r.FormValue("email")
		driverData.Identification = r.FormValue("identification")
		driverData.CarLicense = r.FormValue("carLicense")

		driverToAdd, _ := json.Marshal(driverData)

		_, err := http.Post(url+"/"+driverData.Id,
			"application/json", bytes.NewBuffer(driverToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			redirectURL := fmt.Sprintf("/driver/%s", driverData.Id)

			http.Redirect(w, r, redirectURL, http.StatusFound)

		}
	}
}

func passengerViewTrips(w http.ResponseWriter, r *http.Request) {
	// todo: get trip array from trip api
	params := mux.Vars(r)
	url := PassengerAPIbaseURL
	passengerID := params["passengerID"]
	if passengerID != "" {
		url = TripAPIbaseURL + "/passenger/" + passengerID
	}
	response, err := http.Get(url)
	var trips []tripInfo
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &trips)
	}
	fmt.Println(trips)

	tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerViewTrip.html"))
	tmpl.Execute(w, trips)
}

func passengerRequestTrip(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerRequestTrip.html"))
		tmpl.Execute(w, passenger)
	} else {
		r.ParseForm()
		params := mux.Vars(r)
		tripData := new(tripInfo)
		tripData.CustID = params["passengerID"]
		tripData.PickUp = r.FormValue("pickup")
		tripData.DropOff = r.FormValue("dropoff")

		tripToAdd, _ := json.Marshal(tripData)
		_, err := http.Post(TripAPIbaseURL+"/passenger/"+tripData.CustID,
			"application/json", bytes.NewBuffer(tripToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			redirectURL := fmt.Sprintf("/passenger/%s", tripData.CustID)

			http.Redirect(w, r, redirectURL, http.StatusFound)

		}
	}

}

// main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)

	// routes for passenger
	router.HandleFunc("/passengerLogin", passengerLogin)
	router.HandleFunc("/passengerSignup", passengerSignup)
	router.HandleFunc("/passenger/{passengerID}", passengerHome)
	router.HandleFunc("/passenger/{passengerID}/editPDetails", passengerEditDetails)
	router.HandleFunc("/passenger/{passengerID}/viewTrips", passengerViewTrips)
	router.HandleFunc("/passenger/{passengerID}/requestTrip", passengerRequestTrip)

	// routes for driver
	router.HandleFunc("/driverLogin", driverLogin)
	router.HandleFunc("/driverSignup", driverSignup)
	router.HandleFunc("/driver/{driverID}", driverHome)
	router.HandleFunc("/driver/{driverID}/editDDetails", driverEditDetails)

	// html assets
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("Website/css/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/",
		http.FileServer(http.Dir("Website/img/"))))

	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
