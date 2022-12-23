package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	bdb "example/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Couldn't not upgrade %s\n", err.Error())
		return
	}

	books := bdb.GetAllBooks()
	jsonBook, _ := json.Marshal(books)
	err = wsConn.WriteMessage(websocket.TextMessage, []byte(jsonBook))
	if err != nil {
		log.Fatal(err)
	}

	// for _, key := range books {
	// 	key, _ := json.Marshal(key)
	// 	err := wsConn.WriteMessage(websocket.TextMessage, []byte(key))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	defer wsConn.Close()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bar", WsEndpoint)
	fmt.Println("Listen in port 9100")
	log.Fatal(http.ListenAndServe(":9100", router))
}
