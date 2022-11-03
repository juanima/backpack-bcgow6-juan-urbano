# Creación de tablas temporales e Indices

## *Ejercicio 2*

1. En la base de datos `movies`, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. 

* Conocemos la tabla

```sql
SELECT * FROM series;
```

* Creamos el indice al campo `title`

```sql
ALTER TABLE series ADD INDEX `idx_series_title` (`title`);
```

* Probamos que el indice funciona y hace una busqueda sobre el nombre (`hash_map`)

```sql
SELECT * FROM series WHERE title = "The Walking Dead";
```

   * Si queremos conocer el _performance_ de la query podemos y nos mostrará que se hizo la busqueda al campo title 1 row

   ```sql
   EXPLAIN SELECT * FROM series WHERE title = "The Walking Dead";
   ```

2. Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.

* Aquí se eligió la tabla `series` y como indice a la columna `title` ya que de esta manera podemos hacer una busqueda por el nombre de la pelicula


> Curiosidades

   ```sql
   SELECT * FROM series WHERE title = "The Walking Dead" AND release_date = '2011-01-01;
   ```

   * Aquí también tendremos un performance de 1 row ya que primero busca el titulo y como es único y luego sobre ese resultado seguirá con la segunda parte de la condición


> Notas

* Es mala practica tener multiples indices (llenar toda la tabla de indices)
   * el PK es el primer indice y este es el indice `cluster`
   * los demás indices ya ocupan memoria
       * construir en memoria todos esos _hash map_ podría para campo
   * tarda mas en acceder ya que primero accede al indice y luego a la tabla 


