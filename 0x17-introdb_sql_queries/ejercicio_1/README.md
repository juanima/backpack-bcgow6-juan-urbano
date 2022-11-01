# Consultas a Base de Datos


Se propone realizar las siguientes consultas a la base de datos `movies_db.sql`.

Importar el archivo `movies_db.sql` desde PHPMyAdmin o MySQL Workbench y resolver las siguientes consultas:

* Mostrar todos los registros de la tabla de movies. 
* Mostrar el nombre, apellido y rating de todos los actores.
* Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español.
* Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
* Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.
* Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
* Mostrar los títulos de las primeras tres películas en la base de datos.
* Mostrar el top 5 de las películas con mayor rating.
* Listar los primeros 10 actores.
* Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
* Mostrar a todos los actores cuyos nombres empiezan con Sam.
* Mostrar el título de las películas que salieron entre el 2004 y 2008.
* Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.



## Database on my machine

* Import database sql script

```bash
$ mysql -uroot -p < 15.\ movies-db.sql
Enter password:
$
```

* Interative with my new database

```bash
$ mysql -uroot -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 11
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
| information_schema |
| melisprint         |
| movies_db          |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
6 rows in set (0.00 sec)

mysql> USE movies_db;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> SHOW TABLES;
+---------------------+
| Tables_in_movies_db |
+---------------------+
| actor_episode       |
| actor_movie         |
| actors              |
| episodes            |
| genres              |
| migrations          |
| movies              |
| password_resets     |
| seasons             |
| series              |
| users               |
+---------------------+
11 rows in set (0.00 sec)

mysql>
```


# Queries 

1. Mostrar todos los registros de la tabla de movies. 

```sql
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
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
21 rows in set (0.00 sec)
```
---

2. Mostrar el nombre, apellido y rating de todos los actores.

```sql
mysql> SELECT first_name, last_name, rating FROM actors;
+------------+---------------------+--------+
| first_name | last_name           | rating |
+------------+---------------------+--------+
| Sam        | Worthington         |    7.5 |
| Zoe        | Saldana             |    5.5 |
| Sigourney  | Weaver              |    9.7 |
| Leonardo   | Di Caprio           |    3.5 |
| Kate       | Winslet             |    1.5 |
| Billy      | Zane                |    7.5 |
| Mark       | Hamill              |    6.5 |
| Harrison   | Ford                |    7.5 |

. . .
+------------+---------------------+--------+
49 rows in set (0.00 sec)
```
---

3. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español.

```sql
mysql> SELECT se.title as titulo FROM series se;
+---------------------+
| titulo              |
+---------------------+
| Game of Thrones     |
| Supernatural        |
| The Walking Dead    |
| Person of Interest  |
| The Big Bang Theory |
| Breaking Bad        |
+---------------------+
6 rows in set (0.00 sec)
```

---

4. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.

```sql
mysql> SELECT first_name, last_name FROM actors WHERE rating > 7.5;
+------------+-----------+
| first_name | last_name |
+------------+-----------+
| Sigourney  | Weaver    |
| Shia       | LaBeouf   |
| Sean       | Penn      |
| Renee      | Zellweger |
| Emilia     | Clarke    |
| Bryan      | Cranston  |
+------------+-----------+
6 rows in set (0.00 sec)
```

---

5. Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.

```sql
mysql> SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;
+-----------------------------------------+--------+--------+
| title                                   | rating | awards |
+-----------------------------------------+--------+--------+
| Avatar                                  |    7.9 |      3 |
| Titanic                                 |    7.7 |     11 |
| La Guerra de las galaxias: Episodio VI  |    9.1 |      7 |
| La Guerra de las galaxias: Episodio VII |    9.0 |      6 |
| Parque Jurasico                         |    8.3 |      5 |
| El rey león                             |    9.1 |      3 |
| La vida es bella                        |    8.3 |      5 |
| Carrozas de fuego                       |    9.9 |      7 |
| I am Sam                                |    9.0 |      4 |
+-----------------------------------------+--------+--------+
9 rows in set (0.00 sec)

```

---

6. Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.

```sql
mysql> SELECT title, rating FROM movies ORDER BY rating ASC;
+-----------------------------------------------------+--------+
| title                                               | rating |
+-----------------------------------------------------+--------+
| Transformers: el lado oscuro de la luna             |    0.9 |
| Toy Story 2                                         |    3.2 |
| Mi pobre angelito                                   |    3.2 |
| Harry Potter y la cámara de los secretos            |    3.5 |
| Alicia en el país de las maravillas                 |    5.7 |
| Toy Story                                           |    6.1 |
| Hotel Transylvania                                  |    7.1 |
| Big                                                 |    7.3 |
| Titanic                                             |    7.7 |
| Avatar                                              |    7.9 |
| Parque Jurasico                                     |    8.3 |
| Buscando a Nemo                                     |    8.3 |
| La vida es bella                                    |    8.3 |
| La Guerra de las galaxias: Episodio VII             |    9.0 |
| Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |
| Intensamente                                        |    9.0 |
| I am Sam                                            |    9.0 |
| La Guerra de las galaxias: Episodio VI              |    9.1 |
| El rey león                                         |    9.1 |
| Carrozas de fuego                                   |    9.9 |
| Harry Potter y la piedra filosofal                  |   10.0 |
+-----------------------------------------------------+--------+
21 rows in set (0.00 sec)
```
---

7. Mostrar los títulos de las primeras tres películas en la base de datos.

```sql
mysql> SELECT title FROM movies LIMIT 3;
+----------------------------------------+
| title                                  |
+----------------------------------------+
| Avatar                                 |
| Titanic                                |
| La Guerra de las galaxias: Episodio VI |
+----------------------------------------+
3 rows in set (0.00 sec)
```

---

8. Mostrar el top 5 de las películas con mayor rating.

```sql
mysql> SELECT title, rating FROM movies ORDER BY rating DESC LIMIT 5;
+-----------------------------------------------------+--------+
| title                                               | rating |
+-----------------------------------------------------+--------+
| Harry Potter y la piedra filosofal                  |   10.0 |
| Carrozas de fuego                                   |    9.9 |
| El rey león                                         |    9.1 |
| La Guerra de las galaxias: Episodio VI              |    9.1 |
| Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |
+-----------------------------------------------------+--------+
5 rows in set (0.01 sec)
```

---

9. Listar los primeros 10 actores.

```sql
mysql> SELECT * FROM actors LIMIT 10;
+----+------------+------------+------------+-------------+--------+-------------------+
| id | created_at | updated_at | first_name | last_name   | rating | favorite_movie_id |
+----+------------+------------+------------+-------------+--------+-------------------+
|  1 | NULL       | NULL       | Sam        | Worthington |    7.5 |                 1 |
|  2 | NULL       | NULL       | Zoe        | Saldana     |    5.5 |                 2 |
|  3 | NULL       | NULL       | Sigourney  | Weaver      |    9.7 |              NULL |
|  4 | NULL       | NULL       | Leonardo   | Di Caprio   |    3.5 |                 4 |
|  5 | NULL       | NULL       | Kate       | Winslet     |    1.5 |                 5 |
|  6 | NULL       | NULL       | Billy      | Zane        |    7.5 |                 6 |
|  7 | NULL       | NULL       | Mark       | Hamill      |    6.5 |                 7 |
|  8 | NULL       | NULL       | Harrison   | Ford        |    7.5 |                 8 |
|  9 | NULL       | NULL       | Carrie     | Fisher      |    7.5 |                 9 |
| 10 | NULL       | NULL       | Sam        | Neill       |    2.5 |                10 |
+----+------------+------------+------------+-------------+--------+-------------------+
10 rows in set (0.00 sec)
```

---

10. Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.

```sql
mysql> SELECT title, rating FROM movies WHERE title LIKE 'Toy Story%';
+-------------+--------+
| title       | rating |
+-------------+--------+
| Toy Story   |    6.1 |
| Toy Story 2 |    3.2 |
+-------------+--------+
2 rows in set (0.00 sec)
```

---

11. Mostrar a todos los actores cuyos nombres empiezan con Sam.

```sql
mysql> SELECT * FROM actors WHERE first_name LIKE 'Sam%';
+----+------------+------------+------------+-------------+--------+-------------------+
| id | created_at | updated_at | first_name | last_name   | rating | favorite_movie_id |
+----+------------+------------+------------+-------------+--------+-------------------+
|  1 | NULL       | NULL       | Sam        | Worthington |    7.5 |                 1 |
| 10 | NULL       | NULL       | Sam        | Neill       |    2.5 |                10 |
+----+------------+------------+------------+-------------+--------+-------------------+
2 rows in set (0.00 sec)
```

---

12. Mostrar el título de las películas que salieron entre el 2004 y 2008.

```sql
mysql> SELECT title, release_date FROM movies WHERE YEAR(release_date) BETWEEN 2004 AND 2008;
+-----------------------------------------------------+---------------------+
| title                                               | release_date        |
+-----------------------------------------------------+---------------------+
| La Guerra de las galaxias: Episodio VI              | 2004-07-04 00:00:00 |
| Harry Potter y las Reliquias de la Muerte - Parte 2 | 2008-07-04 00:00:00 |
| Transformers: el lado oscuro de la luna             | 2005-07-04 00:00:00 |
| Harry Potter y la piedra filosofal                  | 2008-04-04 00:00:00 |
| Alicia en el país de las maravillas                 | 2008-07-04 00:00:00 |
| Toy Story                                           | 2008-03-04 00:00:00 |
| Intensamente                                        | 2008-07-04 00:00:00 |
+-----------------------------------------------------+---------------------+
7 rows in set (0.00 sec)
```

---

13. Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.

```sql
mysql> SELECT title, rating, awards, release_date FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN 1998 AND 2009 ORDER BY rating ASC;
+-----------------------------------------------------+--------+--------+---------------------+
| title                                               | rating | awards | release_date        |
+-----------------------------------------------------+--------+--------+---------------------+
| Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |
| Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |
| Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |
| Parque Jurasico                                     |    8.3 |      5 | 1999-01-04 00:00:00 |
| Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |
| La Guerra de las galaxias: Episodio VII             |    9.0 |      6 | 2003-11-04 00:00:00 |
| Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |
| Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |
| I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |
| La Guerra de las galaxias: Episodio VI              |    9.1 |      7 | 2004-07-04 00:00:00 |
| El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |
+-----------------------------------------------------+--------+--------+---------------------+
11 rows in set (0.01 sec)
```


