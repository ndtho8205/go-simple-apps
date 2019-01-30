package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func FileServerRegister(router *mux.Router, publicDir string) {
	fs := http.FileServer(http.Dir(publicDir))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public/", fs))
}
