# Ejercicio 1

## Enunciated *_Test Unitario GetAll()_*

Generar un Stub del Store cuya función `“Read”` retorne dos productos con las especificaciones que deseen. Comprobar que `GetAll()` retorne la información exactamente igual a la esperada. Para esto:

1. Dentro de la carpeta `/internal/(producto/usuario/transacción)`, crear un archivo `repository_test.go` con el test diseñado.


## Run 

- Estructura de directorios, podemos observar el directorio generado `docs/` donde almacena la documentación

```bash
$ tree
.├── README.md
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
│       ├── repository_test.go
│       └── service.go
├── patch_products.json
├── pkg
│   ├── store
│   │   ├── file.go
│   │   └── stub_file.go
│   └── web
│       └── response.go
├── products.json
└── update_products.json

11 directories, 21 files
```

> Los test los puedes encontrar siguiendo `repository_test.go` y `stub_file.go`

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go test -v ./internal/products/
=== RUN   TestGetAll
--- PASS: TestGetAll (0.00s)
PASS
ok      github.com/ejercicio_2/internal/products        0.315s
```

---

