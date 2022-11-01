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



