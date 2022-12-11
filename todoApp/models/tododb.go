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

var Db *sql.DB

func init() {
	var err error
	constring := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	Db, err = sql.Open("postgres", constring)

	if err != nil {
		log.Fatal(err)
	}
}

type Todo struct {
	Procces int    `json:"procces"`
	Id      int    `json:"id"`
	Title   string `json:"title"`
	IsDone  bool   `json:"isDone"`
}

func InsertTodo(todo Todo) int {
	result, err := Db.Exec("INSERT INTO todos(title, is_done) VALUES($1, $2)", todo.Title, todo.IsDone)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func GetTodoById(id int) string {
	var todo Todo
	err := Db.QueryRow("SELECT * FROM todos WHERE id=$1", id).Scan(&todo.Id, &todo.Title, &todo.IsDone)

	if err == sql.ErrNoRows {
		return "Rows not found"
	} else if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	result := fmt.Sprintf("id=%d title=%s", todo.Id, todo.Title)
	return result
}

func GetAllTodo() []Todo {
	rows, err := Db.Query("SELECT * FROM todos")

	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal("Rows not found")
			return nil
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.IsDone)
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

func UpdateTodo(todo Todo) int {
	result, err := Db.Exec("UPDATE todos SET is_done=$2 WHERE id=$1", todo.Id, todo.IsDone)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected)
}

func DeleteTodoById(id int) int {
	result, err := Db.Exec("DELETE FROM todos WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}
