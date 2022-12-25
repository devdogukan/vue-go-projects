package models

import "fmt"

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func GetUserByID(id int) int {
	initalize("register")
	var user User
	err := db.QueryRow("SELECT age FROM users WHERE id=$1", id).Scan(&user.Age)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return user.Age
}
