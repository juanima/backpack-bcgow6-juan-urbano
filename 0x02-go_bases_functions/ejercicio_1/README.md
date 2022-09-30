# Ejercicio 1

## Enunciated *_Impuestos de salario_*

Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

## Run 

- Estructura de directorios

```bash
$ ls -l
total 40
-rw-r--r--  1 juurbano  1010544492  1063 Sep 26 18:56 README.md
-rw-r--r--  1 juurbano  1010544492    28 Sep 26 19:34 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 26 19:32 main.go
-rw-r--r--  1 juurbano  1010544492   354 Sep 26 19:42 salary_tax.go
-rw-r--r--  1 juurbano  1010544492    33 Sep 26 16:04 testcase.txt
$
```

- test cases
    * la primera linea es el número de testcases
    * la segunda linea representa un delimitador entre cada test
    * la tercera linea es el input en este caso el salario

```
$ cat testcase.txt
1
-----------------------
58000

$
```
### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt ; echo "" | cat -e
-----------------------
El impuesto del salario es $9860.000 USD
$
$
```

### Option 2
- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
CO000PFW3GXDTQ9:ejercicio_1 juurbano$ go run main.go salary_tax.go < testcase.txt ; echo "" | cat -e
-----------------------
El impuesto del salario es $9860.000 USD
$
CO000PFW3GXDTQ9:ejercicio_1 juurbano$
```


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_1
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

