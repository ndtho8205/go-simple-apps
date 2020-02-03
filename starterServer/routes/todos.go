package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndtho8205/go-simple-apps/starterServer/middlewares"
)

type TodoStatus string

const (
	Working TodoStatus = "WORKING"
	Done    TodoStatus = "DONE"
	Cancel  TodoStatus = "CANCEL"
)

type Todo struct {
	Title  string
	Status TodoStatus
}

type TodosPageData struct {
	PageTitle string
	Todos     []Todo
}

func TodosRegister(router *mux.Router, publicDir string) {
	tmpl := template.Must(template.ParseFiles(publicDir + "todos.html"))

	todosData := TodosPageData{
		PageTitle: "TODO",
		Todos: []Todo{
			{"Task 1", Working},
			{"Task 2", Working},
			{"Task 3", Done},
			{"Task 4", Cancel},
		},
	}

	todosHandler := func(writer http.ResponseWriter, request *http.Request) {
		err := tmpl.Execute(writer, todosData)
		if err != nil {
			log.Fatal(err)
		}
	}

	handlerFunc := middlewares.Chain(
		todosHandler,
		middlewares.MethodMiddleware(http.MethodGet),
		middlewares.LoggingMiddleware())

	router.HandleFunc("/todos", handlerFunc)
}
