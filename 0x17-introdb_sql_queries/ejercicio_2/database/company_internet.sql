-- -----------------------------------------------------
-- Schema empresa_internet
-- -----------------------------------------------------
DROP DATABASE IF EXISTS empresa_internet;
CREATE DATABASE empresa_internet;
USE empresa_internet;

-- -----------------------------------------------------
-- Table internet_pack
-- -----------------------------------------------------
DROP TABLE IF EXISTS internet_pack;
CREATE TABLE internet_pack (
  plan_id int(10) NOT NULL,
  speed decimal(3, 2) NOT NULL,
  price decimal(3, 2) NOT NULL,
  discount decimal(3, 2) NOT NULL,
  PRIMARY KEY (plan_id)
);

INSERT INTO internet_pack (plan_id, speed, price, discount) VALUES 
(1234, 3.5, 8.8, 0.1),
(1235, 3.4, 8.7, 0.12),
(1236, 3.3, 8.5, 0.14),
(1237, 3.2, 8.6, 0.1),
(1238, 3.1, 8.9, 0.08)
;


-- -----------------------------------------------------
-- Table `customer`
-- -----------------------------------------------------
DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
  dni int(10) NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  date_of_birth timestamp NULL DEFAULT NULL,
  provincia varchar(100) NOT NULL,
  ciudad varchar(100) NOT NULL,
  plan_id int(10) NOT NULL,
  PRIMARY KEY (`dni`),
  CONSTRAINT customer_internet_pack_foreign FOREIGN KEY (plan_id) REFERENCES internet_pack (plan_id)
);

INSERT INTO customer (dni, first_name, last_name, date_of_birth, provincia, ciudad, plan_id) VALUES 
(34546, 'Camilo', 'Cuadros', '2008-12-01 00:01:01', 'California', 'Bogota', 1234),
(34547, 'Andres', 'Canaviri', '2008-10-01 00:02:01', 'Colorado', 'Paz', 1235),
(34548, 'Hernan', 'Neira', '2008-03-01 00:04:01', 'Madrid', 'Washington', 1236),
(34549, 'Ricardo', 'Jacobs', '2008-04-01 00:03:01', 'Cordoba', 'Londres', 1237),
(34556, 'Sarah', 'Suarez', '2008-05-01 00:06:01', 'Paz', 'Manchester', 1238),
(34557, 'Damian', 'Niz', '2008-06-01 00:09:01', 'Valley', 'Paris', 1234),
(34558, 'Jesus', 'Ortiz', '2008-07-01 00:10:01', 'LA', 'Bangkok', 1235),
(34559, 'David', 'Díaz', '2008-09-01 00:11:01', 'Quito', 'Roma', 1236),
(34551, 'Joel', 'Sanchez', '2008-08-01 00:50:01', 'Patagonia', 'Barcelona', 1237),
(34552, 'Ivanna', 'Cingolani', '2008-11-01 00:40:01', 'Estonia', 'Berlin', 1238)
;

