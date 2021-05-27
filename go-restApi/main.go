package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(writer http.ResponseWriter, request *http.Request) {

	articles := Articles{
		Article{Title: "Cyclone ", Desc: "Cyclone in Eastern-India", Content: "https://www.google.com"},
		Article{Title: "Cyclone1 ", Desc: "Cyclone in Eastern-India", Content: "https://www.google.com"},
		Article{Title: "Cyclone2", Desc: "Cyclone in Eastern-India", Content: "https://www.google.com"},
	}

	fmt.Println("Endpoint Hit :All Article Endpoint")
	json.NewEncoder(writer).Encode(articles)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint hit")
	fmt.Fprintf(writer, "HomePage EndPoint Hit")

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", router))
}

func testPostArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Test Post article endpoint hit")

}

func main() {
	fmt.Println("Executing the main function")
	handleRequests()

}
