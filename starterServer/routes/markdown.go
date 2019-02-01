package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/russross/blackfriday.v2"
	"log"
	"net/http"
)

func markdownGetHandler(publicDir string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, publicDir+"markdown.html")
	}
}

func markdownPostHandler(writer http.ResponseWriter, request *http.Request) {
	var content struct {
		Content string `json:"content"`
	}
	err := json.NewDecoder(request.Body).Decode(&content)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err = fmt.Fprintln(writer, string(blackfriday.Run([]byte(content.Content))))
	if err != nil {
		log.Fatal(err)
	}
}

func MarkdownRegister(router *mux.Router, publicDir string) {
	router.HandleFunc("/markdown", markdownGetHandler(publicDir)).Methods(http.MethodGet)
	router.HandleFunc("/markdown", markdownPostHandler).Methods(http.MethodPost)
}
