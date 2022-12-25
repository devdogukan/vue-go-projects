package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "books"
	password = "dgnk1234"
)

func initalize(nameOfDatabase string) {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, nameOfDatabase, password)

	db, err = sql.Open(user, connString)

	if err != nil {
		log.Fatal(err)
	}
}
