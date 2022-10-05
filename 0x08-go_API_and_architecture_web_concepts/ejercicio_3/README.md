# Ejercicio 1

## Enunciated *_Listar Entidad_*

Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.

	1. Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
	2. Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
	3. Genera un handler para el endpoint llamado “GetAll”.
	4. Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 40
total 48
-rw-r--r--@ 1 juurbano  1010544492  1297 Oct  5 07:29 README.md
-rw-r--r--  1 juurbano  1010544492  1064 Oct  5 07:50 go.mod
-rw-r--r--  1 juurbano  1010544492  6962 Oct  5 07:27 go.sum
-rw-r--r--@ 1 juurbano  1010544492   976 Oct  5 07:54 main.go
-rw-r--r--  1 juurbano  1010544492   316 Oct  5 07:27 products.json
-rw-r--r--@ 1 juurbano  1010544492   304 Oct  4 09:23 main.go

$
```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

*Window 1*

```bash
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /nombre                   --> main.main.func1 (3 handlers)
[GIN-debug] GET    /products                 --> main.GetAll (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
[GIN] 2022/10/05 - 07:53:56 | 200 |     261.209µs |       127.0.0.1 | GET      "/products"
[GIN] 2022/10/05 - 07:54:01 | 200 |      51.375µs |       127.0.0.1 | GET      "/products"
[GIN] 2022/10/05 - 07:54:11 | 200 |      42.334µs |       127.0.0.1 | GET      "/products"

...

```

*Window 2*

```bash
$ curl http://127.0.0.1:8000/products -X GET -vvv ; echo "" | cat -e
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying 127.0.0.1:8000...
* Connected to 127.0.0.1 (127.0.0.1) port 8000 (#0)
> GET /products HTTP/1.1
> Host: 127.0.0.1:8000
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 05 Oct 2022 12:54:11 GMT
< Content-Length: 19
<
* Connection #0 to host 127.0.0.1 left intact
{"message":[{},{}]}$
$

```

>>> Ejemplo 1

## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

