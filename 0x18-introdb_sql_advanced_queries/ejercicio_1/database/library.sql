-- -----------------------------------------------------
-- Schema library
-- -----------------------------------------------------
DROP DATABASE IF EXISTS library;
CREATE DATABASE library;
USE library;

-- -----------------------------------------------------
-- Table book
-- -----------------------------------------------------
DROP TABLE IF EXISTS book;
CREATE TABLE book (
  book_id int(10) NOT NULL,
  title varchar(100) NOT NULL,
  editorial varchar(100)  NOT NULL,
  area varchar(100) NOT NULL,
  PRIMARY KEY (book_id)
);

INSERT INTO book (book_id, title, editorial, area) VALUES 
(1234, "El buscon", "Soledad", "ficcion"),
(1235, "Rojo y negro", "Copy Editors", "narrativo"),
(1237, "En el camino", "Query Letter", "dramatico"),
(1236, "Drácula", "ProofReaders", "poetico"),
(1238, "La comedia humana", "Soledad", "novela")
;


-- -----------------------------------------------------
-- Table author
-- -----------------------------------------------------
DROP TABLE IF EXISTS author;
CREATE TABLE author (
  author_id int(10) NOT NULL,
  name varchar(100) NOT NULL,
  nationality varchar(100)  NOT NULL,
  PRIMARY KEY (author_id)
);

INSERT INTO author (author_id, name, nationality) VALUES 
(12341, "Ramiro", "Venezolana"),
(12342, "Jack", "Stroker"),
(12343, "Albert", "Gogol"),
(12344, "Emily", "Fuentes"),
(12345, "Ray", "Soledad")
;


-- -----------------------------------------------------
-- Table student
-- -----------------------------------------------------
DROP TABLE IF EXISTS student;
CREATE TABLE student (
  student_id int(10) NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  address varchar(100) NOT NULL,
  career varchar(100)  NOT NULL,
  age int(3) NOT NULL,
  PRIMARY KEY (student_id)
);

INSERT INTO student (student_id, first_name, last_name, address, career, age) VALUES 
(123411, "Ramiro", "Lozada", "Calle Santo del Villagomez No. 81", "nurse", 22),
(123422, "Jack", "Stroker", "Real del Vasconcellos No. 914", "Carnicero", 28),
(123433, "Albert", "Gogol", "Calle Bensaad No. 159", "Florista", 32),
(123444, "Emily", "Fuentes", "Cerrada Estero No. 454", "lawyer", 42),
(123455, "Ray", "Soledad", "Privada Pujolas No. 599", "mechanic", 28)
;

-- -----------------------------------------------------
-- Table `booking`
-- -----------------------------------------------------
DROP TABLE IF EXISTS booking;
CREATE TABLE booking (
  student_id int(10) NOT NULL,
  book_id int(10) NOT NULL,
  loan_date timestamp NULL DEFAULT NULL,
  return_date timestamp NULL DEFAULT NULL,
  returned boolean NOT NULL DEFAULT false,
  PRIMARY KEY (student_id, book_id),
  CONSTRAINT booking_book_foreign FOREIGN KEY (book_id) REFERENCES book(book_id),
  CONSTRAINT booking_student_foreign FOREIGN KEY (student_id) REFERENCES student(student_id)
);


INSERT INTO booking (student_id, book_id, loan_date, return_date, returned) VALUES 
(123411, 1234, '2008-12-01 00:01:01', '2008-12-23 00:01:01', false),
(123422, 1235, '2009-11-01 00:11:06', '2010-12-01 00:01:01', true),
(123433, 1236, '2002-10-02 00:12:01', '2003-12-01 00:01:01', true),
(123455, 1237, '2003-09-03 00:01:10', '2004-12-01 00:01:01', false),
(123411, 1235, '2012-01-04 00:11:01', '2013-12-01 00:01:01', false)
;

-- -----------------------------------------------------
-- Table `bookauthor`
-- -----------------------------------------------------
DROP TABLE IF EXISTS bookauthor;
CREATE TABLE bookauthor (
  author_id int(10) NOT NULL,
  book_id int(10) NOT NULL,
  PRIMARY KEY (author_id, book_id),
  CONSTRAINT bookauthor_book_foreign FOREIGN KEY (book_id) REFERENCES book(book_id),
  CONSTRAINT bookauthor_author_foreign FOREIGN KEY (author_id) REFERENCES author(author_id)
);

INSERT INTO bookauthor (author_id, book_id) VALUES 
(12341, 1234),
(12342, 1235),
(12343, 1236),
(12345, 1237),
(12342, 1234)
;

