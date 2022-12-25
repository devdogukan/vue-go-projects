package models

import (
	"database/sql"
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
	AgeRange   int    `json:"age_range"`
}

var databaseName string = "books"

func InsertBook(b Book) int {
	initalize(databaseName)
	defer db.Close()

	insertStr := "INSERT INTO books(no, name, author, book_type, popularity, age_range) VALUES($1, $2, $3, $4, $5, $6)"
	result, err := db.Exec(insertStr, b.No, b.Name, b.Author, b.BookType, b.Popularity, b.AgeRange)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func UpdateBook(id int, b Book) int {
	initalize(databaseName)
	defer db.Close()

	updateStr := "UPDATE books SET no=$2, name=$3, author=$4, book_type=$5, popularity=$5 WHERE id=$1"

	result, err := db.Exec(updateStr, id, b.No, b.Name, b.Author, b.BookType, b.Popularity)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func DeleteBook(id int) int {
	initalize(databaseName)
	defer db.Close()

	result, err := db.Exec("DELETE FROM books WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func DropBooks() int {
	initalize(databaseName)
	defer db.Close()

	result, err := db.Exec("DROP TABLE books")

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected)
}

func GetAllBooks(queryStr string) []Book {
	initalize(databaseName)
	defer db.Close()

	rows, err := db.Query(queryStr)

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
		err = rows.Scan(&book.Id, &book.No, &book.Name, &book.Author, &book.BookType, &book.Popularity, &book.AgeRange)
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
