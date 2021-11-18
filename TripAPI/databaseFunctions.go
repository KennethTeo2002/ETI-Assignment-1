package main

import (
	"database/sql"
	"fmt"
)

func InsertRecord(db *sql.DB, CustID string, DriverID string, PickupLoc string, DropoffLoc string, StartTime string, EndTime string) {
	query := fmt.Sprintf("INSERT INTO Trips (CustomerID, DriverID, PickUp, DropOff, StartTime, EndTime) VALUES ('%s','%s', '%s', '%s', '%s', '%s')",
		CustID, DriverID, PickupLoc, DropoffLoc, StartTime, EndTime)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditRecord(db *sql.DB, ID string, FN string, LN string, MobileNumber string, Email string, CarLicense string) {
	query := fmt.Sprintf(
		"UPDATE Trips SET FirstName='%s', LastName='%s', MobileNumber='%s', EmailAddress='%s', CarLicense='%s' WHERE ID = '%s'",
		FN, LN, MobileNumber, Email, CarLicense, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB, CustID string) ([]tripInfo, bool) {
	var trips []tripInfo
	query := fmt.Sprintf("Select * FROM trip_db.Trips WHERE CustomerID = '%s'", CustID)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var trip tripInfo
		err = results.Scan(&trip.Id, &trip.CustID, &trip.DriverID, &trip.PickUp, &trip.DropOff, &trip.StartTime, &trip.EndTime)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(trip.Id, trip.CustID,
			trip.DriverID)

	}

	return trips, true

}
