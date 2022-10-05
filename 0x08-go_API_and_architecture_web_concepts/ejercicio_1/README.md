# Ejercicio 1

## Enunciated *_Red social_*

Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.

* Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
* Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
* Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.

- Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
- Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.

## Run 

- Estructura de directorios

```bash
$ ls -lR
total 24
-rw-r--r--@ 1 juurbano  1010544492  2843 Oct  4 08:29 README.md
drwxr-xr-x  3 juurbano  1010544492    96 Oct  4 08:17 go-web
-rw-r--r--@ 1 juurbano  1010544492    39 Oct  4 08:10 go.mod
drwxr-xr-x@ 4 juurbano  1010544492   128 Oct  4 08:09 internal
-rw-r--r--@ 1 juurbano  1010544492   443 Oct  4 08:20 main.go

./go-web:
total 8
-rw-r--r--  1 juurbano  1010544492  316 Oct  4 08:09 products.json

./internal:
total 0
drwxr-xr-x@ 3 juurbano  1010544492  96 Oct  4 08:17 file
drwxr-xr-x@ 3 juurbano  1010544492  96 Oct  4 08:15 service

./internal/file:
total 8
-rw-r--r--@ 1 juurbano  1010544492  470 Oct  4 08:16 file.go

./internal/service:
total 8
-rw-r--r--@ 1 juurbano  1010544492  649 Oct  4 08:15 products.go

$
```

- test cases

```json
[
    {
	"id": 1,
    	"stock": 2,
    	"name": "beer",
    	"color": "blue",
    	"code": "AC323", 
    	"publish": false,
    	"date": "December"
    },
    {
	"id": 2,
    	"stock": 4,
    	"name": "headphones",
    	"color": "red",
    	"code": "AC323T", 
    	"publish": false,
    	"date": "November"
    }
]
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

