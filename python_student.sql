CREATE DATABASE python_student;
USE python_student;

CREATE TABLE teachers(
    teacher_id INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    password   VARCHAR(255) NOT NULL
);


CREATE TABLE students(
    student_id INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL
);
CREATE TABLE laboratories(
    lab_id   INT PRIMARY KEY AUTO_INCREMENT,
    location VARCHAR(100) NOT NULL
);
CREATE TABLE courses(
    course_id   INT PRIMARY KEY AUTO_INCREMENT,
    course_name VARCHAR(100) NOT NULL,
    teacher_id  INT          NOT NULL,
    location    VARCHAR(100) NOT NULL,
    class_time  DATETIME     NOT NULL
);
CREATE TABLE student_courses(
    student_id INT NOT NULL,
    course_id  INT NOT NULL,
    PRIMARY KEY (student_id, course_id),

);
CREATE TABLE student_teachers(
    student_id INT NOT NULL,
    teacher_id INT NOT NULL,
    PRIMARY KEY (student_id, teacher_id),
);
INSERT INTO teachers (name, password)
VALUES ('张老师', 'password123'),
       ('李老师', 'password456');

INSERT INTO laboratories (location)
VALUES ('A101'),
       ('B202');

INSERT INTO courses (course_name, teacher_id, location, class_time)
VALUES ('数学', 1, 'A101', '2024-12-21 10:00:00'),
       ('物理', 2, 'B202', '2024-12-22 14:00:00');

INSERT INTO students (name)
VALUES ('学生甲'),
       ('学生乙');

INSERT INTO student_courses (student_id, course_id)
VALUES (1, 1),
       (1, 2),
       (2, 1);

INSERT INTO student_teachers (student_id, teacher_id)
VALUES (1, 1),
       (2, 2);
