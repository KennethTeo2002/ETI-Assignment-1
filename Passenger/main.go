package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type passengerInfo struct {
	Id           string
	Firstname    string
	Lastname     string
	Mobilenumber string
	Email        string
}

var passengers map[string]passengerInfo

/*
func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
*/

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

func passenger(w http.ResponseWriter, r *http.Request) {
	// if !validKey(r) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("401 - Invalid key"))
	// 	return
	// }
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if passenger, ok := GetRecords(db, params["passengerID"]); ok {
			json.NewEncoder(w).Encode(passenger)

		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {
		var newpassenger passengerInfo
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &newpassenger)
			// Check if JSON missing any values
			missingValues := newpassenger.Firstname == "" || newpassenger.Lastname == "" || newpassenger.Mobilenumber == "" || newpassenger.Email == ""
			if missingValues {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"422 - Missing passenger information "))
				return
			}
			// POST is for creating new passenger
			if r.Method == "POST" {
				// check if course exists; add only if
				// course does not exist
				if _, ok := passengers[params["passengerID"]]; !ok {
					InsertRecord(db, newpassenger.Id, newpassenger.Firstname, newpassenger.Lastname, newpassenger.Mobilenumber, newpassenger.Email)

					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - passenger added: " +
						params["passengerID"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate passenger ID"))
				}
			}
			//---PUT is for creating or updating
			// existing passenger---
			if r.Method == "PUT" {

				if _, ok := GetRecords(db, params["passengerID"]); !ok {
					InsertRecord(db, newpassenger.Id, newpassenger.Firstname, newpassenger.Lastname, newpassenger.Mobilenumber, newpassenger.Email)

				} else {
					// update passenger
					fmt.Println("updatiing")
					EditRecord(db, newpassenger.Id, newpassenger.Firstname, newpassenger.Lastname, newpassenger.Mobilenumber, newpassenger.Email)

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
	passengers = make(map[string]passengerInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/passenger/{passengerID}", passenger).Methods(
		"GET", "PUT", "POST")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
