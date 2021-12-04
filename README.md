# ETI Assignment 1

Hi, I am Kenneth Teo, a Year 3 student studying for a Diploma in Information Technology at Ngee Ann Polytechnic.

This respository contains the source code for my ETI assignment 1 project. This project is about the implementation of microservices and REST API in a ride-sharing platform.

## Design consideration of microservices

<ins>Initial design</ins>

![Initial Design](design1.png?raw=true "Title")
During the first draft of my microservice design, I was planning to have the client webapp only able to directly interact with the passenger and driver APIs. And when the user is attempting to post or update a trip, the client would sent a call to their respective user APIs which would then forward the call to the Trip API, making it an indirect connection. However, after completing the passenger API, I realised that all 4 methods of the passenger function was used, which means I would need to create a new function just to forward the call. This was redundant as the trip information needed from passenger was only the ID, making the rerouting not efficient, so I decided that the client webapp should connect to the trip API directly.

## Architecture diagram

After several iterations, the final architecture diagram i settled on is

![Initial Design](design2.png?raw=true "Title")

### Passenger

| API Action                   | API Command                    |
| ---------------------------- | ------------------------------ |
| Get passenger information    | `GET /api/v1/passenger/:id`    |
| Add new passenger            | `POST /api/v1/passenger/:id`   |
| Update passenger information | `PUT /api/v1/passenger/:id`    |
| Delete passenger             | `DELETE /api/v1/passenger/:id` |

The Passenger microservice uses all 4 methods available in 1 function. This function would take in the passenger's id as a parameter, as well as their password key as a URL query. The password variable cannot be set as a parameter, as for the POST method, since the user currently do not have an account, there would not be any records to refer to, thus the password query is not supplied. Whereas for the GET, PUT and DELETE methods, the API would read and validate the user before performing their respective actions.

### Driver

| API Action                   | API Command                 |
| ---------------------------- | --------------------------- |
| Get driver information       | `GET /api/v1/driver/:id`    |
| Add new driver               | `POST /api/v1/driver/:id`   |
| Update driver information    | `PUT /api/v1/driver/:id`    |
| Delete driver                | `DELETE /api/v1/driver/:id` |
| Get available driver         | `GET /api/v1/drivertrip`    |
| Update driver's availability | `PUT /api/v1/drivertrip`    |

The Driver microservice uses 2 functions, `.../driver` is the base function for the user's account, whereas `.../drivertrip` is the function used by the trip API to assign drivers. The base function uses all 4 methods available, as works the same way as the passenger API. The other function only uses the GET and PUT methods, and is only connected and called by the trip API. The GET function searches the database for an available driver, set their availability to 'not available' and returns their ID back to the trip API. The PUT function takes in the driver json object, set the driver's availability back to 'available'.

### Trip

| API Action                               | API Command                       |
| ---------------------------------------- | --------------------------------- |
| Get passenger's trip history             | `GET /api/v1/trip/passenger/:id`  |
| Request new trip                         | `POST /api/v1/trip/passenger/:id` |
| Get driver's active trip                 | `GET /api/v1/trip/driver/:id`     |
| Update trip timing & driver availability | `PUT /api/v1/trip/driver/:id`     |

The trip microservice uses 2 functions, `.../passenger/:id` is the function for passengers, whereas `.../driver/:id` is the function for drivers. The passenger function only uses the GET and POST methods, the GET functionality retrieves and return all trip records in descending order when it is called by the passenger entering the 'view trip history' page. The POST method is invoked when a passenger submit a request for a trip. Firstly, this method would call the driver API to get an available driver, and set the trip's allocated driver to the driver id and posting the trip to the database.

The driver function only uses the GET and PUT methods, the GET functionality checks if there is any active trip under the driver that does not have an end time, and returns it. The PUT method is triggered when the driver clicks on the 'Start trip' or 'End trip' button, which would update the respective database column with the current time. If the call was to end the trip, a call is also sent to the driver API to set the availabilty back to 'available'.

## Instructions for setting up and running your microservices

### Setting up of database

To set up the 3 MySQL databases, connect `fullSetup.sql` from the Database directory to a instance of MySQL, and run the file.

### Running the microservices

To start the client side application, run `go run main.go` within the main root directory.

To start the microservices, navigate to the respective API folders (PassengerAPI, DriverAPI, TripAPI) and run `go run main.go databaseFunctions.go`
