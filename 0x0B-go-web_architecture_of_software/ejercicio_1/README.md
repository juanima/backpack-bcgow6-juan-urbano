# Ejercicio 1

## Enunciated *_Generar paquete internal_*

Se debe separar la estructura del proyecto y como primer paso generando el paquete internal, en el paquete internal deben estar todas las funcionalidades que no dependan de paquetes externos.

Dentro del paquete deben estar las capas:
1. Servicio, debe contener la lógica de nuestra aplicación.
    1. Se debe crear el archivo service.go.
    2. Se debe generar la interface Service con todos sus métodos.
    3. Se debe generar la estructura service que contenga el repositorio.
    4. Se debe generar una función que devuelva el Servicio.
    5. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
2. Repositorio, debe tener el acceso a la variable guardada en memoria.
    1. Se debe crear el archivo repository.go
    2. Se debe crear la estructura de la entidad
    3. Se deben crear las variables globales donde guardar las entidades
    4. Se debe generar la interface Repository con todos sus métodos
    5. Se debe generar la estructura repository
    6. Se debe generar una función que devuelva el Repositorio
    7. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 40
-rw-r--r--@ 1 juurbano  1010544492  1171 Oct  4 09:24 README.md
-rw-r--r--  1 juurbano  1010544492  1064 Oct  4 09:21 go.mod
-rw-r--r--  1 juurbano  1010544492  6962 Oct  4 09:21 go.sum
-rw-r--r--@ 1 juurbano  1010544492   304 Oct  4 09:23 main.go

$
```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run main.go
{Id:1 stock:0 Name:beer Color:blue Code:AC323 Price:0 Publish:false Date:December}
{Id:2 stock:0 Name:headphones Color:red Code:AC323T Price:0 Publish:false Date:November}
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

