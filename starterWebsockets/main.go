package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ndtho8205/go-simple-apps/starterWebsockets/routes"
)

func main() {
	router := mux.NewRouter()

	routes.SessionsRegister(router)
	routes.JsonRegister(router)
	routes.WebsocketsRegister(router)
	routes.PasswordRegister(router)

	server := http.Server{
		Addr:         "localhost:9000",
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(server.ListenAndServe())
}
