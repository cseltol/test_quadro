CREATE DATABASE IF NOT EXISTS bookshelf;

USE bookshelf;

CREATE TABLE IF NOT EXISTS books (
        id serial PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
        id serial PRIMARY KEY,
        author VARCHAR(255) NOT NULL
);

INSERT INTO books (title, author) 
VALUES 
    ('Война и мир','Толстой Л.Н.'),
    ('Анна Каренина', 'Толстой Л.Н.'),
    ('Лирика', 'Пастернак Б.Л.'),
    ('Черный человек', 'Есенин С.А.'),
    ('Белая гвардия', 'Булгаков М.А.'),
    ('Идиот', 'Достоевский Ф.М.');

INSERT INTO authors (author)
VALUES
    ('Толстой Л.Н.'),
    ('Толстой Л.Н.'),
    ('Пастернак Б.Л.'),
    ('Есенин С.А.'),
    ('Булгаков М.А.'),
    ('Достоевский Ф.М.');