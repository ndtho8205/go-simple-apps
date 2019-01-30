package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
}

func encodeHandler(writer http.ResponseWriter, request *http.Request) {
	user := User{
		FirstName: "Minh",
		LastName:  "Tran",
		Username:  "ttminh",
	}
	err := json.NewEncoder(writer).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func decodeHandler(writer http.ResponseWriter, request *http.Request) {
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	_, err = fmt.Fprintln(writer, user)
	if err != nil {
		log.Fatal(err)
	}
}

func JsonRegister(router *mux.Router) {
	router.HandleFunc("/json/encode", encodeHandler)
	router.HandleFunc("/json/decode", decodeHandler)
}
