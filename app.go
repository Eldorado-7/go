package main

import (
	_ "encoding/json"
	Address "go-microservices/Lib/Core/Addresses"
	_ "io/ioutil"
	_ "log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create new Router
	router := mux.NewRouter()
	// route properly to respective handlers
	//router.Handle("/employees", Employees.Employee.Run()).Methods("GET")
	//router.Handle("/employees", handlers.CreateProductHandler()).Methods("POST")

	// Create new server and assign the router
	server := http.Server{
		Addr:    Address.HOST_DEFAULT_PORT,
		Handler: router,
	}

	// Start Server on defined port/host.
	server.ListenAndServe()
}

//Get request handler

func doGet(response http.ResponseWriter, request *http.Request) {

}

//Post request handler
