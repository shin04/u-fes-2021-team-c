CREATE TABLE student_infos
(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(128),
    student_number INT
);

INSERT INTO student_infos (user_id, name, student_number) VALUES (1, 'Bob', '111111111');
INSERT INTO student_infos (user_id, name, student_number) VALUES (2, 'Tom', '222222222');
INSERT INTO student_infos (user_id, name, student_number) VALUES (3, 'Nancy', '333333333');
