-- CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
-- GRANT ALL ON *.* TO 'user'@'localhost'

-- DROP database passenger_db;
-- CREATE database passenger_db;


USE passenger_db;

-- DROP TABLE Passengers;
-- CREATE TABLE Passengers (ID varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), MobileNumber VARCHAR(15), EmailAddress VARCHAR(30)); 



-- INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0001", "Cannerf", "Back", "11111111", "bosspro2002@gmail.com");
-- INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0002", "Kleb", "banana", "22222222", "clapfruit@gmail.com");
-- INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0003", "Oh hak", "eews", "33333333", "jaycemain@hotmail.com");
-- INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0004", "Neinhility", "rart", "44444444", "nine9nein@gmail.com");
-- INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0005", "Axie", "OOF", "55555555", "oofer@hotmail.com");
-- INSERT INTO Passengers (ID, FirstName, LastName, MobileNumber,EmailAddress) VALUES ("0006", "Canshiryou", "The", "66666666", "overflowprez@gmail.com");

SELECT * FROM Passengers;