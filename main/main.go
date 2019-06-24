package main

import (
	"net/http"
	"challenger/endpoints"
	"github.com/gorilla/mux"
	"fmt"
)

func main()  {
	route := mux.NewRouter().StrictSlash(true)
	subroute := route.PathPrefix("/challenger/api/v1").Subrouter()

	subroute.HandleFunc("/", endpoints.IndexHandler).Methods("GET")
	subroute.HandleFunc("/from/{fromRange}/to/{toRange}", 
	endpoints.RangeHandler).Methods("GET")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", subroute)
}
