package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "todosdb"
	password = "dgnk1234"
)

var db *sql.DB

func init() {
	var err error
	connstring := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connstring)

	if err != nil {
		log.Fatal(err)
	}
}

type Todo struct {
	Id         int
	Title      string
	IsDone     bool
	lastUpdate string
}

func InsertTodo(data Todo) int {
	result, err := db.Exec("INSERT INTO todos(title, isDone) VALUES($1, $2)", data.Title, data.IsDone)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func UpdateTodo(data Todo) int {
	result, err := db.Exec("UPDATE todos SET title=$2, isDone=$3 WHERE id=$1", data.Id, data.Title, data.IsDone)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func GetTodoById(id int) string {
	var todo Todo
	err := db.QueryRow("SELECT * FROM todos WHERE ID=$1", id).Scan(&todo.Id, &todo.Title, &todo.IsDone, &todo.lastUpdate)

	switch {
	case err == sql.ErrNoRows:
		return "Rows Not Found"
	case err != nil:
		log.Fatal(err)
	default:
		result := fmt.Sprintf("id=%d, title=%s, isDone=%t, last update=%s", todo.Id, todo.Title, todo.IsDone, todo.lastUpdate)
		return result
	}
	return ""
}

func GetAllTodos() []Todo {
	rows,err := db.Query("SELECT * FROM todos")

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Record Not Found")
			return nil
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.IsDone, &todo.lastUpdate)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return todos
}

func DeleteById(id int) int {
	result, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected,_ := result.RowsAffected()

	return int(rowsAffected)
}
