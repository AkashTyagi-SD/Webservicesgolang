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
	router.HandleFunc("/posts", controller.Index).Methods("GET")

	http.ListenAndServe(":8080", router)

	// var empdata1 models.Emp
	// empdata1.Name = "Akash"
	// empdata1.Address = "Khilwai"
	// empdata1.Age = 20
	// showData(empdata1)

}
