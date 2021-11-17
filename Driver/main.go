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

type driverInfo struct {
	Id             string
	Firstname      string
	Lastname       string
	Mobilenumber   string
	Email          string
	Identification string
	CarLicense     string
}

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

func driver(w http.ResponseWriter, r *http.Request) {
	// if !validKey(r) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("401 - Invalid key"))
	// 	return
	// }
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if driver, ok := GetRecords(db, params["driverID"]); ok {
			json.NewEncoder(w).Encode(driver)

		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No driver found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {
		var newdriver driverInfo
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &newdriver)
			// Check if JSON missing any values
			missingValues := newdriver.Firstname == "" || newdriver.Lastname == "" || newdriver.Mobilenumber == "" || newdriver.Email == "" || newdriver.Identification == "" || newdriver.CarLicense == ""
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
				if _, ok := GetRecords(db, params["driverID"]); !ok {
					InsertRecord(db, newdriver.Id, newdriver.Firstname, newdriver.Lastname, newdriver.Mobilenumber, newdriver.Email, newdriver.Identification, newdriver.CarLicense)

				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate passenger ID"))
				}
			}
			//---PUT is for creating or updating
			// existing passenger---
			if r.Method == "PUT" {

				if _, ok := GetRecords(db, params["driverID"]); !ok {
					InsertRecord(db, newdriver.Id, newdriver.Firstname, newdriver.Lastname, newdriver.Mobilenumber, newdriver.Email, newdriver.Identification, newdriver.CarLicense)

				} else {
					// update passenger
					EditRecord(db, newdriver.Id, newdriver.Firstname, newdriver.Lastname, newdriver.Mobilenumber, newdriver.Email, newdriver.CarLicense)

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
	router.HandleFunc("/api/v1/driver/{driverID}", driver).Methods(
		"GET", "PUT", "POST")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
