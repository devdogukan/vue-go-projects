package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	Id         int    `json:"id"`
	No         int    `json:"no"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	BookType   string `json:"type"`
	Popularity int    `json:"popularity"`
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "books"
	password = "dgnk1234"
)

var db *sql.DB

func initalize() {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}
}

func InsertBook(b Book) int {
	initalize()
	defer db.Close()

	result, err := db.Exec("INSERT INTO books(no, name, author, book_type, popularity) VALUES($1, $2, $3, $4, $5)", b.No, b.Name, b.Author, b.BookType, b.Popularity)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func UpdateBook(id int, b Book) int {
	initalize()
	defer db.Close()

	result, err := db.Exec("UPDATE books SET no=$2, name=$3, author=$4, book_type=$5, popularity=$5 WHERE id=$1", id, b.No, b.Name, b.Author, b.BookType, b.Popularity)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func DeleteBook(id int) int {
	initalize()
	defer db.Close()

	result, err := db.Exec("DELETE FROM books WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func DropBooks() int {
	initalize()
	defer db.Close()

	result, err := db.Exec("DROP TABLE books")

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func GetAllBooks() []Book {
	initalize()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Rows not found")
			return nil
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Id, &book.No, &book.Name, &book.Author, &book.BookType, &book.Popularity)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return books
}
