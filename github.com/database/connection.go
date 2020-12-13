package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//This is a util function for create mysql database connection
//This is called from controller file
// func CreateConnection() (*sql.DB, error) {
// 	db, err := sql.Open("mysql", "root:1qazxsw23edc@tcp(localhost:3306)/golangdb")
// 	if err != nil {
// 		panic(err.Error())
// 		return nil, err
// 	}
// 	return db, nil
// }

//CreateConnection is called from controller file
func CreateConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:1qazxsw23edc@(localhost:3306)/golangdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	return db, nil
}
