
# DER y SQL

## Objetivo

El objetivo de esta guía práctica es poder integrar los contenidos de bases de datos relacionales vistos hasta este momento. Para ello, se propone la siguiente práctica.

> ¡Buena suerte!  


## *Escenario*

Una empresa proveedora de Internet necesita una base de datos para almacenar cada uno de sus clientes junto con el plan/pack que tiene contratado.
Mediante un análisis previo se conoce que se tiene que almacenar la siguiente información:
* De los clientes se debe registrar: dni, nombre, apellido, fecha de nacimiento, provincia, ciudad.
* En cuanto a los planes de internet: identificación del plan, velocidad ofrecida en megas, precio, descuento.

### *Ejercicio 1*

Luego del planteo de los requerimientos de la empresa, se solicita modelar los mismos mediante un DER (Diagrama Entidad-Relación).

![diagram](./images/customer_plan.png "DER diagram")

### *Ejercicio 2*

Una vez modelada y planteada la base de datos, responder a las siguientes preguntas:

a. ¿Cuál es la primary key para la tabla de clientes? Justificar respuesta

* La llave primaria considerada será el campo `dni` ya que es un número único que no cambiará como identificación del cliente

b. ¿Cuál es la primary key para la tabla de planes de internet? Justificar respuesta.

* La llave primaria considerada será el campo `plan_id` ya que es un número podria considerarse único 

c. ¿Cómo serían las relaciones entre tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar respuesta.

*_Relaciones_*

* 1 cliente solo puede contratar un servicio un plan con la compañia IT
* 1 paquete puede tenerlo muchos clientes
        
### *Ejercicio 3*

Una vez realizado el planteo del diagrama y de haber respondido estas preguntas, utilizar PHPMyAdmin o MySQL Workbench para ejecutar lo siguiente:

a. Se solicita crear una nueva base de datos llamada `empresa_internet`. 
b. Incorporar 10 registros en la tabla de clientes y 5 en la tabla de planes de internet.
c. Realizar las asociaciones/relaciones correspondientes entre estos registros.

Run the sql script `./database/company_internet.sql`

* Import your database

```bash
$ mysql -uroot -p < ./database/company_internet.sql
Enter password:
$
```

* Interative with your database

```bash
$ mysql -uroot -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 65
Server version: 8.0.31 Homebrew

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| empresa_internet   |
| information_schema |
| melisprint         |
| movies_db          |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
7 rows in set (0.01 sec)

mysql> exit
Bye
$
```

### Ejercicio 4

Plantear 10 consultas SQL que se podrían realizar a la base de datos. Expresar las sentencias.

1. Listar todos los planes de proveedor

```sql
mysql> SELECT * FROM internet_pack;
+---------+-------+-------+----------+
| plan_id | speed | price | discount |
+---------+-------+-------+----------+
|    1234 |  3.50 |  8.80 |     0.10 |
|    1235 |  3.40 |  8.70 |     0.12 |
|    1236 |  3.30 |  8.50 |     0.14 |
|    1237 |  3.20 |  8.60 |     0.10 |
|    1238 |  3.10 |  8.90 |     0.08 |
+---------+-------+-------+----------+
5 rows in set (0.01 sec)
```

2. Listar todos los clientes

```sql
mysql> SELECT * FROM customer;
+-------+------------+-----------+---------------------+------------+------------+---------+
| dni   | first_name | last_name | date_of_birth       | provincia  | ciudad     | plan_id |
+-------+------------+-----------+---------------------+------------+------------+---------+
| 34546 | Camilo     | Cuadros   | 2008-12-01 00:01:01 | California | Bogota     |    1234 |
| 34547 | Andres     | Canaviri  | 2008-10-01 00:02:01 | Colorado   | Paz        |    1235 |
| 34548 | Hernan     | Neira     | 2008-03-01 00:04:01 | Madrid     | Washington |    1236 |
| 34549 | Ricardo    | Jacobs    | 2008-04-01 00:03:01 | Cordoba    | Londres    |    1237 |
| 34551 | Joel       | Sanchez   | 2008-08-01 00:50:01 | Patagonia  | Barcelona  |    1237 |
| 34552 | Ivanna     | Cingolani | 2008-11-01 00:40:01 | Estonia    | Berlin     |    1238 |
| 34556 | Sarah      | Suarez    | 2008-05-01 00:06:01 | Paz        | Manchester |    1238 |
| 34557 | Damian     | Niz       | 2008-06-01 00:09:01 | Valley     | Paris      |    1234 |
| 34558 | Jesus      | Ortiz     | 2008-07-01 00:10:01 | LA         | Bangkok    |    1235 |
| 34559 | David      | Díaz      | 2008-09-01 00:11:01 | Quito      | Roma       |    1236 |
+-------+------------+-----------+---------------------+------------+------------+---------+
10 rows in set (0.00 sec)
```

3. Listar nombre y apellido de los clientes con `plan_id` 1234

```sql
mysql> SELECT first_name, last_name FROM customer WHERE plan_id = 1234;
+------------+-----------+
| first_name | last_name |
+------------+-----------+
| Camilo     | Cuadros   |
| Damian     | Niz       |
+------------+-----------+
2 rows in set (0.00 sec)
```

4. Listar nombres y apellidos de los clientes de la ciudad `Bogota`

```sql
mysql> SELECT first_name, last_name FROM customer WHERE ciudad = 'Bogota';
+------------+-----------+
| first_name | last_name |
+------------+-----------+
| Camilo     | Cuadros   |
+------------+-----------+
1 row in set (0.01 sec)
```

5. Mostrar el dni, nombre, apellido de los clientes ordenados por fecha de nacimiento de forma ascendente

```sql
mysql> SELECT dni, first_name, last_name, date_of_birth FROM customer ORDER BY date_of_birth ASC;
+-------+------------+-----------+---------------------+
| dni   | first_name | last_name | date_of_birth       |
+-------+------------+-----------+---------------------+
| 34548 | Hernan     | Neira     | 2008-03-01 00:04:01 |
| 34549 | Ricardo    | Jacobs    | 2008-04-01 00:03:01 |
| 34556 | Sarah      | Suarez    | 2008-05-01 00:06:01 |
| 34557 | Damian     | Niz       | 2008-06-01 00:09:01 |
| 34558 | Jesus      | Ortiz     | 2008-07-01 00:10:01 |
| 34551 | Joel       | Sanchez   | 2008-08-01 00:50:01 |
| 34559 | David      | Díaz      | 2008-09-01 00:11:01 |
| 34547 | Andres     | Canaviri  | 2008-10-01 00:02:01 |
| 34552 | Ivanna     | Cingolani | 2008-11-01 00:40:01 |
| 34546 | Camilo     | Cuadros   | 2008-12-01 00:01:01 |
+-------+------------+-----------+---------------------+
10 rows in set (0.00 sec)
```

6. Mostrar los planes donde el descuento sea igual al 10% 

```sql
mysql> SELECT plan_id, discount FROM internet_pack WHERE discount = 0.10;
+---------+----------+
| plan_id | discount |
+---------+----------+
|    1234 |     0.10 |
|    1237 |     0.10 |
+---------+----------+
2 rows in set (0.00 sec)
```

7. Mostrar el id del plan y el precio de los planes que son mas costosos top 3 

```sql
mysql> SELECT plan_id, price FROM internet_pack ORDER BY price DESC LIMIT 3;
+---------+-------+
| plan_id | price |
+---------+-------+
|    1238 |  8.90 |
|    1234 |  8.80 |
|    1235 |  8.70 |
+---------+-------+
3 rows in set (0.00 sec)
```


