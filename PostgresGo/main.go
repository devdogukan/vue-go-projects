package main

import db "dredogu.com/dbgo/models"

func main() {
	//product := db.Product{ProductName: "Mouse", CategoryId: 2, Price: 79.99}
	//db.InsertProduct(product)

	//db.GetProducts()

	//db.GetProductById(2)

	product := db.Product{ID: 2, ProductName: "Mouse", CategoryId: 2, Price: 5.99}
	db.UpdateProduct(product)
}
