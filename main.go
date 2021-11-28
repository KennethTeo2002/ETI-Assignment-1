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
	Password	string
	Firstname    string
	Lastname     string
	Mobilenumber string
	Email        string
}
type driverInfo struct {
	Id             string
	Password string
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
// Temporary caching
var passenger passengerInfo
var driver driverInfo
var activeTrip tripInfo

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Website/homepage.html"))
	tmpl.Execute(w, nil)
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Website/errorpage.html"))
	tmpl.Execute(w, nil)
}

// Passenger webpages
func passengerHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerHome.html"))
	tmpl.Execute(w, passenger)
}

func passengerUpdateCookie(id string, password string)bool{
	url:= PassengerAPIbaseURL
	if id != "" && password != ""  {
		url = PassengerAPIbaseURL + "/" + id + "?password=" + password
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != http.StatusOK{
			passenger = passengerInfo{}
		}else{
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal([]byte(data), &passenger)
			return true
		}
		
	}
	return false
}

func passengerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerLogin.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		id := r.FormValue("id")
		password := r.FormValue("password")

		success := passengerUpdateCookie(id,password)
		if success{
			http.Redirect(w, r, "/passenger", http.StatusFound)
		}else{
			http.Redirect(w, r, "/error", http.StatusFound)
		}
		

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
		passengerData.Password = r.FormValue("password")
		passengerData.Firstname = r.FormValue("firstname")
		passengerData.Lastname = r.FormValue("lastname")
		passengerData.Mobilenumber = r.FormValue("mobileNo")
		passengerData.Email = r.FormValue("email")

		passengerToAdd, _ := json.Marshal(passengerData)

		response, err := http.Post(url+"/"+passengerData.Id+ "?password=" + passengerData.Password,
			"application/json", bytes.NewBuffer(passengerToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode != http.StatusOK {
				http.Redirect(w, r, "/error", http.StatusFound)
			}else{
				success := passengerUpdateCookie(passengerData.Id,passengerData.Password)
				if success{
					http.Redirect(w, r, "/passenger", http.StatusFound)
				}else{
					http.Redirect(w, r, "/error", http.StatusFound)
				}
			}
		}
	}
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
			url+"/"+passenger.Id + "?password=" + passenger.Password,
			bytes.NewBuffer(passengerToUpdate))

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode != http.StatusOK {
				http.Redirect(w, r, "/error", http.StatusFound)
			}else{
				success := passengerUpdateCookie(passenger.Id,passenger.Password)
				if success{
					http.Redirect(w, r, "/passenger", http.StatusFound)
				}else{
					http.Redirect(w, r, "/error", http.StatusFound)
				}
			}
		}
	}
}

func passengerViewTrips(w http.ResponseWriter, r *http.Request) {	
	url := PassengerAPIbaseURL
	passengerID := passenger.Id
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
	
	tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerViewTrip.html"))
	tmpl.Execute(w, trips)
}

func passengerRequestTrip(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Passenger/passengerRequestTrip.html"))
		tmpl.Execute(w, passenger)
	} else {
		r.ParseForm()
		tripData := new(tripInfo)
		tripData.CustID = passenger.Id
		tripData.PickUp = r.FormValue("pickup")
		tripData.DropOff = r.FormValue("dropoff")

		tripToAdd, _ := json.Marshal(tripData)
		response, err := http.Post(TripAPIbaseURL+"/passenger/"+tripData.CustID,
			"application/json", bytes.NewBuffer(tripToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode != http.StatusOK{
				http.Redirect(w, r, "/error", http.StatusFound)
			}else{
				http.Redirect(w, r, "/passenger", http.StatusFound)
			}
		}
	}
}

func passengerDeleteAccount(w http.ResponseWriter, r *http.Request) {
	request, _ := http.NewRequest(http.MethodDelete,
        PassengerAPIbaseURL+"/"+passenger.Id+"?password="+passenger.Password, nil)

    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		if response.StatusCode != http.StatusOK{
			http.Redirect(w, r, "/error", http.StatusFound)
		}else{
			http.Redirect(w, r, "/passengerLogin", http.StatusFound)
		}
    }
}


// Driver webpages
func driverHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
		response, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode != http.StatusOK{
				http.Redirect(w, r, "/error", http.StatusFound)
				
			}else{
				success := driverUpdateCookie(driver.Id,driver.Password)
				if success{
					http.Redirect(w, r, "/driver", http.StatusFound)
				}else{
					http.Redirect(w, r, "/error", http.StatusFound)
				}
			}

		}

	}

}
func driverUpdateCookie(id string, password string)bool{
	driverSuccess := false

	url := DriverAPIbaseURL
	if id != "" && password != "" {
		url = DriverAPIbaseURL + "/" + id + "?password=" + password
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode != http.StatusOK{
			driver = driverInfo{}
			// display alert msg
			_=string(data)
		}else{
			json.Unmarshal([]byte(data), &driver)
			driverSuccess = true
		}

	}

	if id != ""  && password != "" {
		url = TripAPIbaseURL + "/driver/" + id
	}
	responsetrip, errtrip := http.Get(url)

	if errtrip != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(responsetrip.Body)
		if responsetrip.StatusCode != http.StatusOK{
			activeTrip = tripInfo{}
			// display alert msg
			_=string(data)
		}else{
			json.Unmarshal([]byte(data), &activeTrip)
		}
	}
	return driverSuccess
}

func driverLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("Website/Driver/driverLogin.html"))
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		id := r.FormValue("id")
		password := r.FormValue("password")

		success := driverUpdateCookie(id,password)
		if success{
			http.Redirect(w, r, "/driver", http.StatusFound)
		}else{
			http.Redirect(w, r, "/error", http.StatusFound)
		}
		

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
		driverData.Password = r.FormValue("password")
		driverData.Firstname = r.FormValue("firstname")
		driverData.Lastname = r.FormValue("lastname")
		driverData.Mobilenumber = r.FormValue("mobileNo")
		driverData.Email = r.FormValue("email")
		driverData.Identification = r.FormValue("identification")
		driverData.CarLicense = r.FormValue("carLicense")

		driverToAdd, _ := json.Marshal(driverData)

		response, err := http.Post(url+"/"+driverData.Id,
			"application/json", bytes.NewBuffer(driverToAdd))

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode != http.StatusOK{
				http.Redirect(w, r, "/error", http.StatusFound)
			}else{
				success := driverUpdateCookie(driverData.Id,driverData.Password)
				if success{
					http.Redirect(w, r, "/driver", http.StatusFound)
				}else{
					http.Redirect(w, r, "/error", http.StatusFound)
				}
			}

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
		driverData.Password = driver.Password
		driverData.Firstname = r.FormValue("firstname")
		driverData.Lastname = r.FormValue("lastname")
		driverData.Mobilenumber = r.FormValue("mobileNo")
		driverData.Email = r.FormValue("email")
		driverData.CarLicense = r.FormValue("carLicense")
		driverToUpdate, _ := json.Marshal(driverData)

		request, _ := http.NewRequest(http.MethodPut,
			url+"/"+driver.Id + "?password=" + driver.Password,
			bytes.NewBuffer(driverToUpdate))

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if response.StatusCode != http.StatusOK{
				http.Redirect(w, r, "/error",http.StatusFound)
			}else{
				success := driverUpdateCookie(driverData.Id,driverData.Password)
				if success{
					http.Redirect(w, r, "/driver", http.StatusFound)
				}else{
					http.Redirect(w, r, "/error", http.StatusFound)
				}
			}
		}

	}
}

func driverDeleteAccount(w http.ResponseWriter, r *http.Request) {
	request, _ := http.NewRequest(http.MethodDelete,
        DriverAPIbaseURL+"/"+driver.Id+"?password="+driver.Password, nil)

    client := &http.Client{}
    response, err := client.Do(request)
	
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		if response.StatusCode != http.StatusOK{
			http.Redirect(w, r, "/error", http.StatusFound)
		}else{
			http.Redirect(w, r, "/driverLogin", http.StatusOK)
		}
        
    }

}

// main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/error", errorPage)

	// routes for passenger
	router.HandleFunc("/passengerLogin", passengerLogin)
	router.HandleFunc("/passengerSignup", passengerSignup)
	router.HandleFunc("/passenger", passengerHome)
	router.HandleFunc("/passenger/editPDetails", passengerEditDetails)
	router.HandleFunc("/passenger/viewTrips", passengerViewTrips)
	router.HandleFunc("/passenger/requestTrip", passengerRequestTrip)
	router.HandleFunc("/passenger/deleteAccount", passengerDeleteAccount)

	// routes for driver
	router.HandleFunc("/driverLogin", driverLogin)
	router.HandleFunc("/driverSignup", driverSignup)
	router.HandleFunc("/driver", driverHome)
	router.HandleFunc("/driver/editDDetails", driverEditDetails)
	router.HandleFunc("/driver/deleteAccount", driverDeleteAccount)

	// html assets
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("Website/css/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/",
		http.FileServer(http.Dir("Website/img/"))))

	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
