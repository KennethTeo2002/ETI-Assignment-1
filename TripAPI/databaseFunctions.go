package main

import (
	"database/sql"
	"fmt"
)

func InsertRecord(db *sql.DB, CustID string, DriverID string, PickupLoc string, DropoffLoc string) {
	query := fmt.Sprintf("INSERT INTO Trips (CustomerID, DriverID, PickUp, DropOff) VALUES ('%s','%s', '%s', '%s')",
		CustID, DriverID, PickupLoc, DropoffLoc)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditRecord(db *sql.DB, ID string, CustID string, DriverID string, PickupLoc string, DropoffLoc string, StartTime string, EndTime string) {
	query := fmt.Sprintf(
		"UPDATE Trips SET CustomerID='%s', DriverID='%s', PickUp='%s', DropOff='%s', StartTime='%s',EndTime='%s' WHERE ID = '%s'",
		CustID, DriverID, PickupLoc, DropoffLoc, StartTime, EndTime, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetPassengerTrips(db *sql.DB, CustID string) ([]tripInfo, bool) {
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
		trips = append(trips, trip)

	}

	return trips, true

}

func GetDriverTrips(db *sql.DB, DriverID string) (tripInfo, bool) {

	var trip tripInfo
	query := fmt.Sprintf("Select * FROM trip_db.Trips WHERE DriverID = '%s' AND (StartTime IS NULL OR EndTime IS NULL)", DriverID)

	if err := db.QueryRow(query).Scan(&trip.Id, &trip.CustID, &trip.PickUp, &trip.DropOff, &trip.StartTime, &trip.EndTime); err != nil {
		if err == sql.ErrNoRows {
			return trip, false
		}
	}
	fmt.Println(trip)
	return trip, true

}
