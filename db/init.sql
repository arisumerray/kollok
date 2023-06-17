CREATE TABLE IF NOT EXISTS student
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    age  INTEGER     NOT NULL,
    spec VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS grade
(
    id        SERIAL PRIMARY KEY,
    studentId INTEGER REFERENCES student (id) NOT NULL,
    grade     INTEGER                         NOT NULL,
    subject   VARCHAR(64)                     NOT NULL
);