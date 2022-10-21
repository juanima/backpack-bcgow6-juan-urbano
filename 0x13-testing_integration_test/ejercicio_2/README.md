# Ejercicio 1

## Enunciated *_Service/Repo/Db Delete()_*

Diseñar un test que pruebe en la capa service, el método o función `Delete()`. Se debe probar la correcta eliminación de un producto, y el error cuando el producto no existe. Para lograrlo puede:

1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
2. Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de Storage.
3. Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage luego del Delete(). También que cuando se intenta borrar un producto  inexistente, se debe obtener el error correspondiente.

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

> Los test los puedes encontrar siguiendo `service_test.go` y `mock_file_integration.go`

>> `mock_file_integration.go`: creamos una estructura que implementa los metodos del store original, con esto logramos crear una instancia del struct `MockFileStoreIntegration` con ello podemos inyectarlo al repository  

>> `service_test.go`: creamos los test de integración para probar el método `Delete()` y que utiliza el *mock* de store, una integración: ---> relación `service`y `repository`

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go test -v ./internal/products/  -run TestServiceIntegrationDelete
=== RUN   TestServiceIntegrationDelete
--- PASS: TestServiceIntegrationDelete (0.00s)
=== RUN   TestServiceIntegrationDeleteFail
--- PASS: TestServiceIntegrationDeleteFail (0.00s)
PASS
ok      github.com/ejercicio_2/internal/products        0.960s
```

---

