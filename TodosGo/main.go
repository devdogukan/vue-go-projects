package main

import (
	"fmt"

	db "github.com/pg/db.com/models"
)

func main()  {
	todo := db.Todo{Title: "Alisveris yap", IsDone: true}

	fmt.Println(db.InsertTodo(todo))
	//fmt.Println(db.UpdateTodo(todo))
	//fmt.Println(db.GetTodoById(1))
	//fmt.Println(db.DeleteById(1))

	//fmt.Println(db.GetAllTodos())
}