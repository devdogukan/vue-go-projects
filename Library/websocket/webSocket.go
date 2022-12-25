package websocket

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

func BooksEndpoint(w http.ResponseWriter, r *http.Request) {

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

	access_permission := bdb.GetUserByID(i)

	if access_permission {
		queryStr = "SELECT * FROM books;"
	} else {
		queryStr = "SELECT * FROM books WHERE age_range < 18;"
	}

	books := bdb.GetAllBooks(queryStr)
	jsonBook, _ := json.Marshal(books)
	err = wsConn.WriteMessage(websocket.TextMessage, []byte(jsonBook))
	if err != nil {
		log.Fatal(err)
	}

	defer wsConn.Close()
}

func UsersEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Couldn't not upgrade %s\n", err.Error())
		return
	}
	defer wsConn.Close()

	users := bdb.GetAllUsers()
	jsonBook, _ := json.Marshal(users)
	err = wsConn.WriteMessage(websocket.TextMessage, []byte(jsonBook))
	if err != nil {
		log.Fatal(err)
	}

	var user []bdb.User

	err = wsConn.ReadJSON(&user)
	if err != nil {
		fmt.Printf("error reading json %s\n", err.Error())
	}

	for i := 0; i < len(user); i++ {
		bdb.UpdateUser(user[i])
	}
}
