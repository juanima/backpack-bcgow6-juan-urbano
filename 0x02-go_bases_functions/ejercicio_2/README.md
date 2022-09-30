# Ejercicio 2

## Enunciated *_Impuestos de salario_*

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

## Run 

- Estructura de directorios

```bash
$ ls -l
total 40
-rw-r--r--  1 juurbano  1010544492  1063 Sep 26 18:56 README.md
-rw-r--r--  1 juurbano  1010544492    28 Sep 26 19:34 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 26 19:32 main.go
-rw-r--r--  1 juurbano  1010544492   779 Sep 26 23:25 compute_average.go
-rw-r--r--  1 juurbano  1010544492    33 Sep 26 16:04 testcase.txt
$
```

- test cases
    * la primera linea es el número de testcases
    * la segunda linea representa un delimitador entre cada test
    * la tercera linea es el input en este caso el salario

```
$ cat testcase.txt
2
-----------------------1
3
2 3 5
-----------------------2
4
-1 2 3 4

```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt
-----------------------1
El promedio calculado para el estudiante es 3.33
-----------------------2
Hubo un error
El promedio calculado para el estudiante es 0.00

$
```


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

