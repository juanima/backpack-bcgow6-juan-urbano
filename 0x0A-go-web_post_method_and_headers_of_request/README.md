# Ejercicio 1

## Enunciated *_Hola_*


1. Crea dentro de la carpeta go-web un archivo llamado main.go
2. Crea un servidor web con Gin que te responda un JSON que tenga una clave “message”
y diga Hola seguido por tu nombre.
3. Pegale al endpoint para corroborar que la respuesta sea la correcta.

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

