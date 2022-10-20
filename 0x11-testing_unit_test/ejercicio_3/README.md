# Ejercicio 3

## Enunciated *_Test unitario método dividir_*


Para el Método Dividir, visto en la clase:

```go
func Dividir(num, den int) int {
        return num / den
}
```

Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador es igual a 0,  retorna un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.

1. Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
2. Dentro de la carpeta go-testing crear un archivo dividir test.go con el test diseñado.

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
=== RUN   TestDivideGood
--- PASS: TestDivideGood (0.00s)
=== RUN   TestDivideFail
--- PASS: TestDivideFail (0.00s)
PASS
ok      github.com/ejercicio_3/calculator       0.916s
$
```

- Hacemos pruebas desde el main

```bash
$ go run main.go
2
$
```
