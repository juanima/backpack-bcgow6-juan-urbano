-- -----------------------------------------------------
-- Schema company
-- -----------------------------------------------------
DROP DATABASE IF EXISTS company;
CREATE DATABASE company;
USE company;

-- -----------------------------------------------------
-- Table departament
-- -----------------------------------------------------
DROP TABLE IF EXISTS departament;
CREATE TABLE departament (
  depto_nro varchar(10) NOT NULL,
  name_depto varchar(100) NOT NULL,
  locality varchar(100) NOT NULL,
  PRIMARY KEY (depto_nro)
);

INSERT INTO departament (depto_nro, name_depto, locality) VALUES 
('D-000-1', 'Software', 'Los Tigres'),
('D-000-2', 'Sistemas', 'Guadalupe'),
('D-000-3', 'Contabilidad', 'La Roca'),
('D-000-4', 'Ventas', 'Plata')
;


-- -----------------------------------------------------
-- Table `employee`
-- -----------------------------------------------------
DROP TABLE IF EXISTS employee;
CREATE TABLE employee (
  cod_com varchar(10) NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  position varchar(100) NOT NULL,
  alta_date timestamp NULL DEFAULT NULL,
  salary int(10) NOT NULL,
  comission int(10) NOT NULL,
  depto_nro varchar(10) NOT NULL,
  PRIMARY KEY (`cod_com`),
  CONSTRAINT employee_departament_foreign FOREIGN KEY (depto_nro) REFERENCES departament (depto_nro)
);

INSERT INTO employee (cod_com, first_name, last_name, position, alta_date, salary, comission, depto_nro) VALUES 
('E-0001', 'Cesar', 'Piñero', 'Vendedor', '2018-05-12 00:40:01', 80000, 15000, 'D-000-4'),
('E-0002', 'Yosep', 'Kowaleski', 'Analista', '2015-07-14 00:40:01', 140000, 0, 'D-000-2'),
('E-0003', 'Mariela', 'Barrios', 'Director', '2014-06-05 00:40:01', 185000, 0, 'D-000-3'),
('E-0004', 'Jonathan', 'Aguilera', 'Vendedor', '2015-06-03 00:40:01', 85000, 10000, 'D-000-4'),
('E-0005', 'Daniel', 'Brezizicki', 'Vendedor', '2018-03-03 00:40:01', 83000, 10000, 'D-000-4'),
('E-0006', 'Mito', 'Barchuck', 'Presidente', '2014-06-05 00:40:01', 190000, 0, 'D-000-3'),
('E-0007', 'Emilio', 'Gazalarza', 'Desarrollador', '2014-08-02 00:40:01', 60000, 0, 'D-000-1')
;
