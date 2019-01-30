package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-simple-apps/starterServer/routes"
	"log"
	"net/http"
	"time"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello! %v", r.URL.Query().Get("key"))
	if err != nil {
		log.Fatal(err)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "About")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()

	routes.BooksRegister(router)
	routes.Todos(router, "public/")
	routes.ContactRegister(router, "public/")

	server := http.Server{
		Addr:         "localhost:9000",
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(server.ListenAndServe())
}

//	fs := http.FileServer(http.Dir("public/"))

//	r.HandleFunc("/", RootHandler)
//	r.HandleFunc("/about", AboutHandler)
//	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

//	r.HandleFunc("/forms", FormsHandler)
