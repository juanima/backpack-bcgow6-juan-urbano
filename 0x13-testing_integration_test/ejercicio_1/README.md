# Ejercicio 1

## Enunciated *_Service/Repo/Db Update()_*

Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:

1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. 
3. Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el producto con mismo Id y los datos actualizados. Validar también que  Read() del Store haya sido ejecutado durante el test. 


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

>> `mock_file_integration.go`: creamos una estructura que implementa los metodos del store original, con esto logramos crear una instancia del struct `MockFileStoreIntegration` con ello podemos inyectarlo al repository, aquí terminamos la implementación de metodo `Write()` ya que este nos faltaba 

>> `service_test.go`: creamos los test de integración para probar el método `Update()` y que utiliza el *mock* de store, una integración: ---> relación `service`y `repository`

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go test -v ./internal/products/  -run TestServiceIntegrationUpdate
=== RUN   TestServiceIntegrationUpdate
--- PASS: TestServiceIntegrationUpdate (0.00s)
=== RUN   TestServiceIntegrationUpdateFail
--- PASS: TestServiceIntegrationUpdateFail (0.00s)
PASS
ok      github.com/ejercicio_2/internal/products        0.931s
```

---

