package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dgnk1234"
	dbname   = "productsdb"
)

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connString)

	//db.SetMaxIdleConns(5)  // db'ye bagli ama query(sorgu) yapmayan cihaz sinirini belirler.
	//db.SetMaxOpenConns(10)  // es zamanli kac adet connection olacagini belirler
	//db.SetConnMaxIdleTime(1 * time.Second)  // bir bağlantının boşta kalabileceği maksimum süreyi ayarlar.
	//db.SetConnMaxLifetime(30 * time.Second) // bir bağlantının yeniden kullanılabileceği maksimum süreyi ayarlar.

	if err != nil {
		log.Fatal(err)
	}
}

type Product struct {
	ID          int
	ProductName string
	CategoryId  int
	Price       float32
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(productName, categoryId, price) VALUES($1, $2, $3)", data.ProductName, data.CategoryId, data.Price)

	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of Records Affected: %d", rowsAffected)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE products SET productName=$2, categoryID=$3, price=$4 WHERE id=$1", data.ID, data.ProductName, data.CategoryId, data.Price)

	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of Records Affected: %d", rowsAffected)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.ProductName, &prd.CategoryId, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, prd := range products {
		fmt.Printf("%d --> %s, %d, %v\n", prd.ID, prd.ProductName, prd.CategoryId, prd.Price)
	}
}

func GetProductById(id int) {
	var product Product
	err := db.QueryRow("SELECT * FROM products WHERE id=$1", id).Scan(&product.ID, &product.ProductName, &product.CategoryId, &product.Price)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %v", product)
	}
}
