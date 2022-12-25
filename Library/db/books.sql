CREATE TABLE books (
	id SERIAL PRIMARY KEY,
	no INTEGER NOT NULL,
	name VARCHAR(64) NOT NULL,
	author VARCHAR(64) NOT NULL,
	book_type VARCHAR(64) NOT NULL,
	popularity INTEGER NOT NULL,
	age_range INTEGER NOT NULL
);

INSERT INTO books(no, name, author, book_type, popularity, age_range) 
VALUES(100, 'Kürk Mantolu Madonna', 'Sabahattin Ali', 'Roman', 552, 13)

INSERT INTO books(no, name, author, book_type, popularity, age_range) 
VALUES(101, 'Korku', 'Dre', 'Korku-Gerilim', 199, 18)

SELECT * FROM books;

SELECT * FROM books WHERE age_range >= 18;

SELECT * FROM books WHERE age_range >= 13 and age_range < 18;

DROP TABLE books;