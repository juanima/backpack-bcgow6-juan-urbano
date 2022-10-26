# desafio-cierre-testing
Proyecto Base para cumplir el Desafio de Testing

# desafio-cierre-testing

Ya está desarrollada una API que lista artículos disponibles para un vendedor determinado. Esta API no ha podido implementarse porque no cumple con los requisitos mínimos de calidad. Debe contener los tests necesarios para cubrir un 80% de coverage.

El proyecto consiste en una API que cuenta con un método llamado GetProducts y recibe por queryParam el ID (string) del vendedor (`seller_id`).

Link project: [link](https://drive.google.com/file/d/1v4kYwwrR2nYEU6yEb518a8PIicl8vqD5/view)

## *Objetivos*

* Validar que la API esté en correcto funcionamiento.
* Aplicar Mock y Stub tests para desarrollar los Tests de los archivos:
        * Service: crear el archivo *service_test.go*, diseñar los `mocks/stubs` necesarios para probar todo el service.
        * Handler: crear el archivo *handler_test.go*. Será necesario diseñar un mock del framework Gin para emular requests y probar el handler en su totalidad.
* Alcanzar un coverage del 80% total del proyecto.
* Todos los tests ejecutados deben culminar ok.

---

# Part 1: Check status API

*Window 1*

```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/products          --> github.com/bootcamp-go/desafio-cierre-testing/internal/products.(*Handler).GetProducts-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :18085

[GIN] 2022/10/25 - 21:34:54 | 200 |     100.834µs |       127.0.0.1 | GET      "/api/v1/products?seller_id=FEX112AC"
. . .
```

*Window 2*

```bash
$ curl http://127.0.0.1:18085/api/v1/products?seller_id=FEX112AC -X GET -vvv ; echo "" | cat -e
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying 127.0.0.1:18085...
* Connected to 127.0.0.1 (127.0.0.1) port 18085 (#0)
> GET /api/v1/products?seller_id=FEX112AC HTTP/1.1
> Host: 127.0.0.1:18085
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 26 Oct 2022 02:34:54 GMT
< Content-Length: 84
<
* Connection #0 to host 127.0.0.1 left intact
[{"ID":"mock","SellerID":"FEX112AC","Description":"generic product","Price":123.55}]$
$
```

---

## Part 2: Unit test handler using service mock

```bash
$ go test -v ./...
?       github.com/bootcamp-go/desafio-cierre-testing/cmd       [no test files]
?       github.com/bootcamp-go/desafio-cierre-testing/cmd/router        [no test files]
=== RUN   TestServiceGetAllBySeller
[GIN] 2022/10/26 - 14:59:20 | 200 |     183.792µs |       192.0.2.1 | GET      "/api/v1/products?seller_id=FEX112AC"
--- PASS: TestServiceGetAllBySeller (0.00s)
=== RUN   TestServiceRepoGetAllFail
[GIN] 2022/10/26 - 14:59:20 | 500 |       7.625µs |       192.0.2.1 | GET      "/api/v1/products?seller_id=FEX112AC"
--- PASS: TestServiceRepoGetAllFail (0.00s)
PASS
ok      github.com/bootcamp-go/desafio-cierre-testing/internal/products 1.177s
$
```

## Extra

* linter

```bash
$ golangci-lint run
cmd/main.go:12:7: Error return value of `r.Run` is not checked (errcheck)
        r.Run(":18085")
             ^
$
```

* Coverage Report

```bash
$ go test -cover -coverprofile=coverage.out  ./...
?       github.com/bootcamp-go/desafio-cierre-testing/cmd       [no test files]
?       github.com/bootcamp-go/desafio-cierre-testing/cmd/router        [no test files]
ok      github.com/bootcamp-go/desafio-cierre-testing/internal/products 1.136s  coverage: 59.4% of statements
$
```

* Show coverage report

```bash
go tool cover -html=coverage.out
```

* terminal coverage report

```bash
$ go test -cover ./...
?       github.com/bootcamp-go/desafio-cierre-testing/cmd       [no test files]
?       github.com/bootcamp-go/desafio-cierre-testing/cmd/router        [no test files]
ok      github.com/bootcamp-go/desafio-cierre-testing/internal/products 1.116s  coverage: 59.4% of statements
$
```


