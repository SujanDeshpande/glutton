package main

import (
	"glutton/Handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", Handler.HomeHandler).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}
