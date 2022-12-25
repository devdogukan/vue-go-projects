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
	defer wsConn.Close()

	books := bdb.GetAllBooks("SELECT * FROM books")
	jsonBook, _ := json.Marshal(books)
	err = wsConn.WriteMessage(websocket.TextMessage, []byte(jsonBook))
	if err != nil {
		log.Fatal(err)
	}
}

func booksEndpoint(w http.ResponseWriter, r *http.Request) {

	urlParam := mux.Vars(r)
	id := urlParam["id"]
	fmt.Printf("Hello: %v", id)
	fmt.Println()

	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Couldn't not upgrade %s\n", err.Error())
		return
	}

	var i int
	var queryStr string

	fmt.Sscanf(id, "%d", &i)

	age := bdb.GetUserByID(i)

	if age < 18 {
		queryStr = "SELECT * FROM books WHERE age_range >= 13 and age_range < 18;"
	} else if age >= 18 {
		queryStr = "SELECT * FROM books WHERE age_range >= 18;"
	} else {
		queryStr = "SELECT * FROM books"
	}

	books := bdb.GetAllBooks(queryStr)
	jsonBook, _ := json.Marshal(books)
	err = wsConn.WriteMessage(websocket.TextMessage, []byte(jsonBook))
	if err != nil {
		log.Fatal(err)
	}

	defer wsConn.Close()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bar", WsEndpoint)
	router.HandleFunc("/{id}/books", booksEndpoint)
	fmt.Println("Listen in port 9100")
	log.Fatal(http.ListenAndServe(":9100", router))
}
