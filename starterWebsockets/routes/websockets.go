package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketsEchoHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(message))

		if err = conn.WriteMessage(messageType, message); err != nil {
			log.Fatal(err)
		}
	}
}

func websocketsHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "public/websockets.html")
}

func WebsocketsRegister(router *mux.Router) {
	router.HandleFunc("/websockets", websocketsHandler)
	router.HandleFunc("/websockets/echo", websocketsEchoHandler)
}
