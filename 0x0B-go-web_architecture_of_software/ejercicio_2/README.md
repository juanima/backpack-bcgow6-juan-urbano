# Ejercicio 2

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


## Enunciated *_Generar paquete server_*

Se debe separar la estructura del proyecto, como segundo paso se debe generar el paquete server donde se agregaran las funcionalidades del proyecto que dependan de paquetes externos y el main del programa.

1. Dentro del paquete deben estar:
	1. El main del programa.
	2. Se debe importar e inyectar el repositorio, servicio y handler
2. Se debe implementar el router para los diferentes endpoints
	1. El paquete handler con el controlador de la entidad seleccionada.
	2. Se debe generar la estructura request
	3. Se debe generar la estructura del controlador que tenga como campo el servicio
	4. Se debe generar la función que retorne el controlador
	5. Se deben generar todos los métodos correspondientes a los endpoints


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 32
-rw-r--r--@ 1 juurbano  1010544492  2802 Oct  4 22:14 README.md
drwxr-xr-x  3 juurbano  1010544492    96 Oct  4 21:59 cmd
-rw-r--r--  1 juurbano  1010544492  1064 Oct  4 22:17 go.mod
-rw-r--r--  1 juurbano  1010544492  6962 Oct  4 22:17 go.sum
drwxr-xr-x  3 juurbano  1010544492    96 Oct  4 21:58 internal

./cmd:
total 0
drwxr-xr-x  4 juurbano  1010544492  128 Oct  4 22:34 server

./cmd/server:
total 8
drwxr-xr-x  3 juurbano  1010544492   96 Oct  4 22:36 handler
-rw-r--r--  1 juurbano  1010544492  403 Oct  4 22:19 main.go

./cmd/server/handler:
total 8
-rw-r--r--  1 juurbano  1010544492  2331 Oct  4 22:36 product.go

./internal:
total 0
drwxr-xr-x  4 juurbano  1010544492  128 Oct  4 22:25 products

./internal/products:
total 16
-rw-r--r--  1 juurbano  1010544492  1368 Oct  4 21:58 repository.go
-rw-r--r--  1 juurbano  1010544492  1012 Oct  4 21:58 service.go

$
```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

---

*Windows 1*

```bash
$ go run cmd/server/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /products/                --> github.com/ejercicio_2/cmd/server/handler.(*Product).Store.func1 (3 handlers)
[GIN-debug] GET    /products/                --> github.com/ejercicio_2/cmd/server/handler.(*Product).GetAll.func1 (3 handlers)
[GIN-debug] PUT    /products/:id             --> github.com/ejercicio_2/cmd/server/handler.(*Product).Update.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2022/10/04 - 22:35:54 | 200 |      42.208µs |       127.0.0.1 | POST     "/products/"

...


```
---

*Windows 2*

```bash
$ curl http://127.0.0.1:8080/products/ -X GET -H "token: 123456" -vvv ; echo "" | cat -e
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /products/ HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
> token: 123456
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 05 Oct 2022 03:35:41 GMT
< Content-Length: 54
<
* Connection #0 to host 127.0.0.1 left intact
null$

$
```

#### *Note*

El json obtenido arriba en el ejemplo retorna `null` ya que no hay recursos agregados en el array creado al momento de la ejecución del servidor.

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

