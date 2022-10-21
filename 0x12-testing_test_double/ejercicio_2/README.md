# Ejercicio 2

## Enunciated *_Test Unitario UpdateName()_*

Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:

1. Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo nombre puede ser “Before Update”.
2. 2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase. 
3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test. 

## Estructura de directorios

```bash
$ tree
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
│       └── service.go
├── patch_products.json
├── pkg
│   ├── store
│   │   ├── file.go
│   │   ├── mock_file.go
│   │   └── stub_file.go
│   └── web
│       └── response.go
├── products.json
└── update_products.json

11 directories, 23 files
```

> Los test los puedes encontrar siguiendo `repository_mock_test.go` y `mock_file.go`

>> `mock_file.go`: creamos una estructura que implementa los metodos del repository original, con esto logramos crear una instancia del struct `MockStoreFile` con ello podemos inyectarlo al repository

>> `repository_mock_test.go`: creamos los test para probar el método `UpdateName()` y que utiliza el *mock* de store

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go test -v ./internal/products/
=== RUN   TestUpdateNameGood
--- PASS: TestUpdateNameGood (0.00s)
=== RUN   TestUpdateNameBad
--- PASS: TestUpdateNameBad (0.00s)
=== RUN   TestGetAll
--- PASS: TestGetAll (0.00s)
PASS
ok      github.com/ejercicio_2/internal/products        0.987s
```

---

