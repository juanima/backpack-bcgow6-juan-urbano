# Ejercicio 3

## Enunciated *_Calcular salario_*

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
  
## Run 

- Estructura de directorios

```bash
$ ls -l
total 40
-rw-r--r--  1 juurbano  1010544492  1063 Sep 26 18:56 README.md
-rw-r--r--  1 juurbano  1010544492    28 Sep 26 19:34 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 26 19:32 main.go
-rw-r--r--  1 juurbano  1010544492   779 Sep 26 23:25 compute_salary.go
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
A 10
-----------------------2
C 10
$

```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt
-----------------------1
45000
-----------------------2
10000

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

