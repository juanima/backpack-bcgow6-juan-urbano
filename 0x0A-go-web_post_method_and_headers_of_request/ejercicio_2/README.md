# Ejercicio 2

## Enunciated *_Crear Entidad_*

Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los siguientes pasos:

1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro nuevo registro (sin tener una variable de último ID a nivel global).


## Enunciated *_Validación de campos_*

Se debe implementar las validaciones de los campos al momento de enviar la petición, para eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400 con el mensaje “el campo %s es requerido”. (En %s debe ir el nombre del campo que no está completo).


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 48
-rw-r--r--@ 1 juurbano  1010544492  3497 Oct  4 23:54 README.md
-rw-r--r--  1 juurbano  1010544492  1064 Oct  4 23:21 go.mod
-rw-r--r--  1 juurbano  1010544492  6962 Oct  4 23:21 go.sum
-rw-r--r--  1 juurbano  1010544492  1673 Oct  4 23:51 main.go
-rw-r--r--  1 juurbano  1010544492   138 Oct  4 23:50 products.json

$
```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

*Window 1*

```bash
CO000PFW3GXDTQ9:ejercicio_2 juurbano$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /products/                --> main.saveProduct.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2022/10/04 - 23:50:27 | 400 |         960µs |       127.0.0.1 | POST     "/products/"

...

```

*Window 2*

```bash
$ cat products.json
{
    "id": 1,
    "stock": 2,
    "name": "beer",
    "price": 3.4,
    "color": "blue",
    "code": "AC323",
    "date": "December"
}

$ curl http://127.0.0.1:8080/products/ -X POST -H "Content-Type: application/json" -d @products.json -vvv ; echo "" | cat -e
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /products/ HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 128
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 400 Bad Request
< Content-Type: application/json; charset=utf-8
< Date: Wed, 05 Oct 2022 04:50:27 GMT
< Content-Length: 35
<
* Connection #0 to host 127.0.0.1 left intact
{"error":"El publish es requerido"}$
$

```

_Note_: Por ejemplo aquí nos hace falta el campo `publish` del json, nos debería retornar el error


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

