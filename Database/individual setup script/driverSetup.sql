-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost'
-- CREATE database driver_db;
-- DROP database driver_db;
USE driver_db;

-- CREATE TABLE Drivers (ID varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), MobileNumber VARCHAR(15), EmailAddress VARCHAR(30), Identification VARCHAR(30), CarLicense VARCHAR(30), Driving BOOL); 
-- DROP TABLE Drivers;


-- INSERT INTO Drivers (ID, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0001", "Bron", "son", "65423324", "bonzey@hotmail.com", "S2391782Z","SGL2782D",false);

SELECT * FROM Drivers