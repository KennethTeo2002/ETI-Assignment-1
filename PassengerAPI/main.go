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
	Password string
	Firstname    string
	Lastname     string
	Mobilenumber string
	Email        string
}


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Passenger REST API!")
}

func validPassword(r *http.Request, db *sql.DB ,id string ) bool {
    v := r.URL.Query()
    if password, ok := v["password"]; ok {
		if passenger, ok := GetRecords(db, id); ok {
			if password[0] == passenger.Password {
				return true
			}else{
				return false
			}
		} else {
			return false
		}
    } else {
        return false
    }
}

func passenger(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	params := mux.Vars(r)
	if r.Method == "GET" {
		if !validPassword(r,db,params["passengerID"]) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Invalid credentials"))
			return
		}
		if passenger, ok := GetRecords(db, params["passengerID"]); ok {
			json.NewEncoder(w).Encode(passenger)

		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found"))
		}
	}
	if r.Method == "DELETE" {
		if !validPassword(r,db,params["passengerID"]) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Invalid credentials"))
			return
		}
        w.WriteHeader(http.StatusForbidden)
        w.Write([]byte("403 - Unable to delete due to auditing reasons"))
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
				// check if passenger exists; add only if
				// passenger does not exist
				if _, ok := GetRecords(db, params["passengerID"]); !ok {
					InsertRecord(db, newpassenger.Id, newpassenger.Password,newpassenger.Firstname, newpassenger.Lastname, newpassenger.Mobilenumber, newpassenger.Email)

				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate passenger ID"))
				}
			}
			//---PUT is for creating or updating
			// existing passenger---
			if r.Method == "PUT" {
				if !validPassword(r,db,params["passengerID"]) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 - Invalid credentials"))
					return
				}
				if _, ok := GetRecords(db, params["passengerID"]); !ok {
					InsertRecord(db, newpassenger.Id, newpassenger.Password, newpassenger.Firstname, newpassenger.Lastname, newpassenger.Mobilenumber, newpassenger.Email)

				} else {
					// update passenger
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
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passenger", home)
	router.HandleFunc("/api/v1/passenger/{passengerID}", passenger).Methods(
		"GET", "PUT", "POST","DELETE")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
