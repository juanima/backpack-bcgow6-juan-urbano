# Advanced SQL queries

## *Practica grupal*

*_Resolver las siguientes consignas_*

Tomando la base de datos `movies_db.sql`, se solicita:


1. Agregar una película a la tabla movies.

```sql
mysql> INSERT INTO movies VALUES(DEFAULT, NULL, NULL, "NUEVA PELI", 7.9, 3, NOW(), 120, NULL);
Query OK, 1 row affected (0.01 sec)

mysql> SELECT * FROM movies;
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
| id | created_at | updated_at | title                                               | rating | awards | release_date        | length | genre_id |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
|  1 | NULL       | NULL       | Avatar                                              |    7.9 |      3 | 2010-10-04 00:00:00 |    120 |        5 |
|  2 | NULL       | NULL       | Titanic                                             |    7.7 |     11 | 1997-09-04 00:00:00 |    320 |        3 |
|  3 | NULL       | NULL       | La Guerra de las galaxias: Episodio VI              |    9.1 |      7 | 2004-07-04 00:00:00 |   NULL |        5 |
|  4 | NULL       | NULL       | La Guerra de las galaxias: Episodio VII             |    9.0 |      6 | 2003-11-04 00:00:00 |    180 |        5 |
|  5 | NULL       | NULL       | Parque Jurasico                                     |    8.3 |      5 | 1999-01-04 00:00:00 |    270 |        5 |
|  6 | NULL       | NULL       | Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |    190 |        6 |
|  7 | NULL       | NULL       | Transformers: el lado oscuro de la luna             |    0.9 |      1 | 2005-07-04 00:00:00 |   NULL |        5 |
|  8 | NULL       | NULL       | Harry Potter y la piedra filosofal                  |   10.0 |      1 | 2008-04-04 00:00:00 |    120 |        8 |
|  9 | NULL       | NULL       | Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |    200 |        8 |
| 10 | NULL       | NULL       | El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |   NULL |       10 |
| 11 | NULL       | NULL       | Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |    120 |     NULL |
| 12 | NULL       | NULL       | Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |    110 |        7 |
| 13 | NULL       | NULL       | Toy Story                                           |    6.1 |      0 | 2008-03-04 00:00:00 |    150 |        7 |
| 14 | NULL       | NULL       | Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |    120 |        7 |
| 15 | NULL       | NULL       | La vida es bella                                    |    8.3 |      5 | 1994-10-04 00:00:00 |   NULL |        3 |
| 16 | NULL       | NULL       | Mi pobre angelito                                   |    3.2 |      1 | 1989-01-04 00:00:00 |    120 |        1 |
| 17 | NULL       | NULL       | Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |    120 |        7 |
| 18 | NULL       | NULL       | Carrozas de fuego                                   |    9.9 |      7 | 1980-07-04 00:00:00 |    180 |     NULL |
| 19 | NULL       | NULL       | Big                                                 |    7.3 |      2 | 1988-02-04 00:00:00 |    130 |        8 |
| 20 | NULL       | NULL       | I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |    130 |        3 |
| 21 | NULL       | NULL       | Hotel Transylvania                                  |    7.1 |      1 | 2012-05-04 00:00:00 |     90 |       10 |
| 23 | NULL       | NULL       | NUEVA PELI                                          |    7.9 |      3 | 2022-11-06 15:31:16 |    120 |     NULL |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
22 rows in set (0.00 sec)
```

2. Agregar un género a la tabla genres.

```sql
mysql> INSERT INTO genres VALUES(DEFAULT, NULL, NULL, "NUEVO GENERO", 23, 1);
Query OK, 1 row affected (0.00 sec)

mysql> SELECT * FROM genres;
+----+---------------------+------------+-----------------+---------+--------+
| id | created_at          | updated_at | name            | ranking | active |
+----+---------------------+------------+-----------------+---------+--------+
|  1 | 2016-07-03 22:00:00 | NULL       | Comedia         |       1 |      1 |
|  2 | 2014-07-03 22:00:00 | NULL       | Terror          |       2 |      1 |
|  3 | 2013-07-03 22:00:00 | NULL       | Drama           |       3 |      1 |
|  4 | 2011-07-03 22:00:00 | NULL       | Accion          |       4 |      1 |
|  5 | 2010-07-03 22:00:00 | NULL       | Ciencia Ficcion |       5 |      1 |
|  6 | 2013-07-03 22:00:00 | NULL       | Suspenso        |       6 |      1 |
|  7 | 2005-07-03 22:00:00 | NULL       | Animacion       |       7 |      1 |
|  8 | 2003-07-03 22:00:00 | NULL       | Aventuras       |       8 |      1 |
|  9 | 2008-07-03 22:00:00 | NULL       | Documental      |       9 |      1 |
| 10 | 2013-07-03 22:00:00 | NULL       | Infantiles      |      10 |      1 |
| 11 | 2011-07-03 22:00:00 | NULL       | Fantasia        |      11 |      1 |
| 12 | 2013-07-03 22:00:00 | NULL       | Musical         |      12 |      1 |
| 13 | NULL                | NULL       | NUEVO GENERO    |      23 |      1 |
+----+---------------------+------------+-----------------+---------+--------+
13 rows in set (0.00 sec)
```

3. Asociar a la película del Ej 2. con el género creado en el Ej. 3.

```sql 
mysql> UPDATE movies SET genre_id = 13 WHERE id = 23; -- genre_id es el id del nuevo genero agregado y id 23 es la última pelicula creada
Query OK, 1 row affected (0.00 sec)
Rows matched: 1  Changed: 1  Warnings: 0

mysql> SELECT * FROM movies;
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
| id | created_at | updated_at | title                                               | rating | awards | release_date        | length | genre_id |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
|  1 | NULL       | NULL       | Avatar                                              |    7.9 |      3 | 2010-10-04 00:00:00 |    120 |        5 |
|  2 | NULL       | NULL       | Titanic                                             |    7.7 |     11 | 1997-09-04 00:00:00 |    320 |        3 |
|  3 | NULL       | NULL       | La Guerra de las galaxias: Episodio VI              |    9.1 |      7 | 2004-07-04 00:00:00 |   NULL |        5 |
|  4 | NULL       | NULL       | La Guerra de las galaxias: Episodio VII             |    9.0 |      6 | 2003-11-04 00:00:00 |    180 |        5 |
|  5 | NULL       | NULL       | Parque Jurasico                                     |    8.3 |      5 | 1999-01-04 00:00:00 |    270 |        5 |
|  6 | NULL       | NULL       | Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |    190 |        6 |
|  7 | NULL       | NULL       | Transformers: el lado oscuro de la luna             |    0.9 |      1 | 2005-07-04 00:00:00 |   NULL |        5 |
|  8 | NULL       | NULL       | Harry Potter y la piedra filosofal                  |   10.0 |      1 | 2008-04-04 00:00:00 |    120 |        8 |
|  9 | NULL       | NULL       | Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |    200 |        8 |
| 10 | NULL       | NULL       | El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |   NULL |       10 |
| 11 | NULL       | NULL       | Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |    120 |     NULL |
| 12 | NULL       | NULL       | Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |    110 |        7 |
| 13 | NULL       | NULL       | Toy Story                                           |    6.1 |      0 | 2008-03-04 00:00:00 |    150 |        7 |
| 14 | NULL       | NULL       | Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |    120 |        7 |
| 15 | NULL       | NULL       | La vida es bella                                    |    8.3 |      5 | 1994-10-04 00:00:00 |   NULL |        3 |
| 16 | NULL       | NULL       | Mi pobre angelito                                   |    3.2 |      1 | 1989-01-04 00:00:00 |    120 |        1 |
| 17 | NULL       | NULL       | Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |    120 |        7 |
| 18 | NULL       | NULL       | Carrozas de fuego                                   |    9.9 |      7 | 1980-07-04 00:00:00 |    180 |     NULL |
| 19 | NULL       | NULL       | Big                                                 |    7.3 |      2 | 1988-02-04 00:00:00 |    130 |        8 |
| 20 | NULL       | NULL       | I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |    130 |        3 |
| 21 | NULL       | NULL       | Hotel Transylvania                                  |    7.1 |      1 | 2012-05-04 00:00:00 |     90 |       10 |
| 23 | NULL       | NULL       | NUEVA PELI                                          |    7.9 |      3 | 2022-11-06 15:31:16 |    120 |       13 |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
22 rows in set (0.00 sec)
```

4. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

```sql
mysql> SELECT * FROM actors;
+----+------------+------------+------------+---------------------+--------+-------------------+
| id | created_at | updated_at | first_name | last_name           | rating | favorite_movie_id |
+----+------------+------------+------------+---------------------+--------+-------------------+
|  1 | NULL       | NULL       | Sam        | Worthington         |    7.5 |                 1 |
|  2 | NULL       | NULL       | Zoe        | Saldana             |    5.5 |                 2 |
|  3 | NULL       | NULL       | Sigourney  | Weaver              |    9.7 |              NULL |
|  4 | NULL       | NULL       | Leonardo   | Di Caprio           |    3.5 |                 4 |

. . .

| 45 | NULL       | NULL       | Jim        | Parsons             |    6.9 |                 3 |
| 46 | NULL       | NULL       | Kaley      | Cuoco               |    2.3 |                 4 |
| 47 | NULL       | NULL       | Bryan      | Cranston            |    7.9 |              NULL |
| 48 | NULL       | NULL       | Aaron      | Paul                |    5.9 |                 6 |
| 49 | NULL       | NULL       | Anna       | Gunn                |    3.1 |                 7 |
+----+------------+------------+------------+---------------------+--------+-------------------+
49 rows in set (0.01 sec)

mysql>
mysql> UPDATE actors SET favorite_movie_id = 23 WHERE id=47;
Query OK, 1 row affected (0.01 sec)
Rows matched: 1  Changed: 1  Warnings: 0

mysql> SELECT * FROM actors;
+----+------------+------------+------------+---------------------+--------+-------------------+
| id | created_at | updated_at | first_name | last_name           | rating | favorite_movie_id |
+----+------------+------------+------------+---------------------+--------+-------------------+
|  1 | NULL       | NULL       | Sam        | Worthington         |    7.5 |                 1 |
|  2 | NULL       | NULL       | Zoe        | Saldana             |    5.5 |                 2 |
|  3 | NULL       | NULL       | Sigourney  | Weaver              |    9.7 |              NULL |
|  4 | NULL       | NULL       | Leonardo   | Di Caprio           |    3.5 |                 4 |

. . .

| 45 | NULL       | NULL       | Jim        | Parsons             |    6.9 |                 3 |
| 46 | NULL       | NULL       | Kaley      | Cuoco               |    2.3 |                 4 |
| 47 | NULL       | NULL       | Bryan      | Cranston            |    7.9 |                23 |
| 48 | NULL       | NULL       | Aaron      | Paul                |    5.9 |                 6 |
| 49 | NULL       | NULL       | Anna       | Gunn                |    3.1 |                 7 |
+----+------------+------------+------------+---------------------+--------+-------------------+
49 rows in set (0.01 sec)

mysql>
mysql> SELECT a.first_name, a.last_name, m.title FROM actors AS a INNER JOIN movies AS m ON a.favorite_movie_id = m.id WHERE a.favorite_movie_id = 23;
+------------+-----------+------------+
| first_name | last_name | title      |
+------------+-----------+------------+
| Bryan      | Cranston  | NUEVA PELI |
+------------+-----------+------------+
1 row in set (0.00 sec)
```

---

5. Crear una tabla temporal copia de la tabla movies.

```sql
mysql> CREATE TEMPORARY TABLE movies_temporary AS SELECT * FROM movies;
mysql> 
mysql> SHOW CREATE TABLE movies_temporary \G;
*************************** 1. row ***************************
       Table: movies_temporary
Create Table: CREATE TEMPORARY TABLE `movies_temporary` (
  `id` int unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `title` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
  `rating` decimal(3,1) unsigned NOT NULL,
  `awards` int unsigned NOT NULL DEFAULT '0',
  `release_date` datetime NOT NULL,
  `length` int unsigned DEFAULT NULL,
  `genre_id` int unsigned DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)

ERROR:
No query specified
mysql>
mysql>
mysql> SELECT * FROM movies_temporary;
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
| id | created_at | updated_at | title                                               | rating | awards | release_date        | length | genre_id |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
|  1 | NULL       | NULL       | Avatar                                              |    7.9 |      3 | 2010-10-04 00:00:00 |    120 |        5 |
|  2 | NULL       | NULL       | Titanic                                             |    7.7 |     11 | 1997-09-04 00:00:00 |    320 |        3 |
|  3 | NULL       | NULL       | La Guerra de las galaxias: Episodio VI              |    9.1 |      7 | 2004-07-04 00:00:00 |   NULL |        5 |
|  4 | NULL       | NULL       | La Guerra de las galaxias: Episodio VII             |    9.0 |      6 | 2003-11-04 00:00:00 |    180 |        5 |
|  5 | NULL       | NULL       | Parque Jurasico                                     |    8.3 |      5 | 1999-01-04 00:00:00 |    270 |        5 |
|  6 | NULL       | NULL       | Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |    190 |        6 |
|  7 | NULL       | NULL       | Transformers: el lado oscuro de la luna             |    0.9 |      1 | 2005-07-04 00:00:00 |   NULL |        5 |
|  8 | NULL       | NULL       | Harry Potter y la piedra filosofal                  |   10.0 |      1 | 2008-04-04 00:00:00 |    120 |        8 |
|  9 | NULL       | NULL       | Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |    200 |        8 |
| 10 | NULL       | NULL       | El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |   NULL |       10 |
| 11 | NULL       | NULL       | Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |    120 |     NULL |
| 12 | NULL       | NULL       | Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |    110 |        7 |
| 13 | NULL       | NULL       | Toy Story                                           |    6.1 |      0 | 2008-03-04 00:00:00 |    150 |        7 |
| 14 | NULL       | NULL       | Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |    120 |        7 |
| 15 | NULL       | NULL       | La vida es bella                                    |    8.3 |      5 | 1994-10-04 00:00:00 |   NULL |        3 |
| 16 | NULL       | NULL       | Mi pobre angelito                                   |    3.2 |      1 | 1989-01-04 00:00:00 |    120 |        1 |
| 17 | NULL       | NULL       | Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |    120 |        7 |
| 18 | NULL       | NULL       | Carrozas de fuego                                   |    9.9 |      7 | 1980-07-04 00:00:00 |    180 |     NULL |
| 19 | NULL       | NULL       | Big                                                 |    7.3 |      2 | 1988-02-04 00:00:00 |    130 |        8 |
| 20 | NULL       | NULL       | I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |    130 |        3 |
| 21 | NULL       | NULL       | Hotel Transylvania                                  |    7.1 |      1 | 2012-05-04 00:00:00 |     90 |       10 |
| 23 | NULL       | NULL       | NUEVA PELI                                          |    7.9 |      3 | 2022-11-06 15:31:16 |    120 |       13 |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
22 rows in set (0.00 sec)

mysql>

```

* alternativa 2) crear una tabla temporal

```sql
-- or 
CREATE TEMPORARY TABLE IF NOT EXISTS movies_temporary(
    id int not null primary key auto_increment,
    created_at timestamp not null,
    updated_at timestamp not null,
    title varchar(255) not null, 
    rating decimal(3,1) not null, 
    awards int not null,
    release_date datetime not null, 
    length int not null,
    genre_id int not null
);
```

* Insertar valores en la tabla temporal

```sql
INSERT INTO movies_temporary VALUES(
	DEFAULT,
    current_timestamp(),
    current_timestamp(),
    "titulo",
    5.9,
    4,
    NOW(),
    4,
    2
)
```

6. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

```sql
mysql> SELECT * FROM movies_temporary WHERE awards < 5;
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
| id | created_at | updated_at | title                                               | rating | awards | release_date        | length | genre_id |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
|  1 | NULL       | NULL       | Avatar                                              |    7.9 |      3 | 2010-10-04 00:00:00 |    120 |        5 |
|  6 | NULL       | NULL       | Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |    190 |        6 |
|  7 | NULL       | NULL       | Transformers: el lado oscuro de la luna             |    0.9 |      1 | 2005-07-04 00:00:00 |   NULL |        5 |
|  8 | NULL       | NULL       | Harry Potter y la piedra filosofal                  |   10.0 |      1 | 2008-04-04 00:00:00 |    120 |        8 |
|  9 | NULL       | NULL       | Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |    200 |        8 |
| 10 | NULL       | NULL       | El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |   NULL |       10 |
| 11 | NULL       | NULL       | Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |    120 |     NULL |
| 12 | NULL       | NULL       | Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |    110 |        7 |
| 13 | NULL       | NULL       | Toy Story                                           |    6.1 |      0 | 2008-03-04 00:00:00 |    150 |        7 |
| 14 | NULL       | NULL       | Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |    120 |        7 |
| 16 | NULL       | NULL       | Mi pobre angelito                                   |    3.2 |      1 | 1989-01-04 00:00:00 |    120 |        1 |
| 17 | NULL       | NULL       | Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |    120 |        7 |
| 19 | NULL       | NULL       | Big                                                 |    7.3 |      2 | 1988-02-04 00:00:00 |    130 |        8 |
| 20 | NULL       | NULL       | I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |    130 |        3 |
| 21 | NULL       | NULL       | Hotel Transylvania                                  |    7.1 |      1 | 2012-05-04 00:00:00 |     90 |       10 |
| 23 | NULL       | NULL       | NUEVA PELI                                          |    7.9 |      3 | 2022-11-06 15:31:16 |    120 |       13 |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
16 rows in set (0.00 sec)

mysql>
mysql> DELETE FROM movies_temporary WHERE awards < 5;
Query OK, 16 rows affected (0.00 sec)
```

7. Obtener la lista de todos los géneros que tengan al menos una película.

```sql 
mysql> SELECT * FROM movies m LEFT JOIN genres g ON m.genre_id = g.id;
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+------+---------------------+------------+-----------------+---------+--------+
| id | created_at | updated_at | title                                               | rating | awards | release_date        | length | genre_id | id   | created_at          | updated_at | name            | ranking | active |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+------+---------------------+------------+-----------------+---------+--------+
|  1 | NULL       | NULL       | Avatar                                              |    7.9 |      3 | 2010-10-04 00:00:00 |    120 |        5 |    5 | 2010-07-03 22:00:00 | NULL       | Ciencia Ficcion |       5 |      1 |
|  2 | NULL       | NULL       | Titanic                                             |    7.7 |     11 | 1997-09-04 00:00:00 |    320 |        3 |    3 | 2013-07-03 22:00:00 | NULL       | Drama           |       3 |      1 |
|  3 | NULL       | NULL       | La Guerra de las galaxias: Episodio VI              |    9.1 |      7 | 2004-07-04 00:00:00 |   NULL |        5 |    5 | 2010-07-03 22:00:00 | NULL       | Ciencia Ficcion |       5 |      1 |
|  4 | NULL       | NULL       | La Guerra de las galaxias: Episodio VII             |    9.0 |      6 | 2003-11-04 00:00:00 |    180 |        5 |    5 | 2010-07-03 22:00:00 | NULL       | Ciencia Ficcion |       5 |      1 |
|  5 | NULL       | NULL       | Parque Jurasico                                     |    8.3 |      5 | 1999-01-04 00:00:00 |    270 |        5 |    5 | 2010-07-03 22:00:00 | NULL       | Ciencia Ficcion |       5 |      1 |
|  6 | NULL       | NULL       | Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |    190 |        6 |    6 | 2013-07-03 22:00:00 | NULL       | Suspenso        |       6 |      1 |
|  7 | NULL       | NULL       | Transformers: el lado oscuro de la luna             |    0.9 |      1 | 2005-07-04 00:00:00 |   NULL |        5 |    5 | 2010-07-03 22:00:00 | NULL       | Ciencia Ficcion |       5 |      1 |
|  8 | NULL       | NULL       | Harry Potter y la piedra filosofal                  |   10.0 |      1 | 2008-04-04 00:00:00 |    120 |        8 |    8 | 2003-07-03 22:00:00 | NULL       | Aventuras       |       8 |      1 |
|  9 | NULL       | NULL       | Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |    200 |        8 |    8 | 2003-07-03 22:00:00 | NULL       | Aventuras       |       8 |      1 |
| 10 | NULL       | NULL       | El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |   NULL |       10 |   10 | 2013-07-03 22:00:00 | NULL       | Infantiles      |      10 |      1 |
| 11 | NULL       | NULL       | Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |    120 |     NULL | NULL | NULL                | NULL       | NULL            |    NULL |   NULL |
| 12 | NULL       | NULL       | Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |    110 |        7 |    7 | 2005-07-03 22:00:00 | NULL       | Animacion       |       7 |      1 |
| 13 | NULL       | NULL       | Toy Story                                           |    6.1 |      0 | 2008-03-04 00:00:00 |    150 |        7 |    7 | 2005-07-03 22:00:00 | NULL       | Animacion       |       7 |      1 |
| 14 | NULL       | NULL       | Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |    120 |        7 |    7 | 2005-07-03 22:00:00 | NULL       | Animacion       |       7 |      1 |
| 15 | NULL       | NULL       | La vida es bella                                    |    8.3 |      5 | 1994-10-04 00:00:00 |   NULL |        3 |    3 | 2013-07-03 22:00:00 | NULL       | Drama           |       3 |      1 |
| 16 | NULL       | NULL       | Mi pobre angelito                                   |    3.2 |      1 | 1989-01-04 00:00:00 |    120 |        1 |    1 | 2016-07-03 22:00:00 | NULL       | Comedia         |       1 |      1 |
| 17 | NULL       | NULL       | Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |    120 |        7 |    7 | 2005-07-03 22:00:00 | NULL       | Animacion       |       7 |      1 |
| 18 | NULL       | NULL       | Carrozas de fuego                                   |    9.9 |      7 | 1980-07-04 00:00:00 |    180 |     NULL | NULL | NULL                | NULL       | NULL            |    NULL |   NULL |
| 19 | NULL       | NULL       | Big                                                 |    7.3 |      2 | 1988-02-04 00:00:00 |    130 |        8 |    8 | 2003-07-03 22:00:00 | NULL       | Aventuras       |       8 |      1 |
| 20 | NULL       | NULL       | I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |    130 |        3 |    3 | 2013-07-03 22:00:00 | NULL       | Drama           |       3 |      1 |
| 21 | NULL       | NULL       | Hotel Transylvania                                  |    7.1 |      1 | 2012-05-04 00:00:00 |     90 |       10 |   10 | 2013-07-03 22:00:00 | NULL       | Infantiles      |      10 |      1 |
| 23 | NULL       | NULL       | NUEVA PELI                                          |    7.9 |      3 | 2022-11-06 15:31:16 |    120 |       13 |   13 | NULL                | NULL       | NUEVO GENERO    |      23 |      1 |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+------+---------------------+------------+-----------------+---------+--------+
22 rows in set (0.01 sec)

mysql>
mysql> SELECT  g.name  FROM movies m LEFT JOIN genres g ON m.genre_id = g.id GROUP BY g.name HAVING COUNT(m.id) >= 1;
+-----------------+
| name            |
+-----------------+
| NULL            |
| Comedia         |
| Drama           |
| Ciencia Ficcion |
| Suspenso        |
| Animacion       |
| Aventuras       |
| Infantiles      |
| NUEVO GENERO    |
+-----------------+
9 rows in set (0.00 sec)

```
---

8. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 

```sql
SELECT a.* FROM actors AS a INNER JOIN movies AS m ON m.id = a.favorite_movie_id WHERE m.awards > 3;
```

9. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

```sql
mysql> EXPLAIN INSERT INTO movies_temporary VALUES(DEFAULT, NULL, NULL, "titulo", 5.9, 4, NOW(), 4, 2);
+----+-------------+------------------+------------+------+---------------+------+---------+------+------+----------+-------+
| id | select_type | table            | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra |
+----+-------------+------------------+------------+------+---------------+------+---------+------+------+----------+-------+
|  1 | INSERT      | movies_temporary | NULL       | ALL  | NULL          | NULL | NULL    | NULL | NULL |     NULL | NULL  |
+----+-------------+------------------+------------+------+---------------+------+---------+------+------+----------+-------+
1 row in set, 1 warning (0.00 sec)
```

```sql
mysql> EXPLAIN DELETE FROM movies_temporary WHERE awards < 5;
+----+-------------+------------------+------------+------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table            | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+------------------+------------+------+---------------+------+---------+------+------+----------+-------------+
|  1 | DELETE      | movies_temporary | NULL       | ALL  | NULL          | NULL | NULL    | NULL |   21 |   100.00 | Using where |
+----+-------------+------------------+------------+------+---------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.01 sec)
```

10. ¿Qué son los índices? ¿Para qué sirven?

```
Un índice contiene claves generadas a partir de una o varias columnas de la tabla o la vista. Dichas claves están almacenadas en una estructura (árbol b) que permite que el motor de base de datos SQL busque de forma rápida y eficiente la fila o filas asociadas a los valores de cada clave.
```

11. Crear un índice sobre el nombre en la tabla movies.

```sql
CREATE INDEX title_idx ON movies_temporary (title);
CREATE INDEX title_idx ON movies (title);
```

12. Chequee que el índice fue creado correctamente.

```sql
SHOW INDEX FROM movies_temporary;
SHOW INDEX FROM movies;
```

