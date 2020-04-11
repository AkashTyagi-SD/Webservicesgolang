package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1qazxsw23edc@tcp(localhost:3306)/dev")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("db is connected")
		return db, nil
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err == nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}
