package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndtho8205/go-simple-apps/starterServer/middlewares"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello! %v", r.URL.Query().Get("key"))
	if err != nil {
		log.Fatal(err)
	}
}

func HomeRegister(router *mux.Router) {
	router.
		HandleFunc("/", middlewares.Chain(homeHandler, middlewares.LoggingMiddleware())).
		Methods(http.MethodGet)
}
