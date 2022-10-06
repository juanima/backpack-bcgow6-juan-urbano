# Ejercicio 2

## Enunciated *_Get one endpoint_*

Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática. Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.

1. Genera una nueva ruta.
2. Genera un handler para la ruta creada.
3. Dentro del handler busca el item que necesitas.
4. Devuelve el item según el id.

Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 56
-rw-r--r--@ 1 juurbano  1010544492  3759 Oct  5 09:53 README.md
-rw-r--r--  1 juurbano  1010544492  1064 Oct  5 09:44 go.mod
-rw-r--r--  1 juurbano  1010544492  6962 Oct  5 09:43 go.sum
-rw-r--r--  1 juurbano  1010544492  1849 Oct  5 20:17 main.go

$
```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

*Window 1*

```bash
$ go run main.go
CO000PFW3GXDTQ9:ejercicio_2 juurbano$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /products/:id             --> main.ProductHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
[GIN] 2022/10/05 - 21:15:26 | 200 |      36.708µs |       127.0.0.1 | GET      "/products/2"

...

```

*Window 2*

```bash
$ curl http://127.0.0.1:8000/products/2 --request GET -vvv ; echo "" | cat -e
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying 127.0.0.1:8000...
* Connected to 127.0.0.1 (127.0.0.1) port 8000 (#0)
> GET /products/2 HTTP/1.1
> Host: 127.0.0.1:8000
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 06 Oct 2022 02:15:26 GMT
< Content-Length: 110
<
* Connection #0 to host 127.0.0.1 left intact
{"ID":2,"Stock":123,"Name":"pants","Color":"red","Code":"AC124","Price":3.4,"Publish":false,"Date":"November"}$
$

```

_Note_: Encontramos el producto con el `id=2`


---


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

