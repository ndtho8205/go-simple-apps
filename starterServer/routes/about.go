package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndtho8205/go-simple-apps/starterServer/middlewares"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "About")
	if err != nil {
		log.Fatal(err)
	}
}

func AboutRegister(router *mux.Router) {
	router.
		HandleFunc("/about", middlewares.Chain(aboutHandler, middlewares.LoggingMiddleware())).
		Methods(http.MethodGet)
}
