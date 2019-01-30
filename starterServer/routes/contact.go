package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-simple-apps/starterServer/middlewares"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func ContactRegister(router *mux.Router, publicDir string) {
	tmpl := template.Must(template.ParseFiles(publicDir + "contact.html"))
	contactHandler := func(writer http.ResponseWriter, request *http.Request) {
		result := struct {
			Success bool
			Contact *ContactDetails
		}{false, nil}

		if request.Method == http.MethodPost {
			contact := ContactDetails{
				Email:   request.FormValue("email"),
				Subject: request.FormValue("subject"),
				Message: request.FormValue("message"),
			}
			fmt.Println(contact)
			result.Success = true
			result.Contact = &contact
		}
		err := tmpl.Execute(writer, result)
		if err != nil {
			log.Fatal(err)
		}
	}

	router.HandleFunc(
		"/contact",
		middlewares.Chain(contactHandler, middlewares.LoggingMiddleware()))
}
