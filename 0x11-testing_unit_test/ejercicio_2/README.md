# Ejercicio 2

## Enunciated *_Test Unitario Método Ordenar_*

Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente, posteriormente diseñar un test unitario que valide el funcionamiento del mismo.
1. Dentro de la carpeta `go-testing` crear un archivo `ordenamiento.go` con la función a probar.
2. Dentro de la carpeta go-testing crear un archivo `ordenamiento_test.go` con el test diseñado.


- Estructura de directorios

```bash
$ tree 
.
├── README.md
├── calculator
│   ├── calculadora.go
│   └── calculadora_test.go
├── go.mod
├── go.sum
└── main.go

1 directory, 6 files
$
```

### Run

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go test -v ./calculator/
=== RUN   TestSortSliceGood
--- PASS: TestSortSliceGood (0.00s)
=== RUN   TestSortSliceGoodSameArray
--- PASS: TestSortSliceGoodSameArray (0.00s)
PASS
ok      github.com/ejercicio_2/calculator       1.549s
$
```

- Hacemos pruebas desde el main

```bash
$ go run main.go
[8 9 13 30]
$
```
