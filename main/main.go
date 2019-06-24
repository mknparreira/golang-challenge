package main

import (
	"fmt"
	"golang-challenge/endpoints"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter().StrictSlash(true)
	subroute := route.PathPrefix("/challenger/api/v1").Subrouter()

	subroute.HandleFunc("/", endpoints.IndexHandler).Methods("GET")
	subroute.HandleFunc("/from/{fromRange}/to/{toRange}",
		endpoints.RangeHandler).Methods("GET")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", subroute)
}
