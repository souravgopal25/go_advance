package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

//[]byte
var mySigningKey []byte = []byte(os.Getenv("MY_JWT_TOKEN"))

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Sourav Sharma"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong", err)
		return "", err
	}
	return tokenString, nil

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello WOrld")
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func main() {
	fmt.Println("My Simple Client")
	handleRequests()

}
