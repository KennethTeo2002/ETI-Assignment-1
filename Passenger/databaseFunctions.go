package main

import (
	"database/sql"
	"fmt"
)

func InsertRecord(db *sql.DB, ID string, FN string, LN string, MobileNumber string, Email string) {
	query := fmt.Sprintf("INSERT INTO Passengers (ID,FirstName, LastName,MobileNumber, EmailAddress) VALUES ('%s','%s', '%s', '%s', '%s')",
		ID, FN, LN, MobileNumber, Email)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditRecord(db *sql.DB, ID string, FN string, LN string, MobileNumber string, Email string) {
	query := fmt.Sprintf(
		"UPDATE Passengers SET FirstName='%s', LastName='%s', MobileNumber='%s', Email='%s' WHERE ID = '%s'",
		FN, LN, MobileNumber, Email, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB, ID string) {
	query := fmt.Sprintf("Select * FROM passenger_db.Passengers WHERE ID = '%s'", ID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		// map this type to the record in the table
		var passenger passengerInfo
		err = results.Scan(&passenger.Id, &passenger.Firstname,
			&passenger.Lastname, &passenger.Mobilenumber, &passenger.Email)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(passenger.Firstname,
			passenger.Lastname, passenger.Mobilenumber, passenger.Email)
		// return passenger
	}

}
