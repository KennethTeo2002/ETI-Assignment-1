package main

import (
	"database/sql"
	"fmt"
)

func InsertRecord(db *sql.DB, ID string, FN string, LN string, MobileNumber string, Email string, Identification string, CarLicense string) {
	query := fmt.Sprintf("INSERT INTO Drivers (ID,FirstName, LastName,MobileNumber, EmailAddress,Identification,CarLicense,Driving) VALUES ('%s','%s', '%s', '%s', '%s', '%s', '%s',false)",
		ID, FN, LN, MobileNumber, Email, Identification, CarLicense)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditRecord(db *sql.DB, ID string, FN string, LN string, MobileNumber string, Email string, CarLicense string) {
	query := fmt.Sprintf(
		"UPDATE Drivers SET FirstName='%s', LastName='%s', MobileNumber='%s', EmailAddress='%s', CarLicense='%s' WHERE ID = '%s'",
		FN, LN, MobileNumber, Email, CarLicense, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB, ID string) (driverInfo, bool) {
	var driver driverInfo
	query := fmt.Sprintf("Select * FROM driver_db.Drivers WHERE ID = '%s'", ID)
	if err := db.QueryRow(query).Scan(&driver.Id, &driver.Firstname,
		&driver.Lastname, &driver.Mobilenumber, &driver.Email, &driver.Identification, &driver.CarLicense, &driver.Driving); err != nil {
		if err == sql.ErrNoRows {
			return driver, false
		}
	}
	return driver, true

}

func GetAvailableDriver(db *sql.DB) (driverInfo, bool) {
	var driver driverInfo
	if err := db.QueryRow("Select * FROM driver_db.Drivers WHERE Driving = false").Scan(&driver.Id, &driver.Firstname,
		&driver.Lastname, &driver.Mobilenumber, &driver.Email, &driver.Identification, &driver.CarLicense, &driver.Driving); err != nil {
		if err == sql.ErrNoRows {
			return driver, false
		}
	}
	return driver, true
}

func ToggleDriving(db *sql.DB, ID string) {
	query := fmt.Sprintf(
		"UPDATE Drivers SET Driving = NOT Driving WHERE ID = '%s'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

}
