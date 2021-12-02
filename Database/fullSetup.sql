-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost';

CREATE database passenger_db;
USE passenger_db;

CREATE TABLE Passengers (ID varchar (5) NOT NULL PRIMARY KEY,  Password VARCHAR(30), FirstName VARCHAR(30), LastName VARCHAR(30), MobileNumber VARCHAR(15), EmailAddress VARCHAR(30)); 
INSERT INTO Passengers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0001", "password", "Cannerf", "Back", "11111111", "bosspro2002@gmail.com");
INSERT INTO Passengers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0002", "password1", "Kleb", "banana", "22222222", "clapfruit@gmail.com");
INSERT INTO Passengers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0003", "password2", "Oh hak", "eews", "33333333", "jaycemain@hotmail.com");
INSERT INTO Passengers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0004", "password3", "Neinhility", "rart", "44444444", "nine9nein@gmail.com");
INSERT INTO Passengers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0005", "password4", "Axie", "OOF", "55555555", "oofer@hotmail.com");
INSERT INTO Passengers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0006", "password5", "Canshiryou", "The", "66666666", "overflowprez@gmail.com");

CREATE database driver_db;
USE driver_db;

CREATE TABLE Drivers (ID varchar (5) NOT NULL PRIMARY KEY, Password VARCHAR(30), FirstName VARCHAR(30), LastName VARCHAR(30), MobileNumber VARCHAR(15), EmailAddress VARCHAR(30), Identification VARCHAR(30), CarLicense VARCHAR(30), Driving BOOL); 
INSERT INTO Drivers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0001", "pass1", "Bron", "son", "65423324", "bonzey@hotmail.com", "S2391782Z","SGL2782D",false);
INSERT INTO Drivers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0002", "pass2", "Vander", "wick", "74542214", "definitelynotwarwick@gmail.com", "S8952857C","SGB8402Q",false);
INSERT INTO Drivers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0003", "pass3", "Jinx", "powder", "56234523", "powpow@gmail.com", "S0856867A","SGS2460D",false);

CREATE database trip_db;
USE trip_db;

CREATE TABLE Trips (ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT, CustomerID VARCHAR(30), DriverID VARCHAR(30), PickUp VARCHAR(6), DropOff VARCHAR(6), StartTime DateTime, EndTime DateTime); 
INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0001", "0001", "693222", "489293", '2021-11-19 12:34:05', '2021-11-19 12:36:45');
INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0001", "0003", "489293", "693222", '2021-11-19 14:42:10', '2021-11-19 15:02:31');
INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0004", "0001", "413232", "454533", '2021-11-20 8:56:21', '2021-11-20 9:00:25');
INSERT INTO Trips (CustomerID,DriverID,PickUp,DropOff,StartTime,EndTime) VALUES ("0005", "0002", "654352", "567453", '2021-11-20 9:05:42', '2021-11-20 9:05:45');
