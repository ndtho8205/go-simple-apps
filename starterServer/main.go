package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ndtho8205/go-simple-apps/starterServer/routes"
)

func main() {
	publicDir := "public/"
	router := mux.NewRouter()

	routes.BooksRegister(router)
	routes.TodosRegister(router, publicDir)
	routes.ContactRegister(router, publicDir)
	routes.FileServerRegister(router, publicDir)
	routes.MarkdownRegister(router, publicDir)
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
