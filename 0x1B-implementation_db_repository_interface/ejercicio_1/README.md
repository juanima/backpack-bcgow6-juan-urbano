# Test

## Ejercicio 1 - Implementar GetByName()
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:

-  Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorna una estructura del tipo Product.

-  Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.



## Ejercicio 2 - Replicar Store()
Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.



## Ejercicio 3 - Ejecutar Store()
Diseñar un Test que permita ejecutar Store y comprobar la correcta inserción del registro en la tabla.

## Ejercicio 4 - Ejecutar GetByName()
Diseñar un Test que permita ejecutar GetByName y comprobar que retorne el registro buscado en caso de que exista. 

# Run

* Import database from sql script file

```bash
$ mysql -uroot -p < db/db.sql
Enter password:
$
```

* Run web server form our API

```bash
CO000PFW3GXDTQ9:ejercicio_1 juurbano$ go run cmd/server/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/products/         --> github.com/bootcamp-go/storage/cmd/server/handler.(*Product).Store.func1 (3 handlers)
[GIN-debug] GET    /api/v1/products/         --> github.com/bootcamp-go/storage/cmd/server/handler.(*Product).GetByName.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080

. . .
```

# Test

* Consume `POST /api/v1/movies/`

> New example `create_movie.json`

```json
{
        "title": "my new movie",
        "rating": 3.4,
        "awards": 4,
        "length": 20,
        "release_date": "2018-07-21 18:12:42.023"
}
```

```bash
$ curl http://127.0.0.1:8080/api/v1/movies/ -X POST -H "Content-Type: application/json" -d @create_movie.json -vvv; echo "" | cat -e
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /api/v1/movies/ HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 118
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 201 Created
< Content-Type: application/json; charset=utf-8
< Date: Wed, 09 Nov 2022 17:23:58 GMT
< Content-Length: 217
<
* Connection #0 to host 127.0.0.1 left intact
{"data":{"id":24,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"my new movie","rating":3.4,"awards":4,"release_date":"2018-07-21 18:12:42.023","length":20,"genre_id":null},"code":201}$
$
```

* Consume `GET /api/v1/movies`

> New example `getbyname_movie.json`

```json
{
        "title": "my new movie"
}
```


```bash
$ curl http://127.0.0.1:8080/api/v1/movies/ -X GET -H "Content-Type: application/json" -d @getbyname_movie.json -vvv; echo "" | cat -e
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /api/v1/movies/ HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 33
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 09 Nov 2022 17:35:21 GMT
< Content-Length: 194
<
* Connection #0 to host 127.0.0.1 left intact
{"data":{"id":24,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"my new movie","rating":3.4,"awards":4,"release_date":"","length":20,"genre_id":null},"code":200}$
$
```




