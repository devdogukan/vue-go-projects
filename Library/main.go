package main

import (
	"fmt"
	"log"
	"net/http"

	ws "example/websocket"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bar", ws.WsEndpoint)
	router.HandleFunc("/{id}/books", ws.BooksEndpoint)
	router.HandleFunc("/users", ws.UsersEndpoint)
	fmt.Println("Listen in port 9100")
	log.Fatal(http.ListenAndServe(":9100", router))
}
