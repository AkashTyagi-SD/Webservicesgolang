package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Create Database connection by calling CreateConnection
func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1qazxsw23edc@tcp(localhost:3306)/dev")
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	return db, nil
}
