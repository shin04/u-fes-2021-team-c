use team_4;

CREATE TABLE users
(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(128),
    password VARCHAR(128)
);

INSERT INTO user (name, password) VALUES ('Bob', 'pass');
INSERT INTO user (name, password) VALUES ('Tom', 'pass');
INSERT INTO user (name, password) VALUES ('Nancy', 'pass');