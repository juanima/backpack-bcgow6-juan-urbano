# Ejercicio 2

## Enunciated *_Filtremos nuestro endpoint_*


Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo
se tiene que poder filtrar por todos los campos.
	1. Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
	2. Luego genera la lógica de filtrado de nuestro array.
	3. Devolver por el endpoint el array filtrado.


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 56
-rw-r--r--@ 1 juurbano  1010544492  3759 Oct  5 09:53 README.md
-rw-r--r--  1 juurbano  1010544492  1064 Oct  5 09:44 go.mod
-rw-r--r--  1 juurbano  1010544492  6962 Oct  5 09:43 go.sum
-rw-r--r--  1 juurbano  1010544492  1849 Oct  5 20:17 main.go
-rw-r--r--  1 juurbano  1010544492   142 Oct  5 11:11 products.json
-rwxr--r--  1 juurbano  1010544492   303 Oct  5 20:19 script.sh

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

[GIN-debug] GET    /filter-product           --> main.FiltrarProductsHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
2022/10/05 20:18:05 1
2022/10/05 20:18:05 12
2022/10/05 20:18:05 shirt
2022/10/05 20:18:05 blue
2022/10/05 20:18:05 AC123
2022/10/05 20:18:05 12.4
2022/10/05 20:18:05 true
2022/10/05 20:18:05 December
{ID:1 Stock:12 Name:shirt Color:blue Code:AC123 Price:12.4 Publish:true Date:December}
[0x140001accc0]
[GIN] 2022/10/05 - 20:18:05 | 200 |     599.875µs |       127.0.0.1 | GET      "/filter-product?id=1&stock=12&name=shirt&color=blue&code=AC123&price=12.4&publish=true&date=December"

...

```

*Window 2*

```bash
CO000PFW3GXDTQ9:ejercicio_1 juurbano$ cat script.sh
# Run script.sh

# {ID: 1, Stock: 12, Name: "shirt", Color: "blue", Code: "AC123", Price: 12.4, Publish: true, Date: "December"},

curl "http://127.0.0.1:8000/filter-product?id=1&stock=12&name=shirt&color=blue&code=AC123&price=12.4&publish=true&date=December" \
--request GET \
-vvv ; echo "" | cat -e

CO000PFW3GXDTQ9:ejercicio_1 juurbano$
CO000PFW3GXDTQ9:ejercicio_1 juurbano$ ./script.sh
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying 127.0.0.1:8000...
* Connected to 127.0.0.1 (127.0.0.1) port 8000 (#0)
> GET /filter-product?id=1&stock=12&name=shirt&color=blue&code=AC123&price=12.4&publish=true&date=December HTTP/1.1
> Host: 127.0.0.1:8000
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 06 Oct 2022 01:19:08 GMT
< Content-Length: 112
<
* Connection #0 to host 127.0.0.1 left intact
[{"ID":2,"Stock":123,"Name":"pants","Color":"red","Code":"AC124","Price":3.4,"Publish":false,"Date":"November"}]$
CO000PFW3GXDTQ9:ejercicio_1 juurbano$

```

_Note_: Encontramos el producto con el filtro `id=1&stock=12&name=shirt&color=blue&code=AC123&price=12.4&publish=true&date=December`


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

