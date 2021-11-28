# ETI Assignment 1

Hi, I am Kenneth Teo, a Year 3 student studying for a Diploma in Information Technology at Ngee Ann Polytechnic.

This respository contains the source code for my ETI assignment 1 project. This project is about the implementation of microservices and REST API in a ride-sharing platform.

## Design consideration of microservices

Initial design
[insert image]
During the first draft of my microservice design, I was planning to have the client webapp directly interact with the passenger and driver APIs. And when the user is attempting to post or update a trip, the client would sent a call to their respective APIs which would then forward the call to the Trip API, making it an indirect connection.

In both the passenger and driver APIs, I have a function that takes in their ID as a parameter and their password as a URL query. This function uses all 4 methods (GET, POST, PUT, DELETE) to manipulate their account record information. For the trip API, since passenger and driver require the GET method, I decided to split both of them up into 2 functions. Thus, the trippassenger function with the GET and POST methods are used by the passenger to retrieve their trip history and book a trip. Whereas the tripdriver function with the GET and PUT methods are used by the driver to retrieve their allocated trip and modify the trip record with start and end time.

## Architecture diagram

After several iterations, the final architecture diagram i settled on is
[insert image]

## Instructions for setting up and running your microservices

Setting up of database

Running the microservices
To start the client side application, run `go run main.go` within the main root directory.

To start the microservices, navigate to the respective API folders (PassengerAPI, DriverAPI, TripAPI) and run `go run main.go databaseFunctions.go`
