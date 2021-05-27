package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "log"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err := connectToDatabse()
	if err != nil {
		panic(err)
		return
	}

	//Migrating the schema
	er := db.AutoMigrate(&User{})
	if er != nil {
		panic(er)
	} else {
		fmt.Printf("Migration Success")
	}
	//var db1 *gorm.DB
	//Create
	/*if err := db.Create(&User{Name: "Sourav Sharma ", Email: "souravgopal25@gmail.com"}).Error; err != nil {
		panic(err)
	}*/

}
func connectToDatabse() (*gorm.DB, error) {
	dbDriver := "mysql"
	dbUser := "Sourav"
	dbPass := "rootpass11"
	dbName := "studentDB"
	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                dbDriver,
		DSN:                       dsn,
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}), &gorm.Config{})
	if err != nil {
		panic("error")
		fmt.Printf("Failed to connect to Database")
		return nil, err
	}
	return db, nil
}
func AllUsers(writter http.ResponseWriter, reader *http.Request) {
	var db *gorm.DB
	var err error
	db, err = connectToDatabse()
	if err != nil {
		panic(err)
	}

	var users []User
	result := db.Find(&users)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	json.NewEncoder(writter).Encode(users)

}

func NewUsers(writer http.ResponseWriter, request *http.Request) {
	var db *gorm.DB
	var err error
	db, err = connectToDatabse()
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(request)
	if err := db.Create(&User{Name: vars["name"], Email: vars["email"]}).Error; err != nil {
		panic(err)
	}
	fmt.Fprintf(writer, "New Users Successfully Created")
}

func DeleteUsers(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Executing Delete Function")
	var db *gorm.DB
	var err error
	db, err = connectToDatabse()
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(request)
	name := vars["name"]

	var user []User
	//soft delete
	er := db.Where("Name=?", name).Delete(&user).Error

	if er != nil {
		fmt.Fprintf(writer, er.Error())
	}
	//TODO Learn how to handle the ouptut and check whether user has been deleted or not.

	/*	fmt.Println(name)
		fmt.Println(strconv.Itoa(int(results.RowsAffected)))

		if results.RowsAffected!=0 {
			db.Delete(&user)
			fmt.Fprintf(writer, strconv.Itoa(int(results.RowsAffected)) +"User Successfully deleted")
		}else{
			fmt.Fprintf(writer, "No User Found in Database")
		}*/

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var db *gorm.DB
	var err error
	db, err = connectToDatabse()
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(request)
	name := vars["name"]
	email := vars["email"]
	var user User
	//soft delete
	er := db.Where("Name=?", name).Find(&user)
	if er.Error != nil {
		panic(er)
	}
	user.Email = email
	db.Save(&user)

	fmt.Fprintf(writer, " User Updated Success")
}
