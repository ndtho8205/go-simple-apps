package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func passwordHandler(writer http.ResponseWriter, request *http.Request) {
	password := request.URL.Query().Get("password")
	checkPassword := request.URL.Query().Get("checkPassword")

	result := ""
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		log.Fatal(err)
	}

	result += "Hash: " + string(hash[:]) + "\n"

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(checkPassword))
	result += fmt.Sprintf("Match result: %t", err == nil)

	_, err = fmt.Fprintln(writer, result)
	if err != nil {
		log.Fatal(err)
	}
}

func PasswordRegister(router *mux.Router) {
	router.HandleFunc("/password", passwordHandler)
}
