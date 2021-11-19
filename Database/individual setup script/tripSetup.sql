-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost'

-- CREATE database trip_db;
-- DROP database trip_db;

USE trip_db;

-- CREATE TABLE Trips (ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT, CustomerID VARCHAR(30), DriverID VARCHAR(30), PickUp VARCHAR(6), DropOff VARCHAR(6), StartTime DateTime, EndTime DateTime); 
-- DROP TABLE Trips;



Select * FROM trip_db.Trips ;