# Ejercicio 2

## Enunciated *_Medir Coverage_*

Se debe medir el Coverage del proyecto Go - Web, para eso se deben seguir los siguientes pasos:

1. Medir el code coverage del proyecto Go - Web.
2. Generar el coverage report del proyecto Go - Web
3. Identificar todos las partes del código que no tienen o tienen poca cobertura.

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

# Part 1: Install golangci-lint

```bash
$ golangci-lint --version
golangci-lint has version 1.50.1 built from 8926a95 on 2022-10-22T10:48:48Z
$
```

# Part 2:

* Nos ubicamos en la raíz del proyecto
* Errores encontrados `before`

```bash
$ golangci-lint run
cmd/server/handler/handler.go:42:6: func `getToken` is unused (unused)
func getToken(name string) (string, error) {
     ^
internal/domain/product.go:4:9: structtag: struct field tag `json: "id"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Id      int     `json: "id"`
        ^
internal/domain/product.go:5:2: structtag: struct field tag `json: "stock"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Stock   int     `json: "stock"`
        ^
internal/domain/product.go:6:2: structtag: struct field tag `json: "name"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Name    string  `json: "name"`
        ^
internal/domain/product.go:7:2: structtag: struct field tag `json: "color"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Color   string  `json: "color"`
        ^
internal/domain/product.go:8:2: structtag: struct field tag `json: "code"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Code    string  `json: "code"`
        ^
internal/domain/product.go:9:2: structtag: struct field tag `json: "price"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Price   float64 `json: "price"`
        ^
internal/domain/product.go:10:2: structtag: struct field tag `json: "publish"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Publish *bool   `json: "publish"`
        ^
internal/domain/product.go:11:2: structtag: struct field tag `json: "date"` not compatible with reflect.StructTag.Get: bad syntax for struct tag value (govet)
        Date    string  `json: "date"`
        ^
internal/products/service_test.go:132:2: ineffectual assignment to err (ineffassign)
        err := service.Delete(1399)
        ^
$
```

# Part 3

* Luego de resolver los errores `after`

```bash
$ golangci-lint run
$
```

---

