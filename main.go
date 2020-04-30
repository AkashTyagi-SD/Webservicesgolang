package main

import (
	"log"
	"net/http"

	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/controller"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter()
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/getuser/{userid}", controller.Getuser).Methods("GET")

	http.ListenAndServe(":8080", router)

}
