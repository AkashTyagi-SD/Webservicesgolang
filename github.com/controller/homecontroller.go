package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/database"
	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/models"
)

//This function used for fetch data without input param
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.CreateConnection()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	defer selDB.Close()
	emp := models.Employee{}
	res := []models.Employee{}
	for selDB.Next() {
		var id int
		var name, email, address, telephone string
		err = selDB.Scan(&id, &name, &email, &address, &telephone)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Email = email
		emp.Address = address
		emp.Telephone = telephone
		res = append(res, emp)
	}
	json.NewEncoder(w).Encode(res)

}

//Login function
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	var result models.User
	var res models.ResponseResult

}
