package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = "5432"
	dbname = "persondb"
	user = "postgres"
	password = "dgnk1234"
)

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
}

var db *sql.DB

func init()  {
	var err error
	db, err := sql.Open("host=%s port=%s dbname=%s user=%s password=%s, sslomde=disable", host, port, dbname, user, )
}
