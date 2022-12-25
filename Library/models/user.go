package models

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Age               int    `json:"age"`
	Access_permission bool   `json:"access_permission"`
}

func GetUserByID(id int) bool {
	initalize("register")
	defer db.Close()

	var user User
	err := db.QueryRow("SELECT access_permission FROM users WHERE id=$1", id).Scan(&user.Access_permission)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return user.Access_permission
}

func UpdateUser(user User) int {
	initalize("register")
	defer db.Close()

	result, err := db.Exec("UPDATE users SET access_permission=$2 WHERE id=$1", user.ID, user.Access_permission)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func GetAllUsers() []User {
	initalize("register")
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Record Not Found")
			return nil
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Age, &user.Access_permission)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
