CREATE TABLE products (
	id 			SERIAL PRIMARY KEY,
	productName VARCHAR(100) NOT NULL,
	categoryId  NUMERIC(1,0) NOT NULL,
	price		NUMERIC(6,3) NOT NULL
);

INSERT INTO products(productName, categoryId, price) VALUES('Laptop', 1, 11.500)

SELECT * FROM products;

DROP TABLE products;