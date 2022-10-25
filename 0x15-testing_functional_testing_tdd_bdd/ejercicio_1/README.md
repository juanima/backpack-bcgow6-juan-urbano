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
│       │   ├── handler_test.go
│       │   └── middleware.go
│       └── main.go
├── coverage.out
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
├── servic
├── test
│   ├── hello
│   │   ├── mock_repository.go
│   │   ├── mock_service.go
│   │   └── mock_store.go
│   └── products_test.go
└── update_products.json

13 directories, 32 files
```

## Testing Update

* Functional test para la UpdateProduct

```bash
$ go test -v ./test/
=== RUN   TestUpdateGood
[GIN] 2022/10/25 - 17:06:34 | 200 |     307.208µs |       192.0.2.1 | PUT      "/api/v1/products/1379"
--- PASS: TestUpdateGood (0.00s)
PASS
ok      github.com/ejercicio_2/test     0.958s
```


