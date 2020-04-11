package main

import (
	"fmt"

	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/database"
)

func main() {
	fmt.Println("Hello custome workspace")
	db, err := database.CreateConnection()
	fmt.Println("db", db)
	fmt.Println("err", err)

}
