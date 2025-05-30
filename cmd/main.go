package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xxii22w/fullstackgo/handlers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.New().HandleHome).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", "localhost:8080")
	http.ListenAndServe(":8080", router)
}