package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-simple-apps/starterServer/middlewares"
	"log"
	"net/http"
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
