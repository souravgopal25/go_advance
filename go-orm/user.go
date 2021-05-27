package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB
var err error

func InitialMigration() {
	//	dbDriver := "mysql"
	dbUser := "Sourav"
	dbPass := "rootpass11"
	dbName := "usertable"
	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
		fmt.Printf("Failed to connect to Database")
	}

	//Migrating the schema
	db.AutoMigrate(&User{})

	//Create
	db.Create(&User{Name: "Sourav Sharma", Email: "souravgopal25@gmail.com"})

}

func AllUsers(writter http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writter, "All Users Endpoint Hit")
}

func NewUsers(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "New Users Endpoint hit")
}

func DeleteUsers(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Delete user Endpoint Hit")
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Update User Endpoint Hit")
}
