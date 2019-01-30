package routes

import (
	"github.com/gorilla/mux"
	"go-simple-apps/starterServer/middlewares"
	"html/template"
	"log"
	"net/http"
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

func Todos(router *mux.Router, publicDir string) {
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
