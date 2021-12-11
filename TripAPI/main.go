package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const DriverAPIbaseURL = "http://localhost:5001/api/v1/driver"

type tripInfo struct {
	Id        string
	CustID    string
	DriverID  string
	PickUp    string
	DropOff   string
	StartTime *time.Time
	EndTime   *time.Time
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


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Trip REST API!")
}

func tripPassenger(w http.ResponseWriter, r *http.Request) {
	
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)

	if r.Method == "GET" {
		if tripArr, ok := GetPassengerTrips(db, params["passengerID"]); ok {
			json.NewEncoder(w).Encode(tripArr)

		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No trips found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {
		var tripdetails tripInfo
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &tripdetails)
			// Check if JSON missing any values
			missingValues := tripdetails.CustID == "" || tripdetails.PickUp == "" || tripdetails.DropOff == ""
			if missingValues {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"422 - Missing trip information "))
				return
			}
			// POST is for creating new trip
			if r.Method == "POST" {

				url := DriverAPIbaseURL + "trip"

				response, err := http.Get(url)
				if err != nil {
					fmt.Printf("The HTTP request failed with error %s\n", err)
				} else {
					data, _ := ioutil.ReadAll(response.Body)
					if string(data) != "" {
						tripdetails.DriverID = string(data)
						InsertRecord(db, tripdetails.CustID, tripdetails.DriverID, tripdetails.PickUp, tripdetails.DropOff)

					}else{
						w.WriteHeader(
							http.StatusUnprocessableEntity)
						w.Write([]byte(
							"422 - No available driver "))
						return
					}

				}

			}

		} else {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Please supply passenger information " +
				"in JSON format"))
		}

	}
}
func tripDriver(w http.ResponseWriter, r *http.Request) {
	
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if trip, ok := GetDriverTrips(db, params["driverID"]); ok {
			json.NewEncoder(w).Encode(trip)

		} else {
			json.NewEncoder(w).Encode(trip)
		}
	}

	if r.Header.Get("Content-type") == "application/json" {
		var tripdetails tripInfo
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &tripdetails)
			// Check if JSON missing any values
			missingValues := tripdetails.Id == ""
			if missingValues {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"422 - Missing trip information "))
				return
			}
			// PUT updates trip details
			if r.Method == "PUT" {
				if !tripdetails.StartTime.IsZero() {
					EditRecord(db, tripdetails.Id, "StartTime")
				} else {
					EditRecord(db, tripdetails.Id, "EndTime")
					jsonString := driverInfo{
						Id: tripdetails.DriverID,
					}
					jsonValue, _ := json.Marshal(jsonString)

					request, _ := http.NewRequest(http.MethodPut,
						DriverAPIbaseURL+"trip",
						bytes.NewBuffer(jsonValue))

					request.Header.Set("Content-Type", "application/json")

					client := &http.Client{}
					_, err := client.Do(request)

					if err != nil {
						fmt.Printf("The HTTP request failed with error %s\n", err)
					}
				}

			}

		} else {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Please supply passenger information " +
				"in JSON format"))
		}

	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/trip/passenger/{passengerID}", tripPassenger).Methods(
		"GET", "POST")
	router.HandleFunc("/api/v1/trip/driver/{driverID}", tripDriver).Methods(
		"GET", "PUT")

	fmt.Println("Listening at port 5002")
	log.Fatal(http.ListenAndServe(":5002", router))
}
