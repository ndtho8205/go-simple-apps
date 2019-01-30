package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

func sessionsHandler(writer http.ResponseWriter, request *http.Request) {
	session1, err := store.Get(request, "session-one")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	session1.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	session1.Values["MaxAge"] = "0"

	session2, err := store.Get(request, "session-two")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	session2.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60,
		HttpOnly: true,
	}
	session2.Values["MaxAge"] = "60"

	err = sessions.Save(request, writer)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(writer, "Hello!")
	if err != nil {
		log.Fatal(err)
	}
}

func SessionsRegister(router *mux.Router) {
	router.HandleFunc("/sessions", sessionsHandler)
}
