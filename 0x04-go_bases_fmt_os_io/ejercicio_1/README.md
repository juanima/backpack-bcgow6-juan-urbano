# Ejercicio 1

## Enunciated *_Guardar archivo_*

Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.

## Run 

- Estructura de directorios

```bash
$ ls -l
total 40
-rw-r--r--  1 juurbano  1010544492  1063 Sep 26 18:56 README.md
-rw-r--r--  1 juurbano  1010544492    28 Sep 26 19:34 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 26 19:32 main.go
-rw-r--r--  1 juurbano  1010544492   779 Sep 26 23:25 save_file.go
-rw-r--r--  1 juurbano  1010544492    33 Sep 26 16:04 testcase.txt
$
```

- test cases
    * la primera linea es el número de testcases
    * la segunda linea es el número de productos para guardar
    * las siguientes líneas correspondes a los datos de los productos 
	* `ID` `Price `cantidad`

```
$ cat testcase.txt
1
-----------------------1
3
111223 30012.00 1
444321 1000000.00 4
434321 50.50 1
-----------------------2
1
111223 30012.00 1

```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt
-----------------------1
$
$ cat myFile.csv
111223,30012.00,1
444321,1000000.00,4
434321,50.50,1

$
```


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

