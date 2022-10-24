# Ejercicio 2

## Enunciated *_Functional Testing Update()_*

Se requiere probar la funcionalidad de “actualización de producto”, pudiendo reutilizar las funciones creadas en la clase. Para lograrlo realizar los siguientes pasos:

* Dentro de la carpeta /test, crear un archivo `products_test.go`.
* Levantar el Servidor y definir la ruta para este test.
* Crear Request y Response apropiados.
* Solicitar al servidor que atienda al Request.
* Validar Response.
 

## Estructura de directorios

```bash
.
├── README.md
├── cmd
│   └── server
│       ├── handler
│       │   ├── handler.go
│       │   └── middleware.go
│       └── main.go
├── create_products.json
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── images
│   └── swagger-web.png
├── internal
│   ├── domain
│   │   └── product.go
│   └── products
│       ├── repository.go
│       ├── repository_mock_test.go
│       ├── repository_stub_test.go
│       ├── service.go
│       └── service_test.go
├── patch_products.json
├── pkg
│   ├── store
│   │   ├── file.go
│   │   ├── mock_file.go
│   │   ├── mock_file_integration.go
│   │   └── stub_file.go
│   └── web
│       └── response.go
├── products.json
└── update_products.json

11 directories, 25 files
```

# Part 1: Initial coverage

* Analisis de archivos con test propuestos

```bash
$ go test -cover ./...
?       github.com/ejercicio_2/cmd/server       [no test files]
?       github.com/ejercicio_2/cmd/server/handler       [no test files]
?       github.com/ejercicio_2/docs     [no test files]
?       github.com/ejercicio_2/internal/domain  [no test files]
ok      github.com/ejercicio_2/internal/products        0.860s  coverage: 51.7% of statements
?       github.com/ejercicio_2/pkg/store        [no test files]
?       github.com/ejercicio_2/pkg/web  [no test files]
$
```

* Directorio `./internal/products` donde se presentó el análisis de covertura

```bash
$ ls -1 internal/products/
repository.go
repository_mock_test.go
repository_stub_test.go
service.go
service_test.go
$
```

> Archivos test: `repository_mock_test.go`,  `repository_mock_test.go`,  `repository_mock_test.go` y `service_test.go`  

# Part 2: Coverage report

```bash
$ go test -cover -coverprofile=coverage.out ./...
?       github.com/ejercicio_2/cmd/server       [no test files]
?       github.com/ejercicio_2/cmd/server/handler       [no test files]
?       github.com/ejercicio_2/docs     [no test files]
?       github.com/ejercicio_2/internal/domain  [no test files]
ok      github.com/ejercicio_2/internal/products        0.386s  coverage: 51.7% of statements
?       github.com/ejercicio_2/pkg/store        [no test files]
?       github.com/ejercicio_2/pkg/web  [no test files]
$
```

* Archivo `coverage.out` generado luego del ejecutar el comando para generar el reporte

```bash
$ ls -l
total 112
-rw-r--r--@ 1 juurbano  1010544492   4517 Oct 23 16:29 README.md
drwxr-xr-x  3 juurbano  1010544492     96 Oct 23 10:41 cmd
-rw-r--r--  1 juurbano  1010544492   5301 Oct 23 16:32 coverage.out
-rw-r--r--  1 juurbano  1010544492    144 Oct 23 10:41 create_products.json
drwxr-xr-x  5 juurbano  1010544492    160 Oct 23 10:41 docs
-rw-r--r--  1 juurbano  1010544492   1916 Oct 23 10:41 go.mod
-rw-r--r--  1 juurbano  1010544492  16542 Oct 23 10:41 go.sum
drwxr-xr-x  3 juurbano  1010544492     96 Oct 23 10:41 images
drwxr-xr-x  4 juurbano  1010544492    128 Oct 23 10:41 internal
-rw-r--r--  1 juurbano  1010544492     21 Oct 23 10:41 patch_products.json
drwxr-xr-x  4 juurbano  1010544492    128 Oct 23 10:41 pkg
-rw-r--r--  1 juurbano  1010544492    156 Oct 23 10:41 products.json
-rw-r--r--  1 juurbano  1010544492    144 Oct 23 10:41 update_products.json
$
```

```bash

```

# Part 3

* Luego de aumentar agregar covertura a los test existentes y agregar nuevos test tanto para `service.go` y `repository.go`

```bash
$ go test -cover -coverprofile=coverage.out ./...
?       github.com/ejercicio_2/cmd/server       [no test files]
?       github.com/ejercicio_2/cmd/server/handler       [no test files]
?       github.com/ejercicio_2/docs     [no test files]
?       github.com/ejercicio_2/internal/domain  [no test files]
ok      github.com/ejercicio_2/internal/products        0.357s  coverage: 80.5% of statements
?       github.com/ejercicio_2/pkg/store        [no test files]
?       github.com/ejercicio_2/pkg/web  [no test files]
$
```

# Parte 4

```
ok      github.com/ejercicio_2/internal/products        0.386s  coverage: 51.7% of statements
ok      github.com/ejercicio_2/internal/products        0.357s  coverage: 80.5% of statements
```

* Con esto podemos notar que `30%` de la covertura ha aumentado, cubriendo en la mayoria los casos positivos 


---

