package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/database"
	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/models"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Login function
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	//db, err := database.CreateConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//var result models.User
	//var res models.ResponseResult

}

//Register function
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res models.ResponseResult
	if err != nil {
		res.Status = false
		res.Error = err.Error()
		res.Message = "Registration Unsuccessful"
		json.NewEncoder(w).Encode(res)
		return
	}
	db, err := database.CreateConnection()
	if err != nil {
		res.Status = false
		res.Error = err.Error()
		res.Message = "Registration Unsuccessful"
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	query, err := db.Prepare("Insert User SET firstname=?, lastname=? ,emailid=? ,password=? ,isremember=?")
	catch(err)
	_, er := query.Exec(user.FirstName, user.LastName, user.EmailID, user.Password, user.IsRemember)
	if er != nil {
		catch(er)
	} else {
		res.Status = true
		res.Message = "Registration Successful"
		response := []models.User{}

		selDB, err := db.Query("SELECT * FROM User ORDER BY userid DESC LIMIT 1")
		if err != nil {
			catch(er)
		} else {
			defer selDB.Close()
			for selDB.Next() {
				user := models.User{}
				var userid int
				var firstname, lastname, emailid, password string
				var isremember bool
				err = selDB.Scan(&userid, &firstname, &lastname, &emailid, &password, &isremember)
				if err != nil {
					panic(err.Error())
				}
				user.UserID = userid
				user.FirstName = firstname
				user.LastName = lastname
				user.EmailID = emailid
				user.Password = password
				user.IsRemember = isremember
				response = append(response, user)
				res.Result = response
			}
		}
		json.NewEncoder(w).Encode(res)
	}

	defer query.Close()
	//respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Registration Successful"})
}

//Getuser function used for fetch data without input param
func Getuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.CreateConnection()
	selDB, err := db.Query("SELECT userid,firstname,lastname FROM User ORDER BY userid DESC")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	defer selDB.Close()
	user := models.User{}
	var userid int
	var firstname, lastname string
	res := []models.User{}
	for selDB.Next() {

		err = selDB.Scan(&userid, &firstname, &lastname)
		if err != nil {
			panic(err.Error())
		}
		user.UserID = userid
		user.FirstName = firstname
		user.LastName = lastname
		res = append(res, user)
	}
	json.NewEncoder(w).Encode(res)

}
