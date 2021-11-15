package main

import (
	"database/sql"
	"fmt"
)

func InsertRecord(db *sql.DB, MobileNumber string, FN string, LN string, Email string) {
	query := fmt.Sprintf("INSERT INTO Passengers (MobileNumber, FirstName, LastName, EmailAddress) VALUES (%s, '%s', '%s', %s)",
		MobileNumber, FN, LN, Email)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditRecord(db *sql.DB, ID string, MobileNumber string, FN string, LN string, Email string) {
	query := fmt.Sprintf(
		"UPDATE Passengers SET FirstName='%s', LastName='%s', Email=%s, MobileNumber=%s WHERE ID = '%s",
		FN, LN, Email, MobileNumber, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB, ID string) {
	results, err := db.Query("Select * FROM my_db.Passengers WHERE ID = '%s", ID)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		// map this type to the record in the table
		// var passenger Passenger
		// err = results.Scan(&person.ID, &person.FirstName,
		// 	&person.LastName, &person.Age)
		if err != nil {
			panic(err.Error())
		}

		// fmt.Println(person.ID, person.FirstName,
		// 	person.LastName, person.Age)
	}

}
