# Ejercicio 1

## Enunciated *_Registro de estudiantes_*

Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle


## Run 

- Estructura de directorios

```bash
$ ls -l
total 40
-rw-r--r--  1 juurbano  1010544492  1063 Sep 26 18:56 README.md
-rw-r--r--  1 juurbano  1010544492    28 Sep 26 19:34 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 26 19:32 main.go
-rw-r--r--  1 juurbano  1010544492   779 Sep 26 23:25 register_students.go
-rw-r--r--  1 juurbano  1010544492    33 Sep 26 16:04 testcase.txt
$
```

- test cases
    * la primera linea es el número de testcases
    * la segunda linea es el número de estudiantes para registrar
    * las siguientes líneas correspondes a los datos de los estudiantes
	* `Name` `LastName` `DNI` `date`

```
$ cat testcase.txt
2
-----------------------1
2
Pedro Jose 1038498115 15/09/2022
Maria Hernandez 1038468115 16/10/2022
-----------------------2
1
Pedro Jose 1038498115 15/09/2022

```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt
-----------------------1
Guau!
Info del estudiante {"name":"Pedro","lastname":"Jose","dni":1038498115,"date":"15/09/2022"}
Guau!
Info del estudiante {"name":"Maria","lastname":"Hernandez","dni":1038468115,"date":"16/10/2022"}
-----------------------2
Guau!
Info del estudiante {"name":"Pedro","lastname":"Jose","dni":1038498115,"date":"15/09/2022"}

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

