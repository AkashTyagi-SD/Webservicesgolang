package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//This is a util function for create mysql database connection
//This is called from controller file
func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1qazxsw23edc@tcp(localhost:3306)/golangdb")
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	return db, nil
}
