DROP TABLE IF EXIST posts, authors;

CREATE TABLE authors(
id SERIAL PRIMARY KEY,
name TEXT NOT NULL
);

CREATE TABLE posts(
id SERIAL PRIMARY KEY,
author_id INTEGER REFERENCES authors(id) NOT NULL,
title TEXT NOT NULL,
content TEXT NOT NULL,
created_at BIGINT NOT NULL
);

INSERT INTO authors(name) VALUES ("Дмитрий");
INSERT INTO posts(title, content, author_id) VALUES 
("Статья", "Текст статьи", 1);
