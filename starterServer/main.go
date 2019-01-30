package main

import (
	"github.com/gorilla/mux"
	"go-simple-apps/starterServer/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	routes.BooksRegister(router)
	routes.TodosRegister(router, "public/")
	routes.ContactRegister(router, "public/")
	routes.FileServerRegister(router, "public/")
	routes.AboutRegister(router)
	routes.HomeRegister(router)

	server := http.Server{
		Addr:         "localhost:9000",
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(server.ListenAndServe())
}
