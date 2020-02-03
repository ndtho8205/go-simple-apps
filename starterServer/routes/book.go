package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndtho8205/go-simple-apps/starterServer/middlewares"
)

func booksHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	title, page := vars["title"], vars["page"]
	_, err := fmt.Fprintf(writer, "Book: %s\nPage: %s", title, page)
	if err != nil {
		log.Fatal(err)
	}
}

func BooksRegister(router *mux.Router) {
	handlerFunc := middlewares.Chain(
		booksHandler,
		middlewares.MethodMiddleware(http.MethodGet),
		middlewares.LoggingMiddleware())

	router.HandleFunc("/book/{title}/page/{page}", handlerFunc)
}
