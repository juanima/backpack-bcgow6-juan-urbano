# Ejercicio 1

## Enunciated *_Test unitario restar_*


Para el método `Restar()` visto en la clase, realizar el test unitario correspondiente. Para esto:
1. Dentro de la carpeta go-testing crear un archivo `calculadora.go` con la función a probar.
2. Dentro de la carpeta go-testing crear un archivo `calculadora_test.go` con el test diseñado.

## Run 

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
=== RUN   TestSubGood
--- PASS: TestSubGood (0.00s)
=== RUN   TestGoodSubNegativeMinuendSubtrahend
--- PASS: TestGoodSubNegativeMinuendSubtrahend (0.00s)
PASS
ok      github.com/ejercicio_1/calculator       1.043s
$
```
