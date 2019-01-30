package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-simple-apps/starterServer/middlewares"
	"log"
	"net/http"
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
