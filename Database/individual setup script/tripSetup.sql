-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost'

-- CREATE database trip_db;
-- DROP database trip_db;

USE trip_db;

-- CREATE TABLE Trips (ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT, CustomerID VARCHAR(30), DriverID VARCHAR(30), PickUp VARCHAR(6), DropOff VARCHAR(6), StartTime DateTime, EndTime DateTime); 
-- DROP TABLE Trips;

-- INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0001", "0001", "693222", "489293", '2021-11-19 12:34:05', '2021-11-19 12:36:45');
-- INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0001", "0003", "489293", "693222", '2021-11-19 14:42:10', '2021-11-19 15:02:31');
-- INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0004", "0001", "413232", "454533", '2021-11-20 8:56:21', '2021-11-20 9:00:25');
-- INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0005", "0002", "654352", "567453", '2021-11-20 9:05:42', '2021-11-20 9:05:45');

Select * FROM trip_db.Trips ;