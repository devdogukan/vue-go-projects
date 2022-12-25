CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	surname VARCHAR(30) NOT NULL,
	age INTEGER NOT NULL,
	access_permission BOOLEAN NULL
);

INSERT INTO users(name, surname, age, access_permission) 
VALUES('Dogukan', 'Aksoy', 20, true);
INSERT INTO users(name, surname, age, access_permission) 
VALUES('Mehmet', 'Solak', 19, true);
INSERT INTO users(name, surname, age, access_permission) 
VALUES('Ali', 'Veli', 15, false);
INSERT INTO users(name, surname, age, access_permission) 
VALUES('Selim', 'Ay', 12, false);

SELECT * FROM users;

DROP TABLE users;