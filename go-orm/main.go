package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUsers).Methods("PUT")
	myRouter.HandleFunc("user/{name}", DeleteUsers).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func main() {
	fmt.Printf("Executing Main Function")
	InitialMigration()
	handleRequests()

}
