-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost'

-- CREATE database passenger_db;
-- DROP database passenger_db;

USE passenger_db;

CREATE TABLE Passengers (ID varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), MobileNumber VARCHAR(15), EmailAddress VARCHAR(30); 
DROP TABLE Passengers;


INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0001", "Jake", "Lee", "98765432", "JakeLee@gmail.com");
