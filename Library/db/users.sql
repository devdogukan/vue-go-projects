CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	surname VARCHAR(30) NOT NULL,
	age INTEGER NOT NULL
);

INSERT INTO users(name, surname, age) VALUES('Dogukan', 'Aksoy', 20);
INSERT INTO users(name, surname, age) VALUES('Mehmet', 'Solak', 19);
INSERT INTO users(name, surname, age) VALUES('Ali', 'Veli', 15);
INSERT INTO users(name, surname, age) VALUES('Selim', 'Ay', 12);

SELECT * FROM users;

DROP TABLE users;