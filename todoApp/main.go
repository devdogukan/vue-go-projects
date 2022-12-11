package main

import (
	"fmt"
	"log"
	"net/http"
	tdb "vue2/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
	todo   tdb.Todo
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Couldnt not upgrade %s\n", err.Error())
		return
	}

	todos := tdb.GetAllTodo()
	for _, key := range todos {
		err := wsConn.WriteMessage(websocket.TextMessage, []byte(key.Title))
		if err != nil {
			log.Fatal(err)
		}
	}

	defer wsConn.Close()

	for {

		err := wsConn.ReadJSON(&todo)
		if err != nil {
			fmt.Printf("error reading json %s\n", err.Error())
			continue
		}

		fmt.Printf("Todo Recived: %s\n", todo.Title)

		if todo.Procces == 0 {
			tdb.InsertTodo(todo)
		} else if todo.Procces == 1 {
			tdb.DeleteTodoById(todo.Id)
		} else if todo.Procces == 2 {
			tdb.UpdateTodo(todo)
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/socket", WsEndpoint)

	fmt.Println("Listening on port 9100")
	log.Fatal(http.ListenAndServe(":9100", router))
}
