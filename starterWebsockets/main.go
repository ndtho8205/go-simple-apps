package main

import (
	"github.com/gorilla/mux"
	"go-simple-apps/starterWebsockets/routes"
	"log"
	"net/http"
	"time"
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
