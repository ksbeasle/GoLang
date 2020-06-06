CREATE TABLE games (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title TEXT NOT NULL,
genre TEXT NOT NULL,
rating INTEGER NOT NULL,
platform TEXT NOT NULL,
releaseDate TEXT NOT NULL
);



INSERT INTO games (title, genre, rating, platform, releaseDate) VALUES (
'Game Title', 'Game Genre', 1, 'Game Platform', 'January 1, 2000');