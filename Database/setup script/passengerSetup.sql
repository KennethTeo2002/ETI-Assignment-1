CREATE user 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost'

CREATE database passenger_db;

USE passenger_db;

CREATE TABLE Passengers (ID varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), Age INT); 

INSERT INTO Persons (ID, FirstName, LastName, Age) VALUES ("0001", "Jake", "Lee", 25);
