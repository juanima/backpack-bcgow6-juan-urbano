# Ejercicio 1

## Enunciated *_Red social_*

Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.

La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:

* Cambiar nombre: me permite cambiar el nombre y apellido.
* Cambiar edad: me permite cambiar la edad.
* Cambiar correo: me permite cambiar el correo.
* Cambiar contraseña: me permite cambiar la contraseña.

## Run 

- Estructura de directorios

```bash
$ ls -lR
total 40
-rw-r--r--  1 juurbano  1010544492  1844 Sep 27 16:41 README.md
-rw-r--r--  1 juurbano  1010544492    47 Sep 27 22:23 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 27 17:01 main.go
drwxr-xr-x  3 juurbano  1010544492    96 Sep 27 23:42 models
-rw-r--r--  1 juurbano  1010544492   867 Sep 27 23:56 social_network.go
-rw-r--r--  1 juurbano  1010544492   232 Sep 27 23:20 testcase.txt

./models:
total 8
-rw-r--r--  1 juurbano  1010544492  920 Sep 27 23:42 User.go

$
```

- test cases
    * la primera linea es el número de testcases
    * la segunda linea es el número de users para editar 
    * las siguientes líneas correspondes a los datos de los usuarios 
 `nombre` `apellido` `email` `contraseña` `edad`

``
$ cat testcase.txt
1
-----------------------1
3
Nicolas Salva nico29@gmail.com hola1234 29 
Torre Pilar salv29@gmail.com hola1235 30
Andres Hernandez andres9@gmail.com hola1236 35
-----------------------2
1
Nicolas Salva 29 nico29@gmail.com hola1234

```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt
-----------------------1
0x1400010c140
Info User:
        Name: Nicolas,
        Last name: Salva,
        Email: nico29@gmail.com,
        Age: 29
Info User:
        Name: Nicolas,
        Last name: Salva,
        Email: patrick@gmail.com,
        Age: 29
0x1400010c190
Info User:
        Name: Torre,
        Last name: Pilar,
        Email: salv29@gmail.com,
        Age: 30
Info User:
        Name: Torre,
        Last name: Pilar,
        Email: patrick@gmail.com,
        Age: 30
0x1400010c1e0
Info User:
        Name: Andres,
        Last name: Hernandez,
        Email: andres9@gmail.com,
        Age: 35
Info User:
        Name: Andres,
        Last name: Hernandez,
        Email: patrick@gmail.com,
        Age: 35

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

