CREATE SEQUENCE todos_colname_seq AS integer;

CREATE TABLE todos (
	id 			integer DEFAULT nextval('todos_colname_seq'::regclass),
	title 		VARCHAR(100) NOT NULL,
	isDone 		BOOLEAN NOT NULL,
	last_update DATE DEFAULT NOW()
);

ALTER SEQUENCE todos_colname_seq OWNED BY todos.id;

INSERT INTO todos(title, isDone) VALUES('Scrub your teeths', true) RETURNING *;

UPDATE todos SET title='Scrub your teeths', isDone=true WHERE id=1;

DELETE FROM todos WHERE id=1;

SELECT * FROM todos;

DROP TABLE todos;