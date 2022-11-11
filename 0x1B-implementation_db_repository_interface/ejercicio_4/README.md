# Práctica clase 1 - Storage Implementation

## Extra - Diseñar test para las funcionalidades 

Diseñar test para todos los métodos de la capa de repository

---

# Test

```bash
$ go test -v ./internal/movies/
=== RUN   TestSave
=== RUN   TestSave/Store_Ok
=== RUN   TestSave/Store_Fail
--- PASS: TestSave (0.00s)
    --- PASS: TestSave/Store_Ok (0.00s)
    --- PASS: TestSave/Store_Fail (0.00s)
=== RUN   Test_RepositoryGetAll
=== RUN   Test_RepositoryGetAll/GetAll_Ok
=== RUN   Test_RepositoryGetAll/GetAll_Fail
--- PASS: Test_RepositoryGetAll (0.00s)
    --- PASS: Test_RepositoryGetAll/GetAll_Ok (0.00s)
    --- PASS: Test_RepositoryGetAll/GetAll_Fail (0.00s)
PASS
ok      github.com/bootcamp-go/storage/internal/movies  0.906s
$
```

---

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

* Consume `GET /api/v1/movies`

```bash
$ curl http://127.0.0.1:8080/api/v1/movies -X GET -vvv ; echo "" | cat -e
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /api/v1/movies HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 09 Nov 2022 22:31:22 GMT
< Transfer-Encoding: chunked
<
{"data":[{"id":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Avatar","rating":7.9,"awards":3,"release_date":"","length":120,"genre_id":5},{"id":2,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Titanic","rating":7.7,"awards":11,"release_date":"","length":320,"genre_id":3},{"id":3,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"La Guerra de las galaxias: Episodio VI","rating":9.1,"awards":7,"release_date":"","length":null,"genre_id":5},{"id":4,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"La Guerra de las galaxias: Episodio VII","rating":9,"awards":6,"release_date":"","length":180,"genre_id":5},{"id":5,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Parque Jurasico","rating":8.3,"awards":5,"release_date":"","length":270,"genre_id":5},{"id":6,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Harry Potter y las Reliquias de la Muerte - Parte 2","rating":9,"awards":2,"release_date":"","length":190,"genre_id":6},{"id":7,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Transformers: el lado oscuro de la luna","rating":0.9,"awards":1,"release_date":"","length":null,"genre_id":5},{"id":8,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Harry Potter y la piedra filosofal","rating":10,"awards":1,"release_date":"","length":120,"genre_id":8},{"id":9,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Harry Potter y la cámara de los secretos","rating":3.5,"awards":2,"release_date":"","length":200,"genre_id":8},{"id":10,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"El rey león","rating":9.1,"awards":3,"release_date":"","length":null,"genre_id":10},{"id":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Alicia en el país de las maravillas","rating":5.7,"awards":2,"release_date":"","length":120,"genre_id":null},{"id":12,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Buscando a Nemo","rating":8.3,"awards":2,"release_date":"","length":110,"genre_id":7},{"id":13,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Toy Story","rating":6.1,"awards":0,"release_date":"","length":150,"genre_id":7},{"id":14,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Toy Story 2","rating":3.2,"awards":2,"release_date":"","length":120,"genre_id":7},{"id":15,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"La vida es bella","rating":8.3,"awards":5,"release_date":"","length":null,"genre_id":3},{"id":16,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Mi pobre angelito","rating":3.2,"awards":1,"release_date":"","length":120,"genre_id":1},{"id":17,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Intensamente","rating":9,"awards":2,"release_date":"","length":120,"genre_id":7},{"id":18,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Carrozas de fuego","rating":9.9,"awards":7,"release_date":"","length":180,"genre_id":null},{"id":19,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Big","rating":7.3,"awards":2,"release_date":"","length":130,"genre_id":8},{"id":20,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"I am Sam","rating":9,"awards":4,"release_date":"","length":130,"genre_id":3},{"id":21,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"Hotel Transylvania","rating":7.1,"awards":1,"release_date":"","length":90,"genre_id":10},{"id":23,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"NUEVA PELI","rating":7.9,"awards":3,"release_date":"","length":120,"genre_id":13},{"id":24,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"my new* Connection #0 to host 127.0.0.1 left intact
 movie","rating":3.4,"awards":4,"release_date":"","length":20,"genre_id":null}],"code":200}$
$
```

* Consume `DELETE /api/v1/movies/:id`

```bash
$ curl http://127.0.0.1:8080/api/v1/movies/24 -X DELETE -vvv ; echo ""
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> DELETE /api/v1/movies/24 HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 204 No Content
< Content-Type: application/json; charset=utf-8
< Date: Thu, 10 Nov 2022 15:32:20 GMT
<
* Connection #0 to host 127.0.0.1 left intact

$
```

# Coverage

* the following test are related about `Repository Layer`

```bash
$ go test -v -cover -run Test_Delete ./internal/movies/
=== RUN   Test_Delete
=== RUN   Test_Delete/DeleteOK
=== RUN   Test_Delete/Delete_Fail
--- PASS: Test_Delete (0.00s)
    --- PASS: Test_Delete/DeleteOK (0.00s)
    --- PASS: Test_Delete/Delete_Fail (0.00s)
PASS
coverage: 18.6% of statements
ok      github.com/bootcamp-go/storage/internal/movies  (cached)        coverage: 18.6% of statements
$
```

## Coverage output html

1. Generate `coverage.out` file 
2. Render output from `coverage.out` file using *_html_*

```bash
$ go test -v -cover -coverprofile=coverage.out ./internal/movies/
=== RUN   TestSave
=== RUN   TestSave/Store_Ok
=== RUN   TestSave/Store_Fail
--- PASS: TestSave (0.00s)
    --- PASS: TestSave/Store_Ok (0.00s)
    --- PASS: TestSave/Store_Fail (0.00s)
=== RUN   Test_RepositoryGetAll
=== RUN   Test_RepositoryGetAll/GetAll_Ok
=== RUN   Test_RepositoryGetAll/GetAll_Fail
--- PASS: Test_RepositoryGetAll (0.00s)
    --- PASS: Test_RepositoryGetAll/GetAll_Ok (0.00s)
    --- PASS: Test_RepositoryGetAll/GetAll_Fail (0.00s)
=== RUN   Test_Delete
=== RUN   Test_Delete/DeleteOK
=== RUN   Test_Delete/Delete_Fail
--- PASS: Test_Delete (0.00s)
    --- PASS: Test_Delete/DeleteOK (0.00s)
    --- PASS: Test_Delete/Delete_Fail (0.00s)
PASS
coverage: 57.6% of statements
ok      github.com/bootcamp-go/storage/internal/movies  1.037s  coverage: 57.6% of statements
$
$ go tool cover -html=coverage.out
$
```

# Linter

*Important* check the code

```bash
$ golangci-lint run
tests/functional/movies_test.go:31:15: NewProduct not declared by package handler (typecheck)
        p := handler.NewProduct(serv)
                     ^
cmd/server/main.go:27:7: Error return value of `r.Run` is not checked (errcheck)
        r.Run()
             ^
$
```

