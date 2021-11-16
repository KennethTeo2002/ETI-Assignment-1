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

	fmt.Fprintf(w, "Detail for passenger "+params["passengerID"])
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, r.Method)

	if r.Method == "GET" {
		GetRecords(db, params["passengerID"])
		if _, ok := passengers[params["passengerID"]]; ok {
			json.NewEncoder(w).Encode(
				passengers[params["passengerID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new passenger
		if r.Method == "POST" {

			// read the string sent to the service
			var newpassenger passengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newpassenger)

				if newpassenger.Firstname == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply passenger " +
							"information " + "in JSON format"))
					return
				}

				// check if course exists; add only if
				// course does not exist
				if _, ok := passengers[params["passengerID"]]; !ok {
					InsertRecord(db, newpassenger.Firstname, newpassenger.Lastname, newpassenger.Mobilenumber, newpassenger.Email)
					passengers[params["passengerID"]] = newpassenger
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - passenger added: " +
						params["passengerID"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate passenger ID"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply passenger information " +
					"in JSON format"))
			}
		}

		//---PUT is for creating or updating
		// existing passenger---
		if r.Method == "PUT" {
			var newpassenger passengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newpassenger)

				if newpassenger.Firstname == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply passenger " +
							" information " +
							"in JSON format"))
					return
				}

				// check if passenger exists; add only if
				// passenger does not exist
				if _, ok := passengers[params["passengerID"]]; !ok {
					passengers[params["passengerID"]] =
						newpassenger
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - passenger added: " +
						params["passengerID"]))
				} else {
					// update passenger
					passengers[params["passengerID"]] = newpassenger
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - passenger updated: " +
						params["passengerID"]))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"passenger information " +
					"in JSON format"))
			}
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
