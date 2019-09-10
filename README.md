# go-adega
Simple CRUD using Rest and [GO](https://github.com/golang/go)


### Basic usage

Server is set to up on localhost:8000

Application starts on main.go, the main function load the routes file and set a mock on the server file initializing it with two wines on the wines list.

adega/route directory

Routes were built with package [gorilla/mux](https://github.com/gorilla/mux). These routes are bound to the resource file methods corresponding to http requests.

adega/resource directory

Resource file takes the responsability over the http requests, it calls the service file who has a specific method to deal with each different request. It takes the request parameters and pass it to the service, and then return the service response.

adega/service directory

Service file it is similar to the resource file, but this one deals only with the rules over the wine model.
A few basic tests were implemented for this file.

adega/model directory

Model file holds the custom types used on the application.

