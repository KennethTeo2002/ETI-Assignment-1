-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost'
-- CREATE database driver_db;
-- DROP database driver_db;
USE driver_db;

-- CREATE TABLE Drivers (ID varchar (5) NOT NULL PRIMARY KEY, Password VARCHAR(30), FirstName VARCHAR(30), LastName VARCHAR(30), MobileNumber VARCHAR(15), EmailAddress VARCHAR(30), Identification VARCHAR(30), CarLicense VARCHAR(30), Driving BOOL); 
-- DROP TABLE Drivers;


INSERT INTO Drivers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0001", "pass1", "Bron", "son", "65423324", "bonzey@hotmail.com", "S2391782Z","SGL2782D",false);
INSERT INTO Drivers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0002", "pass2", "Vander", "wick", "74542214", "definitelynotwarwick@gmail.com", "S8952857C","SGB8402Q",false);
INSERT INTO Drivers (ID, Password, FirstName, LastName, MobileNumber,EmailAddress,Identification,CarLicense,Driving) VALUES ("0003", "pass3", "Jinx", "powder", "56234523", "powpow@gmail.com", "S0856867A","SGS2460D",false);

SELECT * FROM Drivers;

